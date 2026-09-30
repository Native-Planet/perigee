package libprg

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Native-Planet/perigee/roller"
	"github.com/Native-Planet/perigee/types"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func TestBreachReportsRPCFailureWithoutKeyMaterialError(t *testing.T) {
	var point types.Point
	point.Dominion = "l1"
	point.Network.Keys.Life = "7"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req ethereumTestRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Error(err)
			return
		}
		if req.Method != "getPoint" {
			t.Errorf("unexpected roller method: %s", req.Method)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "id": req.ID, "result": point})
	}))
	previous := roller.Client
	roller.Client = roller.New(roller.Config{Endpoint: server.URL, HTTPClient: server.Client()})
	t.Cleanup(func() { roller.Client = previous; server.Close() })
	withEthereumRPC(t, func(w http.ResponseWriter, req ethereumTestRequest) {
		if req.Method == "eth_getTransactionCount" {
			writeEthereumError(w, req.ID, -32603, "Internal error")
			return
		}
		writeEthereumResult(t, w, req)
	})
	_, err := Breach(context.Background(), "~dilner-digfer", "sampel-ticlyt-migfun-falmel", "", "")
	var rpcErr rpc.Error
	if err == nil || errors.Is(err, ErrKeyMaterial) || !errors.As(err, &rpcErr) || rpcErr.ErrorCode() != -32603 {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(err.Error(), "configure L1 network keys for ~dilner-digfer") || !strings.Contains(err.Error(), "eth_getTransactionCount (pending)") {
		t.Fatalf("missing operation context: %v", err)
	}
}

func TestEthereumPreparationRetriesTransientFailures(t *testing.T) {
	for _, method := range []string{"net_version", "eth_gasPrice", "eth_getCode", "eth_estimateGas", "eth_getTransactionCount"} {
		t.Run(method, func(t *testing.T) {
			calls := make(map[string]int)
			withEthereumRPC(t, func(w http.ResponseWriter, req ethereumTestRequest) {
				calls[req.Method]++
				if req.Method == method && calls[req.Method] == 1 {
					writeEthereumError(w, req.ID, -32603, "Internal error")
					return
				}
				writeEthereumResult(t, w, req)
			})
			tx := prepareTestEthereumTransaction(t)
			if tx.Gas() != 63463 || tx.Nonce() != 30 || tx.ChainId().Cmp(big.NewInt(1)) != 0 {
				t.Fatalf("unexpected transaction: gas=%d nonce=%d chain=%s", tx.Gas(), tx.Nonce(), tx.ChainId())
			}
			if calls[method] != 2 || calls["eth_sendRawTransaction"] != 0 {
				t.Fatalf("unexpected RPC calls: %v", calls)
			}
		})
	}
}

func TestEthereumPreparationRetriesRateLimit(t *testing.T) {
	calls := 0
	withEthereumRPC(t, func(w http.ResponseWriter, req ethereumTestRequest) {
		if req.Method == "eth_estimateGas" {
			calls++
			if calls == 1 {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
		}
		writeEthereumResult(t, w, req)
	})
	prepareTestEthereumTransaction(t)
	if calls != 2 {
		t.Fatalf("gas estimation calls = %d, want 2", calls)
	}
}

func TestEthereumReadFailureReportsMethod(t *testing.T) {
	for _, tc := range []struct {
		name     string
		code     int
		attempts int
	}{
		{"internal error", -32603, 3},
		{"execution reverted", 3, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			withEthereumRPC(t, func(w http.ResponseWriter, req ethereumTestRequest) {
				calls++
				writeEthereumError(w, req.ID, tc.code, tc.name)
			})
			client, err := ethclient.Dial(EthProvider)
			if err != nil {
				t.Fatal(err)
			}
			defer client.Close()
			_, err = (ethereumBackend{client}).PendingNonceAt(context.Background(), common.Address{})
			var rpcErr rpc.Error
			if err == nil || !strings.Contains(err.Error(), "eth_getTransactionCount (pending)") || !errors.As(err, &rpcErr) || rpcErr.ErrorCode() != tc.code {
				t.Fatalf("unexpected error: %v", err)
			}
			if calls != tc.attempts {
				t.Fatalf("calls = %d, want %d", calls, tc.attempts)
			}
		})
	}
}

func TestEthereumSubmissionIsNotRetried(t *testing.T) {
	calls := 0
	withEthereumRPC(t, func(w http.ResponseWriter, req ethereumTestRequest) {
		if req.Method != "eth_sendRawTransaction" {
			t.Errorf("unexpected method: %s", req.Method)
		}
		calls++
		writeEthereumError(w, req.ID, -32603, "Internal error")
	})
	client, err := ethclient.Dial(EthProvider)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	tx := ethTypes.NewTx(&ethTypes.LegacyTx{Gas: 21000, GasPrice: big.NewInt(1)})
	err = (ethereumBackend{client}).SendTransaction(context.Background(), tx)
	if calls != 1 || err == nil || !strings.Contains(err.Error(), "eth_sendRawTransaction") || !strings.Contains(err.Error(), tx.Hash().Hex()) {
		t.Fatalf("calls=%d error=%v", calls, err)
	}
}

func TestEthereumRetryRespectsCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	calls := 0
	_, err := ethereumRead(ctx, "eth_getCode", 3, func() (int, error) {
		calls++
		cancel()
		return 0, rpc.HTTPError{StatusCode: http.StatusTooManyRequests}
	})
	if calls != 1 || !errors.Is(err, context.Canceled) {
		t.Fatalf("calls=%d error=%v", calls, err)
	}
}

type ethereumTestRequest struct {
	ID     json.RawMessage `json:"id"`
	Method string          `json:"method"`
}

func withEthereumRPC(t *testing.T, handle func(http.ResponseWriter, ethereumTestRequest)) {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req ethereumTestRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Error(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		handle(w, req)
	}))
	previous := EthProvider
	EthProvider = server.URL
	t.Cleanup(func() { EthProvider = previous; server.Close() })
}

func writeEthereumError(w http.ResponseWriter, id json.RawMessage, code int, message string) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"jsonrpc": "2.0", "id": id,
		"error": map[string]interface{}{"code": code, "message": message},
	})
}

func writeEthereumResult(t *testing.T, w http.ResponseWriter, req ethereumTestRequest) {
	t.Helper()
	results := map[string]string{
		"net_version": "1", "eth_gasPrice": "0x5f5e100", "eth_getCode": "0x6000",
		"eth_estimateGas": "0xf7e7", "eth_getTransactionCount": "0x1e",
	}
	result, ok := results[req.Method]
	if !ok {
		t.Errorf("unexpected method: %s", req.Method)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "id": req.ID, "result": result})
}

func prepareTestEthereumTransaction(t *testing.T) *ethTypes.Transaction {
	t.Helper()
	key, err := crypto.HexToECDSA(strings.Repeat("1", 64))
	if err != nil {
		t.Fatal(err)
	}
	session, client, err := newEclipticSession(context.Background(), key)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	session.TransactOpts.NoSend = true
	tx, err := session.ConfigureKeys(1, [32]byte{1}, [32]byte{2}, 1, true)
	if err != nil {
		t.Fatal(err)
	}
	return tx
}
