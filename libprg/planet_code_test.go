package libprg

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/Native-Planet/perigee/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// These vectors use Bridge's urbit-key-generation@0.19.0 and urbit-ob@4.1.4.
func TestPlanetCodeBridgeVectors(t *testing.T) {
	b, err := os.ReadFile("testdata/bridge_planet_codes.json")
	if err != nil {
		t.Fatal(err)
	}
	var vectors []struct{ Planet, Ticket, Code, URL, Owner string }
	if err = json.Unmarshal(b, &vectors); err != nil {
		t.Fatal(err)
	}
	for _, v := range vectors {
		t.Run(v.Planet, func(t *testing.T) {
			code, err := PlanetCodeFromTicket(v.Planet, v.Ticket)
			if err != nil {
				t.Fatal(err)
			}
			if code.Planet != v.Planet || code.Code != v.Code || code.URL != v.URL || !strings.EqualFold(code.Owner, v.Owner) {
				t.Fatalf("Bridge mismatch: %+v", code)
			}
			parsed, err := ParsePlanetCode(strings.TrimPrefix(code.Code, "~"))
			if err != nil || parsed != code {
				t.Fatal("code round trip", err)
			}
		})
	}
	for _, code := range []string{"", "~doznec-binwes-samper-siglet", "~doznec-binwes-samper-siglet-zod", "~doznec-binwes-samper-siglet-marzod", "~badbad-binwes-samper-siglet-sampel-palnet"} {
		if _, err := ParsePlanetCode(code); err == nil {
			t.Fatal("invalid code accepted")
		}
	}
	if _, err := GeneratePlanetCode("~zod"); err == nil {
		t.Fatal("nonplanet accepted")
	}
}

func TestPlanetCodeSpawnAndAcceptance(t *testing.T) {
	const code = "~doznec-binwes-samper-siglet-sampel-palnet"
	const inviteOwner = "0x5EfED9f04129342e83c92dd8eB26332b05afb989"
	invite, err := ParsePlanetCode(code)
	if err != nil {
		t.Fatal(err)
	}
	_, point, err := types.ValidateAndNormalizePatp(invite.Planet)
	if err != nil {
		t.Fatal(err)
	}
	parent := patpFor(t, point&65535)
	keyHex := strings.Repeat("1", 64)
	key, _ := crypto.HexToECDSA(keyHex)
	signer := crypto.PubkeyToAddress(key.PublicKey).Hex()
	p := types.Point{Dominion: "l2"}
	p.Ownership.SpawnProxy.Address = signer
	child := types.Point{Dominion: "l2"}
	counts := map[string]int{}
	withRPC(t, func(method string, params json.RawMessage) (any, error) {
		counts[method]++
		switch method {
		case "getPoint":
			var req struct{ Ship string }
			json.Unmarshal(params, &req)
			if req.Ship == parent {
				return p, nil
			}
			return child, nil
		case "getNonce":
			return 0, nil
		case "getUnsignedTx":
			return "0x" + strings.Repeat("0", 64), nil
		case "spawn":
			var req types.SpawnRequest
			json.Unmarshal(params, &req)
			if req.From.Proxy != "spawn" || req.Data.Ship != invite.Planet || !strings.EqualFold(req.Data.Address, inviteOwner) {
				t.Error("incorrect invite spawn")
			}
			child.Ownership.TransferProxy.Address = req.Data.Address
			return "spawn-hash", nil
		case "transferPoint":
			var req types.TransferRequest
			json.Unmarshal(params, &req)
			if req.From.Proxy != "transfer" || !req.Data.Reset || !strings.EqualFold(req.Data.Address, inviteOwner) {
				t.Error("incorrect invite acceptance")
			}
			child.Ownership.Owner.Address = req.Data.Address
			child.Ownership.TransferProxy.Address = ""
			return "accept-hash", nil
		default:
			t.Errorf("unexpected method %s", method)
			return nil, errors.New("unexpected method")
		}
	})
	if tx, err := SpawnPlanetCode(context.Background(), parent, keyHex, "", code); err != nil || tx.Hash != "spawn-hash" {
		t.Fatal(tx, err)
	}
	if tx, err := AcceptPlanetCode(context.Background(), code); err != nil || tx.Hash != "accept-hash" {
		t.Fatal(tx, err)
	}
	if tx, err := AcceptPlanetCode(context.Background(), code); err != nil || tx != nil {
		t.Fatal("owned invite submitted again", tx, err)
	}
	p.Dominion = "l1"
	if _, err := SpawnPlanetCode(context.Background(), parent, keyHex, "", code); err == nil {
		t.Fatal("L1 code accepted")
	}
	if counts["spawn"] != 1 || counts["transferPoint"] != 1 || counts["configureKeys"] != 0 {
		t.Fatal(counts)
	}
}
