package libprg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Native-Planet/perigee/roller"
	"github.com/Native-Planet/perigee/types"
	"github.com/deelawn/urbit-gob/co"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nathanlever/keygen"
)

// GeneratePlanetCode creates a Bridge activation code without submitting a
// transaction. Persist the result before calling SpawnPlanetCode.
func GeneratePlanetCode(planet string) (types.PlanetCode, error) {
	name, point, err := types.ValidateAndNormalizePatp(planet)
	if err != nil || point < 65536 {
		return types.PlanetCode{}, ErrInvalidPoint
	}
	ticket, err := generateTicket(point)
	if err != nil {
		return types.PlanetCode{}, err
	}
	return PlanetCodeFromTicket(name, ticket)
}

// PlanetCodeFromTicket encodes a 64-bit invite secret and the planet name in
// Bridge's activation format. The secret belongs to the invite wallet; Bridge
// generates the recipient's master ticket independently during redemption.
func PlanetCodeFromTicket(planet, ticket string) (types.PlanetCode, error) {
	name, point, err := types.ValidateAndNormalizePatp(planet)
	if err != nil || point < 65536 {
		return types.PlanetCode{}, ErrInvalidPoint
	}
	ticket = "~" + strings.TrimPrefix(ticket, "~")
	if len(ticket) != 28 || !co.IsValidPatq(ticket) {
		return types.PlanetCode{}, errors.New("invalid planet invite secret")
	}
	entropy, err := co.Patq2Hex(ticket)
	if err != nil || len(entropy) != 16 {
		return types.PlanetCode{}, errors.New("planet invite requires 64 bits")
	}
	wallet := keygen.GenerateWallet(strings.TrimPrefix(ticket, "~"), point, "", 0, false)
	code := ticket + "-" + strings.TrimPrefix(name, "~")
	return types.PlanetCode{Planet: name, Code: code, Owner: wallet.Ownership.Keys.Address, URL: "https://bridge.urbit.org/#" + strings.TrimPrefix(code, "~")}, nil
}

func splitPlanetCode(code string) (planet, ticket string, point uint32, err error) {
	parts := strings.Split(strings.TrimPrefix(code, "~"), "-")
	if len(parts) != 6 {
		return "", "", 0, errors.New("invalid planet code")
	}
	planet, point, err = types.ValidateAndNormalizePatp("~" + strings.Join(parts[4:], "-"))
	if err != nil || point < 65536 {
		return "", "", 0, ErrInvalidPoint
	}
	ticket = "~" + strings.Join(parts[:4], "-")
	if len(ticket) != 28 || !co.IsValidPatq(ticket) {
		return "", "", 0, errors.New("invalid planet code")
	}
	return
}

// ParsePlanetCode derives the invite ownership address used by Bridge.
func ParsePlanetCode(code string) (types.PlanetCode, error) {
	planet, ticket, _, err := splitPlanetCode(code)
	if err != nil {
		return types.PlanetCode{}, err
	}
	return PlanetCodeFromTicket(planet, ticket)
}

// SpawnPlanetCode submits an L2 spawn to the persisted code's invite address.
// Reconcile an uncertain submission before retrying. AcceptPlanetCode completes
// the invite ownership handoff so Bridge can redeem it.
func SpawnPlanetCode(ctx context.Context, parent, parentKey, passphrase, code string) (*types.Transaction, error) {
	invite, err := ParsePlanetCode(code)
	if err != nil {
		return nil, err
	}
	_, parentPoint, err := types.ValidateAndNormalizePatp(parent)
	if err != nil || parentPoint < 256 || parentPoint >= 65536 {
		return nil, ErrInvalidPoint
	}
	p, err := roller.Client.GetPoint(ctx, parent)
	if err != nil {
		return nil, err
	}
	if p.Dominion != "l2" && p.Dominion != "spawn" {
		return nil, errors.New("planet codes require an L2-spawning star")
	}
	result, err := Spawn(ctx, parent, parentKey, passphrase, invite.Planet, invite.Owner)
	if err != nil {
		return nil, err
	}
	tx, ok := result.(types.Transaction)
	if !ok {
		return nil, errors.New("planet code spawn returned an unexpected transaction")
	}
	return &tx, nil
}

// AcceptPlanetCode moves the spawned point from the parent's ownership to the
// invite wallet. It configures no network keys. A nil transaction means the
// invite wallet already owns the point.
func AcceptPlanetCode(ctx context.Context, code string) (*types.Transaction, error) {
	planet, ticket, point, err := splitPlanetCode(code)
	if err != nil {
		return nil, err
	}
	wallet := keygen.GenerateWallet(strings.TrimPrefix(ticket, "~"), point, "", 0, false)
	owner := wallet.Ownership.Keys.Address
	p, err := roller.Client.GetPoint(ctx, planet)
	if err != nil {
		return nil, err
	}
	if p.Dominion != "l2" {
		return nil, errors.New("planet code requires an L2 planet")
	}
	if strings.EqualFold(p.Ownership.Owner.Address, owner) {
		return nil, nil
	}
	if !strings.EqualFold(p.Ownership.TransferProxy.Address, owner) {
		return nil, errors.New("planet code has no pending transfer")
	}
	key, err := crypto.HexToECDSA(wallet.Ownership.Keys.Private)
	if err != nil {
		return nil, fmt.Errorf("derive invite key: %w", err)
	}
	return roller.Client.TransferPoint(ctx, planet, true, owner, owner, key)
}
