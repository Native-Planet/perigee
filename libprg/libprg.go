package libprg

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Native-Planet/perigee/abi/ecliptic"
	"github.com/Native-Planet/perigee/roller"
	"github.com/Native-Planet/perigee/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nathanlever/keygen"
)

const (
	Version                     = "0.2.0"
	CRYPTO_SUITE_VERSION uint32 = 1
	EclipticContract            = "0x33EeCbf908478C10614626A9D304bfe18B78DD73"
)

var (
	EthProvider        = os.Getenv("ETH_PROVIDER")
	ErrInvalidPoint    = errors.New("invalid point")
	ErrInvalidLife     = errors.New("invalid life value")
	ErrKeyMismatch     = errors.New("public key mismatch with PKI")
	ErrKeyMaterial     = errors.New("invalid key material")
	ErrRollerOperation = errors.New("roller operation failed")
	ctx                = context.Background()
)

type EclipticSession struct {
	Contract     *ecliptic.Ecliptic
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

func GetContext() context.Context {
	return ctx
}

/*
Note: for passphrases and life values, just set default values if
you don't have a specific reason to use them
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

func Breach(point, ticket, passphrase string) (interface{}, error) {
	wallet, pointInfo, patp, err := getWalletAndPoint(point, ticket, passphrase, 0, true)
	if err != nil {
		return types.Transaction{}, err
	}
	if pointInfo.Dominion == "l2" {
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
	} else {
		tx, err := setL1NetworkKeys(point, pointInfo, wallet, true)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
		}
		ethClient, err := ethclient.Dial(EthProvider)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("failed to connect to Ethereum node: %v", err)
		}
		receipt, err := waitForReceipt(ctx, ethClient, tx)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
		}
		return receipt, nil
	}
}

func Pending(addr string) ([]types.PendingTx, error) {
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

func Wallet(point, masterTicket, passphrase string, life int) (keygen.Wallet, error) {
	wallet, _, _, err := getWalletAndPoint(point, masterTicket, passphrase, life, true)
	if err != nil {
		return keygen.Wallet{}, err
	}
	return wallet, nil
}

func Keyfile(point, masterTicket, passphrase string, life int) (string, error) {
	wallet, pInfo, patp, err := getWalletAndPoint(point, masterTicket, passphrase, life, true)
	if err != nil {
		return "", err
	}
	rev, err := strconv.Atoi(pInfo.Network.Keys.Life)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrInvalidLife, err)
	}
	keyfile, err := roller.Keyfile(
		wallet.Network.Keys.Crypt.Private,
		wallet.Network.Keys.Auth.Private,
		patp,
		rev,
	)
	if err != nil {
		return "", fmt.Errorf("generating keyfile: %v", err)
	}
	return keyfile, nil
}

func setL1NetworkKeys(patp string, point *types.Point, wallet keygen.Wallet, breach bool) (*ethTypes.Transaction, error) {
	patp, pointInt, err := types.ValidateAndNormalizePatp(patp)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	if EthProvider == "" {
		return &ethTypes.Transaction{}, fmt.Errorf("must set ETH_PROVIDER (ex: https://mainnet.infura.io/v3/<your-key>)")
	}
	client, err := ethclient.Dial(EthProvider)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("failed to connect to Ethereum: %v", err)
	}
	ctx := context.Background()
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("failed to get network ID: %v", err)
	}
	privateKey, err := crypto.HexToECDSA(wallet.Ownership.Keys.Private)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("failed to get private key: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("failed to create transactor: %v", err)
	}
	if auth.GasPrice == nil || auth.GasPrice.Cmp(big.NewInt(0)) == 0 {
		gasPrice, err := fetchGasPrice(ctx, client)
		if err != nil {
			return &ethTypes.Transaction{}, fmt.Errorf("failed to fetch gas price: %v", err)
		}
		auth.GasPrice = gasPrice
	}
	eclipticAddress := common.HexToAddress(EclipticContract)
	eclip, err := ecliptic.NewEcliptic(eclipticAddress, client)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("failed to bind Ecliptic contract: %v", err)
	}
	publicCrypt, err := addHexPrefix(wallet.Network.Keys.Crypt.Public)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("invalid crypt key for %s: %v", patp, err)
	}
	publicAuth, err := addHexPrefix(wallet.Network.Keys.Auth.Public)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("invalid auth key for %s: %v", patp, err)
	}
	// Check if keys are already set
	if strings.EqualFold(point.Network.Keys.Crypt, hex.EncodeToString(publicCrypt[:])) &&
		strings.EqualFold(point.Network.Keys.Auth, hex.EncodeToString(publicAuth[:])) &&
		!breach {
		return &ethTypes.Transaction{}, fmt.Errorf("the network key is already set for %s", patp)
	} else if strings.EqualFold(point.Network.Keys.Crypt, hex.EncodeToString(publicCrypt[:])) &&
		strings.EqualFold(point.Network.Keys.Auth, hex.EncodeToString(publicAuth[:])) &&
		breach {
		fmt.Printf("Breaching with existing key revision for %s.", patp)
	}
	session := &ecliptic.EclipticSession{
		Contract:     eclip,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From: auth.From,
		},
	}
	return session.ConfigureKeys(pointInt, publicCrypt, publicAuth, CRYPTO_SUITE_VERSION, breach)
}

func Wait(duration time.Duration, keysTx string) error {
	if duration == 0 {
		return nil
	}
	batchInfo, err := roller.Client.WhenNextBatch(ctx)
	if err != nil {
		return fmt.Errorf("getting batch info: %v", err)
	}
	if batchInfo == nil {
		return fmt.Errorf("no batch info")
	}
	waitTime := time.Duration(batchInfo.TimeUntilNext) * time.Second
	if waitTime > duration {
		waitTime = duration
	}
	deadline := time.Now().Add(duration)
	ticker := time.NewTicker(waitTime)
	defer ticker.Stop()

	for {
		pending, err := Pending("")
		if err != nil {
			return fmt.Errorf("error getting pending ships: %v", err)
		}
		found := false
		for _, tx := range pending {
			if tx.RawTx.Sig == keysTx {
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Transaction committed to chain")
			return nil
		}
		select {
		case <-ticker.C:
		case <-time.After(time.Until(deadline)):
			return fmt.Errorf("timeout waiting for transaction after %v", duration)
		}
	}
}

func addHexPrefix(s string) ([32]byte, error) {
	var arr [32]byte
	s = strings.TrimPrefix(s, "0x")
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return arr, err
	}
	if len(bytes) != 32 {
		return arr, errors.New("invalid length for key, must be 32 bytes")
	}
	copy(arr[:], bytes)
	return arr, nil
}

func waitForReceipt(ctx context.Context, client *ethclient.Client, tx *ethTypes.Transaction) (*ethTypes.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			return receipt, nil
		}
		time.Sleep(2 * time.Second)
	}
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
	// 0 life param means look it up
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

// wallet generation
// a note about adjustLife: keyfiles are generated using the private keys
// of the *previous* life value, so we need to generate the previous life's
// keys and then generate the keyfile noun using the current life value
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

func fetchGasPrice(ctx context.Context, client *ethclient.Client) (*big.Int, error) {
	return client.SuggestGasPrice(ctx)
}
