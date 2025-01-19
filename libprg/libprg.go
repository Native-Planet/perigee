package librpg

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Native-Planet/perigee/roller"
	"github.com/Native-Planet/perigee/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nathanlever/keygen"
)

const Version = "0.1.0"

var (
	ErrInvalidPoint    = errors.New("invalid point")
	ErrInvalidLife     = errors.New("invalid life value")
	ErrKeyMismatch     = errors.New("public key mismatch with PKI")
	ErrKeyMaterial     = errors.New("invalid key material")
	ErrRollerOperation = errors.New("roller operation failed")
	ctx                = context.Background()
)

func GetContext() context.Context {
	return ctx
}

/*
Note: for passphrases and life values, just set default values if
you don't have a specific reason to use them
Just use adjustLife if you're generating keyfiles or breaching
*/

func Escape(point, masterTicket, passphrase, sponsor string) (types.Transaction, error) {
	return handleTransaction(point, masterTicket, passphrase, sponsor, roller.Client.Escape)
}

func CancelEscape(point, masterTicket, passphrase, sponsor string) (types.Transaction, error) {
	return handleTransaction(point, masterTicket, passphrase, sponsor, roller.Client.CancelEscape)
}

func Adopt(point, masterTicket, passphrase, adoptee string) (types.Transaction, error) {
	return handleTransaction(point, masterTicket, passphrase, adoptee, roller.Client.Adopt)
}

func Breach(point, ticket, passphrase string, life int) (types.Transaction, error) {
	wallet, _, patp, err := getWalletAndPoint(point, ticket, passphrase, life, false)
	if err != nil {
		return types.Transaction{}, err
	}
	privKey, err := crypto.HexToECDSA(wallet.Ownership.Keys.Private)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	keysTx, err := roller.Client.ConfigureKeys(ctx, patp,
		"0x"+wallet.Network.Keys.Crypt.Public,
		"0x"+wallet.Network.Keys.Auth.Public,
		true, wallet.Ownership.Keys.Address, privKey)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrRollerOperation, err)
	}
	return *keysTx, nil
}

func RollerPending(addr string) ([]types.PendingTx, error) {
	if addr == "" {
		return roller.Client.GetAllPending(ctx)
	}
	if !strings.HasPrefix(addr, "0x") {
		_, pInfo, err := validatePointAndGetInfo(addr)
		if err != nil {
			return nil, err
		}
		addr = pInfo.Ownership.Owner.Address
	}
	return roller.Client.GetPendingByAddress(ctx, addr)
}

func Point(point string) (types.PointResp, error) {
	patp, pInfo, err := validatePointAndGetInfo(point)
	if err != nil {
		return types.PointResp{}, err
	}
	resp := types.PointResp{
		Point:    pInfo,
		PatpName: patp,
	}
	if err := resp.Point.ResolveSponsorPatp(); err != nil {
		return types.PointResp{}, fmt.Errorf("invalid sponsor point: %v", err)
	}
	return resp, nil
}

func Keyfile(point, masterTicket, passphrase, life string) (string, error) {
	lifeInt := 0
	if life != "" {
		var err error
		lifeInt, err = strconv.Atoi(life)
		if err != nil {
			return "", fmt.Errorf("%w: %v", ErrInvalidLife, err)
		}
	}
	wallet, _, patp, err := getWalletAndPoint(point, masterTicket, passphrase, lifeInt, true)
	if err != nil {
		return "", err
	}
	keyfile, err := roller.Keyfile(
		wallet.Network.Keys.Crypt.Private,
		wallet.Network.Keys.Auth.Private,
		patp,
		lifeInt,
	)
	if err != nil {
		return "", fmt.Errorf("generating keyfile: %v", err)
	}
	return keyfile, nil
}

func validatePointAndGetInfo(point string) (string, *types.Point, error) {
	patp, _, err := types.ValidateAndNormalizePatp(point)
	if err != nil {
		return "", nil, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	pInfo, err := roller.Client.GetPoint(ctx, patp)
	if err != nil {
		return "", nil, fmt.Errorf("getting point info: %v", err)
	}
	return patp, pInfo, nil
}

// generic transaction handler
func handleTransaction(point, masterTicket, passphrase, target string,
	operation func(context.Context, string, string, string, *ecdsa.PrivateKey) (*types.Transaction, error)) (types.Transaction, error) {
	masterTicket = strings.TrimPrefix(masterTicket, "~")
	wallet, _, patp, err := getWalletAndPoint(point, masterTicket, passphrase, 0, true) // true for life adjustment
	if err != nil {
		return types.Transaction{}, err
	}
	privKey, err := crypto.HexToECDSA(wallet.Ownership.Keys.Private)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	tx, err := operation(ctx, patp, target, wallet.Ownership.Keys.Address, privKey)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrRollerOperation, err)
	}
	return *tx, nil
}

// wallet generation (used downstream or directly)
func getWalletAndPoint(point, masterTicket, passphrase string, life int, adjustLife bool) (keygen.Wallet, *types.Point, string, error) {
	masterTicket = strings.TrimPrefix(masterTicket, "~")
	patp, pointInt, err := types.ValidateAndNormalizePatp(point)
	if err != nil {
		return keygen.Wallet{}, nil, "", fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	pInfo, err := roller.Client.GetPoint(ctx, patp)
	if err != nil {
		return keygen.Wallet{}, nil, "", fmt.Errorf("getting point info: %v", err)
	}
	var rev int
	if life == 0 {
		rev, err = strconv.Atoi(pInfo.Network.Keys.Life)
		if err != nil {
			return keygen.Wallet{}, nil, "", fmt.Errorf("%w: %v", ErrInvalidLife, err)
		}
	} else {
		rev = life
	}
	walletLife := rev
	if adjustLife {
		walletLife -= 1
	}
	wallet := keygen.GenerateWallet(masterTicket, uint32(pointInt), passphrase, uint(walletLife), true)
	pointKey := strings.TrimPrefix(pInfo.Network.Keys.Crypt, "0x")
	if wallet.Network.Keys.Crypt.Public != pointKey {
		return keygen.Wallet{}, nil, "", fmt.Errorf("%w: expected 0x%s, got %s",
			ErrKeyMismatch, wallet.Network.Keys.Crypt.Public, pInfo.Network.Keys.Crypt)
	}
	return wallet, pInfo, patp, nil
}
