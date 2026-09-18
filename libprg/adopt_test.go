package libprg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Native-Planet/perigee/roller"
	"github.com/Native-Planet/perigee/types"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestAdoptAcceptsManagementProxyKey(t *testing.T) {
	managementKeyHex := strings.Repeat("1", 64)
	managementKey, err := crypto.HexToECDSA(managementKeyHex)
	if err != nil {
		t.Fatal(err)
	}

	var pointInfo types.Point
	pointInfo.Dominion = "l2"
	pointInfo.Ownership.Owner.Address = "0x0000000000000000000000000000000000000001"
	pointInfo.Ownership.ManagementProxy.Address = crypto.PubkeyToAddress(managementKey.PublicKey).Hex()

	var adoptProxy string
	withAdoptRoller(t, &pointInfo, &adoptProxy)

	got, err := Adopt(context.Background(), "~fen", "~fonlyd", managementKeyHex, "")
	if err != nil {
		t.Fatalf("Adopt() error = %v", err)
	}
	tx, ok := got.(types.Transaction)
	if !ok {
		t.Fatalf("Adopt() returned %T, want types.Transaction", got)
	}
	if tx.Type != "adopt" {
		t.Errorf("transaction type = %q, want adopt", tx.Type)
	}
	if adoptProxy != "manage" {
		t.Errorf("adopt proxy = %q, want manage", adoptProxy)
	}
}

func TestAdoptRejectsSpawnProxyKey(t *testing.T) {
	spawnKeyHex := strings.Repeat("2", 64)
	spawnKey, err := crypto.HexToECDSA(spawnKeyHex)
	if err != nil {
		t.Fatal(err)
	}

	var pointInfo types.Point
	pointInfo.Dominion = "l2"
	pointInfo.Ownership.Owner.Address = "0x0000000000000000000000000000000000000001"
	pointInfo.Ownership.ManagementProxy.Address = "0x0000000000000000000000000000000000000002"
	pointInfo.Ownership.SpawnProxy.Address = crypto.PubkeyToAddress(spawnKey.PublicKey).Hex()
	var adoptProxy string
	withAdoptRoller(t, &pointInfo, &adoptProxy)

	_, err = Adopt(context.Background(), "~fen", "~fonlyd", spawnKeyHex, "")
	if !errors.Is(err, ErrKeyMaterial) {
		t.Fatalf("Adopt() error = %v, want ErrKeyMaterial", err)
	}
}

func withAdoptRoller(t *testing.T, pointInfo *types.Point, adoptProxy *string) {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Method string          `json:"method"`
			ID     string          `json:"id"`
			Params json.RawMessage `json:"params"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("decode request: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var result interface{}
		switch req.Method {
		case "getPoint":
			result = pointInfo
		case "getNonce":
			result = 0
		case "getUnsignedTx":
			result = "0x" + strings.Repeat("0", 64)
		case "adopt":
			var params struct {
				From types.FromData `json:"from"`
			}
			if err := json.Unmarshal(req.Params, &params); err != nil {
				t.Errorf("decode adopt params: %v", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			*adoptProxy = params.From.Proxy
			result = "test-transaction-hash"
		default:
			t.Errorf("unexpected method %q", req.Method)
			http.Error(w, "unexpected method", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(struct {
			Version string      `json:"jsonrpc"`
			Result  interface{} `json:"result"`
			ID      string      `json:"id"`
		}{
			Version: "2.0",
			Result:  result,
			ID:      req.ID,
		}); err != nil {
			t.Errorf("encode response: %v", err)
		}
	}))

	oldClient := roller.Client
	roller.Client = roller.New(roller.Config{
		Endpoint:   server.URL,
		HTTPClient: server.Client(),
	})
	t.Cleanup(func() {
		roller.Client = oldClient
		server.Close()
	})
}
