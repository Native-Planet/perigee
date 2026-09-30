package libprg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	azimuth "github.com/Native-Planet/perigee/abi/azimuthcontract"
	"github.com/Native-Planet/perigee/abi/ecliptic"
	"github.com/Native-Planet/perigee/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// The RPC fixture decodes signed transactions and enforces the contract call
// arguments, so routing and permission checks are exercised without broadcasting.
func mockL1Spawn(t *testing.T, parent, child uint32, childOwner common.Address, permitted bool, receiptStatus uint64) *[]*ethTypes.Transaction {
	t.Helper()
	eclip, err := ecliptic.EclipticMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}
	az, err := azimuth.AzimuthMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}
	transactions := new([]*ethTypes.Transaction)
	withEthereumRPC(t, func(w http.ResponseWriter, req ethereumTestRequest) {
		var params []json.RawMessage
		if len(req.Params) > 0 {
			if err := json.Unmarshal(req.Params, &params); err != nil {
				t.Error(err)
				return
			}
		}
		respond := func(result any) {
			json.NewEncoder(w).Encode(map[string]any{"jsonrpc": "2.0", "id": req.ID, "result": result})
		}
		switch req.Method {
		case "eth_call":
			var call struct {
				Input hexutil.Bytes `json:"input"`
			}
			if err := json.Unmarshal(params[0], &call); err != nil {
				t.Error(err)
				return
			}
			data := call.Input
			if len(data) < 4 {
				t.Error("missing call data")
				return
			}
			var encoded []byte
			var err error
			switch {
			case bytes.Equal(data[:4], eclip.Methods["azimuth"].ID):
				encoded, err = eclip.Methods["azimuth"].Outputs.Pack(common.HexToAddress("0x1234"))
			case bytes.Equal(data[:4], az.Methods["getOwner"].ID):
				args, decodeErr := az.Methods["getOwner"].Inputs.Unpack(data[4:])
				if decodeErr != nil || args[0] != child {
					t.Errorf("getOwner args: %v, %v", args, decodeErr)
				}
				encoded, err = az.Methods["getOwner"].Outputs.Pack(childOwner)
			case bytes.Equal(data[:4], az.Methods["canSpawnAs"].ID):
				args, decodeErr := az.Methods["canSpawnAs"].Inputs.Unpack(data[4:])
				if decodeErr != nil || args[0] != parent {
					t.Errorf("canSpawnAs args: %v, %v", args, decodeErr)
				}
				encoded, err = az.Methods["canSpawnAs"].Outputs.Pack(permitted)
			default:
				t.Errorf("unexpected call: %x", data)
			}
			if err != nil {
				t.Error(err)
			}
			respond(hexutil.Encode(encoded))
		case "eth_sendRawTransaction":
			var raw hexutil.Bytes
			if err := json.Unmarshal(params[0], &raw); err != nil {
				t.Error(err)
				return
			}
			tx := new(ethTypes.Transaction)
			if err := tx.UnmarshalBinary(raw); err != nil {
				t.Error(err)
				return
			}
			if tx.To() == nil || *tx.To() != common.HexToAddress(EclipticContract) {
				t.Errorf("transaction destination: %v", tx.To())
			}
			*transactions = append(*transactions, tx)
			respond(tx.Hash())
		case "eth_getTransactionReceipt":
			var hash common.Hash
			json.Unmarshal(params[0], &hash)
			respond(map[string]any{"transactionHash": hash, "status": hexutil.EncodeUint64(receiptStatus), "cumulativeGasUsed": "0x1", "gasUsed": "0x1", "logsBloom": ethTypes.Bloom{}, "logs": []any{}})
		default:
			writeEthereumResult(t, w, req)
		}
	})
	return transactions
}

func TestL1SpawnRoutingAndPermissions(t *testing.T) {
	keyHex := strings.Repeat("1", 64)
	key, _ := crypto.HexToECDSA(keyHex)
	signer := crypto.PubkeyToAddress(key.PublicKey)
	for _, tc := range []struct {
		name, role, proxy string
		parent, child     uint32
		childOwner        common.Address
		denied            bool
	}{
		{name: "galaxy owner with zero proxy", role: "own", proxy: common.Address{}.Hex(), parent: 1, child: 257},
		{name: "star spawn proxy with zero proxy", role: "spawn", proxy: common.Address{}.Hex(), parent: 257, child: 65793},
		{name: "unset transfer proxy", role: "own", parent: 257, child: 65793},
		{name: "nonzero transfer proxy", role: "spawn", proxy: signer.Hex(), parent: 257, child: 65793},
		{name: "already allocated child", role: "own", parent: 257, child: 65793, childOwner: signer},
		{name: "on-chain permission denied", role: "own", parent: 257, child: 65793, denied: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var point types.Point
			point.Dominion = "l1"
			point.Ownership.TransferProxy.Address = tc.proxy
			if tc.role == "own" {
				point.Ownership.Owner.Address = signer.Hex()
			} else {
				point.Ownership.SpawnProxy.Address = signer.Hex()
			}
			withRPC(t, func(method string, _ json.RawMessage) (any, error) {
				if method != "getPoint" {
					t.Errorf("unexpected roller write: %s", method)
				}
				return &point, nil
			})
			txs := mockL1Spawn(t, tc.parent, tc.child, tc.childOwner, !tc.denied, 1)
			got, err := Spawn(context.Background(), patpFor(t, tc.parent), keyHex, "", patpFor(t, tc.child), signer.Hex())
			if tc.denied || tc.childOwner != (common.Address{}) {
				if err == nil || len(*txs) != 0 {
					t.Fatalf("rejected spawn: txs=%d error=%v", len(*txs), err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			tx, ok := got.(*ethTypes.Transaction)
			if !ok || len(*txs) != 1 {
				t.Fatalf("L1 transaction: %T, count=%d", got, len(*txs))
			}
			abi, _ := ecliptic.EclipticMetaData.GetAbi()
			want, _ := abi.Pack("spawn", tc.child, signer)
			if !bytes.Equal(tx.Data(), want) {
				t.Fatalf("spawn calldata = %x", tx.Data())
			}
			sender, err := ethTypes.Sender(ethTypes.LatestSignerForChainID(tx.ChainId()), tx)
			if err != nil || sender != signer {
				t.Fatalf("sender = %s, error=%v", sender, err)
			}
		})
	}
}

func TestL1SpawnRandomOwner(t *testing.T) {
	for _, tc := range []struct {
		name         string
		accept, keys bool
		status       uint64
		want         []string
		wantErr      string
	}{
		{"spawn only", false, false, 1, []string{"spawn"}, ""},
		{"accept", true, false, 1, []string{"spawn", "transferPoint"}, ""},
		{"keys imply acceptance", false, true, 1, []string{"spawn", "transferPoint", "configureKeys"}, ""},
		{"reverted spawn", true, true, 0, []string{"spawn"}, "reverted"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			keyHex := strings.Repeat("1", 64)
			key, _ := crypto.HexToECDSA(keyHex)
			var point types.Point
			point.Dominion = "l1"
			point.Ownership.Owner.Address = crypto.PubkeyToAddress(key.PublicKey).Hex()
			point.Ownership.TransferProxy.Address = common.Address{}.Hex()
			withRPC(t, func(method string, _ json.RawMessage) (any, error) {
				if method != "getPoint" {
					return nil, fmt.Errorf("unexpected roller method %s", method)
				}
				return &point, nil
			})
			txs := mockL1Spawn(t, 257, 65793, common.Address{}, true, tc.status)
			got, err := SpawnRandomOwner(context.Background(), patpFor(t, 257), keyHex, "", patpFor(t, 65793), tc.accept, tc.keys)
			if (tc.wantErr == "" && err != nil) || (tc.wantErr != "" && (err == nil || !strings.Contains(err.Error(), tc.wantErr))) {
				t.Fatalf("error = %v", err)
			}
			if got.Ticket == "" || got.SpawnTx == nil {
				t.Fatal("missing ticket or transaction")
			}
			abi, _ := ecliptic.EclipticMetaData.GetAbi()
			var methods []string
			for _, tx := range *txs {
				method, err := abi.MethodById(tx.Data()[:4])
				if err != nil {
					t.Fatal(err)
				}
				methods = append(methods, method.Name)
				args, err := method.Inputs.Unpack(tx.Data()[4:])
				if err != nil {
					t.Fatal(err)
				}
				if args[0] != uint32(65793) {
					t.Errorf("wrong child: %v", args[0])
				}
				if method.Name == "spawn" || method.Name == "transferPoint" {
					if args[1] != common.HexToAddress(got.Owner) {
						t.Error("wrong recipient")
					}
				}
				if method.Name != "spawn" {
					sender, err := ethTypes.Sender(ethTypes.LatestSignerForChainID(tx.ChainId()), tx)
					if err != nil || sender != common.HexToAddress(got.Owner) {
						t.Error("wrong child signer")
					}
				}
				if method.Name == "configureKeys" {
					crypt, _ := addHexPrefix(got.NetworkKeys.Crypt.Public)
					auth, _ := addHexPrefix(got.NetworkKeys.Auth.Public)
					if args[1] != crypt || args[2] != auth || args[3] != uint32(1) || args[4] != false {
						t.Error("wrong network key configuration")
					}
				}
			}
			if !reflect.DeepEqual(methods, tc.want) {
				t.Fatalf("calls = %v", methods)
			}
			if tc.keys && tc.wantErr == "" && (got.Keyfile == "" || got.KeysTx == nil || got.AcceptTx == nil) {
				t.Fatal("missing child setup result")
			}
		})
	}
}
