package libprg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Native-Planet/perigee/roller"
	"github.com/Native-Planet/perigee/types"
	"github.com/deelawn/urbit-gob/co"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nathanlever/keygen"
)

func withRPC(t *testing.T, respond func(string, json.RawMessage) (any, error)) {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Method string          `json:"method"`
			ID     string          `json:"id"`
			Params json.RawMessage `json:"params"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Error(err)
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		result, err := respond(req.Method, req.Params)
		resp := map[string]any{"jsonrpc": "2.0", "id": req.ID, "result": result}
		if err != nil {
			resp["error"] = map[string]any{"code": -32000, "message": err.Error()}
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Error(err)
		}
	}))
	old := roller.Client
	roller.Client = roller.New(roller.Config{Endpoint: server.URL, HTTPClient: server.Client()})
	t.Cleanup(func() { roller.Client = old; server.Close() })
}

func patpFor(t *testing.T, point uint32) string {
	t.Helper()
	patp, err := co.Point2Patp(new(big.Int).SetUint64(uint64(point)))
	if err != nil {
		t.Fatal(err)
	}
	return patp
}

func TestSpawnAuthorizationAndRouting(t *testing.T) {
	keyHex := strings.Repeat("1", 64)
	key, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		t.Fatal(err)
	}
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	for _, tc := range []struct {
		name, role, dominion, transferProxy, child, owner string
		wantErr                                           error
	}{
		{name: "owner", role: "own", dominion: "l2"},
		{name: "spawn proxy", role: "spawn", dominion: "l2"},
		{name: "spawn dominion", role: "spawn", dominion: "spawn"},
		{name: "zero transfer proxy", role: "spawn", dominion: "l1", transferProxy: "0x" + strings.Repeat("0", 40)},
		{name: "management proxy denied", role: "manage", dominion: "l2", wantErr: ErrKeyMaterial},
		{name: "unrelated child denied", role: "own", dominion: "l2", child: patpFor(t, 258), wantErr: ErrInvalidPoint},
		{name: "parent is not child", role: "own", dominion: "l2", child: patpFor(t, 1), wantErr: ErrInvalidPoint},
		{name: "invalid owner", role: "own", dominion: "l2", owner: "invalid"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var point types.Point
			point.Dominion = tc.dominion
			point.Ownership.TransferProxy.Address = tc.transferProxy
			switch tc.role {
			case "own":
				point.Ownership.Owner.Address = address
			case "spawn":
				point.Ownership.SpawnProxy.Address = address
			case "manage":
				point.Ownership.ManagementProxy.Address = address
			}
			var submitted *types.SpawnRequest
			withRPC(t, func(method string, params json.RawMessage) (any, error) {
				switch method {
				case "getPoint":
					return &point, nil
				case "getNonce":
					return 0, nil
				case "getUnsignedTx":
					return "0x" + strings.Repeat("0", 64), nil
				case "spawn":
					if err := json.Unmarshal(params, &submitted); err != nil {
						t.Error(err)
					}
					return "spawn-hash", nil
				default:
					t.Errorf("unexpected RPC %s", method)
					return nil, fmt.Errorf("unexpected method")
				}
			})
			child := tc.child
			if child == "" {
				child = patpFor(t, 257)
			}
			owner := tc.owner
			if owner == "" {
				owner = address
			}
			got, err := Spawn(context.Background(), patpFor(t, 1), keyHex, "", child, owner)
			if tc.wantErr != nil || tc.owner == "invalid" {
				if err == nil || (tc.wantErr != nil && !errors.Is(err, tc.wantErr)) {
					t.Fatalf("Spawn() error = %v", err)
				}
				if submitted != nil {
					t.Fatal("unauthorized spawn submitted")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if tx, ok := got.(types.Transaction); !ok || tx.Type != "spawn" || tx.Hash != "spawn-hash" {
				t.Fatalf("transaction = %#v", got)
			}
			if submitted == nil || submitted.From.Proxy != tc.role || submitted.Data.Ship != child || !strings.EqualFold(submitted.Data.Address, owner) || submitted.Signature == "" {
				t.Fatalf("spawn request = %#v", submitted)
			}
		})
	}
}

func TestPickRandomChildExcluding(t *testing.T) {
	for _, tc := range []struct {
		name           string
		parent         uint32
		fallback, full bool
	}{
		{"galaxy list", 1, false, false}, {"star list", 257, false, false},
		{"galaxy last available", 1, true, false}, {"star enumeration", 257, true, false},
		{"galaxy exhausted", 1, true, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			shift := uint(8)
			if tc.parent >= 256 {
				shift = 16
			}
			first := patpFor(t, uint32(1)<<shift|tc.parent)
			second := patpFor(t, uint32(2)<<shift|tc.parent)
			var spawned []string
			if tc.fallback && tc.parent < 256 {
				for i := 2; i < 256; i++ {
					if i == 173 && !tc.full {
						continue
					}
					spawned = append(spawned, strings.TrimPrefix(patpFor(t, uint32(i)<<shift|tc.parent), "~"))
				}
			}
			withRPC(t, func(method string, params json.RawMessage) (any, error) {
				switch method {
				case "getUnspawned":
					if tc.fallback {
						return nil, fmt.Errorf("unspawned list unavailable")
					}
					return []string{strings.TrimPrefix(first, "~"), strings.TrimPrefix(second, "~")}, nil
				case "getSpawned":
					return spawned, nil
				default:
					t.Errorf("unexpected RPC %s", method)
					return nil, fmt.Errorf("unexpected method")
				}
			})
			got, point, err := PickRandomChildExcluding(context.Background(), patpFor(t, tc.parent), []string{strings.TrimPrefix(first, "~")})
			if tc.full {
				if err == nil {
					t.Fatalf("selected %s from exhausted parent", got)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got == first || point>>shift == 0 || point&((1<<shift)-1) != tc.parent {
				t.Fatalf("invalid child %s (%d) of %d", got, point, tc.parent)
			}
			if !tc.fallback && got != second {
				t.Fatalf("child = %s, want %s", got, second)
			}
			if tc.fallback && tc.parent < 256 && point != 173<<8|tc.parent {
				t.Fatalf("did not select last child: %s", got)
			}
		})
	}
}

func TestSpawnRandomOwnerAcceptAndKeys(t *testing.T) {
	keyHex := strings.Repeat("1", 64)
	key, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		t.Fatal(err)
	}
	var parent, child types.Point
	parent.Dominion = "l2"
	parent.Ownership.Owner.Address = crypto.PubkeyToAddress(key.PublicKey).Hex()
	parentPatp, childPatp := patpFor(t, 257), patpFor(t, 65793)
	var methods []string
	var configured struct {
		Data types.ConfigureKeysParams `json:"data"`
	}
	withRPC(t, func(method string, params json.RawMessage) (any, error) {
		switch method {
		case "getUnspawned":
			return []uint32{65793}, nil
		case "getPoint":
			var req struct {
				Ship string `json:"ship"`
			}
			if err := json.Unmarshal(params, &req); err != nil {
				t.Error(err)
			}
			if req.Ship == parentPatp {
				return &parent, nil
			}
			return &child, nil
		case "getNonce":
			return 0, nil
		case "getUnsignedTx":
			return "0x" + strings.Repeat("0", 64), nil
		case "spawn":
			var req types.SpawnRequest
			if err := json.Unmarshal(params, &req); err != nil {
				t.Error(err)
			}
			child.Ownership.TransferProxy.Address = req.Data.Address
			methods = append(methods, method)
			return method + "-hash", nil
		case "transferPoint":
			child.Ownership.Owner.Address = child.Ownership.TransferProxy.Address
			methods = append(methods, method)
			return method + "-hash", nil
		case "configureKeys":
			if err := json.Unmarshal(params, &configured); err != nil {
				t.Error(err)
			}
			methods = append(methods, method)
			return method + "-hash", nil
		default:
			t.Errorf("unexpected RPC %s", method)
			return nil, fmt.Errorf("unexpected method")
		}
	})
	got, err := SpawnRandomOwner(context.Background(), parentPatp, keyHex, "", "", true, true)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(methods, []string{"spawn", "transferPoint", "configureKeys"}) {
		t.Fatalf("operations = %v", methods)
	}
	if got.Patp != childPatp || got.ParentPatp != parentPatp || got.Ship != 65793 || got.NetworkKeys == nil || got.AcceptTx == nil || got.KeysTx == nil {
		t.Fatalf("incomplete spawn result for %s", got.Patp)
	}
	if err := ticketValidation(got.Ticket); err != nil {
		t.Fatal(err)
	}
	wallet := keygen.GenerateWallet(strings.TrimPrefix(got.Ticket, "~"), got.Ship, "", 0, true)
	if !strings.EqualFold(got.Owner, wallet.Ownership.Keys.Address) {
		t.Fatal("ticket does not reproduce owner")
	}
	if got.NetworkKeys.Crypt.Public != wallet.Network.Keys.Crypt.Public || got.NetworkKeys.Auth.Public != wallet.Network.Keys.Auth.Public {
		t.Fatal("network keys do not match wallet")
	}
	want, err := roller.Keyfile(wallet.Network.Keys.Crypt.Private, wallet.Network.Keys.Auth.Private, childPatp, 1)
	if err != nil {
		t.Fatal(err)
	}
	if configured.Data.Encrypt != "0x"+got.NetworkKeys.Crypt.Public || configured.Data.Auth != "0x"+got.NetworkKeys.Auth.Public || configured.Data.Breach {
		t.Fatal("configured keys do not match returned keys")
	}
	if got.Keyfile != want {
		t.Fatal("incorrect child keyfile")
	}
}

func TestWaitTimestampAndCancellation(t *testing.T) {
	for _, pendingError := range []bool{false, true} {
		t.Run(fmt.Sprint(pendingError), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			withRPC(t, func(method string, _ json.RawMessage) (any, error) {
				switch method {
				case "whenNextBatch":
					return time.Now().Add(time.Hour).UnixMilli(), nil
				case "getAllPending":
					cancel()
					if pendingError {
						return nil, fmt.Errorf("unavailable")
					}
					var tx types.PendingTx
					tx.RawTx.Sig = "pending"
					return []types.PendingTx{tx}, nil
				default:
					t.Errorf("unexpected RPC %s", method)
					return nil, fmt.Errorf("unexpected method")
				}
			})
			if err := Wait(ctx, time.Hour, "pending"); !errors.Is(err, context.Canceled) {
				t.Fatalf("Wait() error = %v", err)
			}
		})
	}
}

func TestHex2bufRejectsInvalidHex(t *testing.T) {
	if _, err := hex2buf("not hex"); err == nil {
		t.Fatal("invalid hex accepted")
	}
}

func TestSpawnRandomOwnerRetainsTicketOnAcceptFailure(t *testing.T) {
	keyHex := strings.Repeat("1", 64)
	key, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		t.Fatal(err)
	}
	var parent types.Point
	parent.Dominion = "l2"
	parent.Ownership.Owner.Address = crypto.PubkeyToAddress(key.PublicKey).Hex()
	parentPatp := patpFor(t, 257)
	withRPC(t, func(method string, params json.RawMessage) (any, error) {
		switch method {
		case "getPoint":
			var req struct {
				Ship string `json:"ship"`
			}
			if err := json.Unmarshal(params, &req); err != nil {
				t.Error(err)
			}
			if req.Ship != parentPatp {
				return nil, fmt.Errorf("child lookup unavailable")
			}
			return &parent, nil
		case "getNonce":
			return 0, nil
		case "getUnsignedTx":
			return "0x" + strings.Repeat("0", 64), nil
		case "spawn":
			return "spawn-hash", nil
		default:
			t.Errorf("unexpected RPC %s", method)
			return nil, fmt.Errorf("unexpected method")
		}
	})
	got, err := SpawnRandomOwner(context.Background(), parentPatp, keyHex, "", patpFor(t, 65793), true, false)
	if err == nil || !strings.Contains(err.Error(), "accept transfer") {
		t.Fatalf("error = %v", err)
	}
	if got.Ticket == "" || got.Owner == "" || got.SpawnTx == nil {
		t.Fatal("recovery information missing after submitted spawn")
	}
	wallet := keygen.GenerateWallet(strings.TrimPrefix(got.Ticket, "~"), got.Ship, "", 0, true)
	if !strings.EqualFold(wallet.Ownership.Keys.Address, got.Owner) {
		t.Fatal("ticket cannot recover child")
	}
}
