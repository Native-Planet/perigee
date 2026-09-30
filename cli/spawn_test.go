package cli

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Native-Planet/perigee/roller"
	"github.com/Native-Planet/perigee/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestSpawnCommand(t *testing.T) {
	keyHex := strings.Repeat("1", 64)
	key, _ := crypto.HexToECDSA(keyHex)
	owner := crypto.PubkeyToAddress(key.PublicKey).Hex()
	var point types.Point
	point.Dominion = "l2"
	point.Ownership.Owner.Address = owner
	submitted := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID     string          `json:"id"`
			Method string          `json:"method"`
			Params json.RawMessage `json:"params"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Error(err)
			return
		}
		var result any
		switch req.Method {
		case "getPoint":
			result = point
		case "getNonce":
			result = 0
		case "getUnsignedTx":
			result = "0x" + strings.Repeat("0", 64)
		case "spawn":
			var spawn types.SpawnRequest
			if err := json.Unmarshal(req.Params, &spawn); err != nil {
				t.Error(err)
			}
			if spawn.Data.Ship != "~marzod" || !strings.EqualFold(spawn.Data.Address, owner) {
				t.Errorf("wrong spawn: %+v", spawn.Data)
			}
			submitted = true
			result = "spawn-hash"
		default:
			t.Errorf("unexpected method %s", req.Method)
		}
		json.NewEncoder(w).Encode(map[string]any{"jsonrpc": "2.0", "id": req.ID, "result": result})
	}))
	old := roller.Client
	roller.Client = roller.New(roller.Config{Endpoint: server.URL, HTTPClient: server.Client()})
	t.Cleanup(func() { roller.Client = old; server.Close() })
	cmd := NewSpawnCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetArgs([]string{"--point=~zod", "--child=~marzod", "--owner=" + owner, "--private-key=0x" + keyHex})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	var tx types.Transaction
	if err := json.Unmarshal(out.Bytes(), &tx); err != nil {
		t.Fatal(err)
	}
	if !submitted || tx.Hash != "spawn-hash" {
		t.Fatalf("spawn output: %s", out.String())
	}
}

func TestSpawnCommandRequiresArguments(t *testing.T) {
	for _, args := range [][]string{
		{},
		{"--point=~zod", "--child=~marzod", "--owner=0x1234"},
		{"--point=~zod", "--child=~marzod", "--owner=0x1234", "--private-key=key", "--master-ticket=ticket"},
	} {
		cmd := NewSpawnCmd()
		cmd.SetArgs(args)
		cmd.SetOut(new(bytes.Buffer))
		cmd.SetErr(new(bytes.Buffer))
		if err := cmd.Execute(); err == nil {
			t.Fatalf("accepted invalid arguments: %v", args)
		}
	}
}
