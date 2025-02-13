package libprg

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/sha512"
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

	"github.com/deelawn/urbit-gob/co"
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

func Escape(point, masterTicket, passphrase, sponsor string) (interface{}, error) {
	privateKey, _, pointInfo, _, _, err := ValidateKey(point, masterTicket, passphrase, "", false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return l2Transaction(point, masterTicket, passphrase, sponsor, roller.Client.Escape)
	} else {
		return l1Escape(point, sponsor, privateKey)
	}
}

func EscapeRawTx(address, pointStr, sponsorStr string) ([]byte, error) {
	pInfo, err := Point(pointStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	point, err := co.Patp2Dec(pInfo.PatpName)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	pointInt, err := strconv.Atoi(point)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	sponsorPatp, sponsor, err := types.ValidateAndNormalizePatp(sponsorStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	if pInfo.Point.Dominion == "l1" {
		ecliptic, err := Ecliptic()
		if err != nil {
			return nil, err
		}
		client, err := ethclient.Dial(EthProvider)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to Ethereum: %v", err)
		}
		addr := common.HexToAddress(address)
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to get gas price: %v", err)
		}
		auth := &bind.TransactOpts{
			From:     addr,
			GasPrice: gasPrice,
			GasLimit: 300000,
			NoSend:   true,
		}
		tx, err := ecliptic.Escape(auth, uint32(pointInt), uint32(sponsor))
		if err != nil {
			return nil, err
		}
		return tx.Data(), nil
	} else {
		ctx := context.Background()
		proxy, err := roller.Client.GetManagementProxyType(ctx, pInfo.PatpName, address)
		if err != nil {
			return nil, err
		}
		params := types.EscapeRequest{
			Address: address,
			From: types.FromData{
				Ship:  pInfo.PatpName,
				Proxy: proxy,
			},
			Data: struct {
				Ship string `json:"ship"`
			}{
				Ship: sponsorPatp,
			},
		}
		hash, _, err := roller.Client.GetRawTx(ctx, "escape", params)
		if err != nil {
			return nil, fmt.Errorf("get unsigned tx: %v", err)
		}
		cleanHash := strings.TrimPrefix(hash, "0x")
		return hex.DecodeString(cleanHash)
	}
}

func CancelEscape(point, masterTicket, passphrase, sponsor string) (interface{}, error) {
	privateKey, _, pointInfo, _, _, err := ValidateKey(point, masterTicket, passphrase, "", false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return l2Transaction(point, masterTicket, passphrase, sponsor, roller.Client.CancelEscape)
	} else {
		return l1CancelEscape(point, privateKey)
	}
}

func Adopt(point, masterTicket, passphrase, adoptee string) (interface{}, error) {
	privateKey, _, pointInfo, _, _, err := ValidateKey(point, masterTicket, passphrase, "", false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return l2Transaction(point, masterTicket, passphrase, adoptee, roller.Client.Adopt)
	} else {
		return l1Adopt(adoptee, privateKey)
	}
}

func Breach(point, ticket, passphrase, seed string) (interface{}, error) {
	var wallet keygen.Wallet
	var pointInfo *types.Point
	var patp string
	var err error
	var cryptPubkey string
	var authPubkey string
	if seed == "" {
		wallet, pointInfo, patp, err = getWalletAndPoint(point, ticket, passphrase, 0, false)
		if err != nil {
			return types.Transaction{}, err
		}
		cryptPubkey = wallet.Network.Keys.Crypt.Public
		authPubkey = wallet.Network.Keys.Auth.Public
	} else {
		// if providing a seed, generate keys
		pointRes, err := Point(point)
		if err != nil {
			return types.Transaction{}, err
		}
		pointInfo = pointRes.Point
	}
	if pointInfo.Dominion == "l2" {
		privKey, derivedPubkey, _, networkKeys, _, err := ValidateKey(point, ticket, passphrase, seed, true)
		cryptPubkey = networkKeys.Crypt.Public
		authPubkey = networkKeys.Auth.Public
		if err != nil {
			return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
		}
		keysTx, err := roller.Client.ConfigureKeys(ctx, patp,
			"0x"+cryptPubkey,
			"0x"+authPubkey,
			true, derivedPubkey, privKey)
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

func TransferOwnership(point, masterTicket, passphrase, newOwner string, reset bool) (interface{}, error) {
	privateKey, derivedPubkey, pointInfo, _, _, err := ValidateKey(point, masterTicket, passphrase, "", false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		tx, err := roller.Client.TransferPoint(
			context.Background(),
			point,
			reset,
			newOwner,
			derivedPubkey,
			privateKey,
		)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("%w: %v", ErrRollerOperation, err)
		}
		return *tx, nil
	} else {
		// Handle L1 logic
		return l1TransferOwnership(common.HexToAddress(newOwner), privateKey)
	}
}

func SetManagementProxy(point, masterTicket, passphrase, proxy string) (interface{}, error) {
	privateKey, _, pointInfo, _, _, err := ValidateKey(point, masterTicket, passphrase, "", false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return l2Transaction(point, masterTicket, passphrase, proxy, roller.Client.SetManagementProxy)
	} else {
		return l1SetManagementProxy(point, common.HexToAddress(proxy), privateKey)
	}
}

func SetTransferProxy(point, masterTicket, passphrase, proxy string) (interface{}, error) {
	privateKey, _, pointInfo, _, _, err := ValidateKey(point, masterTicket, passphrase, "", false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return l2Transaction(point, masterTicket, passphrase, proxy, roller.Client.SetTransferProxy)
	} else {
		return l1SetTransferProxy(point, common.HexToAddress(proxy), privateKey)
	}
}

func SetSpawnProxy(point, masterTicket, passphrase, proxy string) (interface{}, error) {
	privateKey, _, pointInfo, _, _, err := ValidateKey(point, masterTicket, passphrase, "", false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return l2Transaction(point, masterTicket, passphrase, proxy, roller.Client.SetSpawnProxy)
	} else {
		return l1SetSpawnProxy(point, common.HexToAddress(proxy), privateKey)
	}
}

func SetVotingProxy(galaxy, masterTicket, passphrase, voter string) (interface{}, error) {
	privateKey, _, pointInfo, _, _, err := ValidateKey(galaxy, masterTicket, passphrase, "", false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return types.Transaction{}, fmt.Errorf("how did you even submit this transaction")
	} else {
		return l1SetVotingProxy(galaxy, common.HexToAddress(voter), privateKey)
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
	if err := resp.ResolveClan(); err != nil {
		return types.PointResp{}, fmt.Errorf("invalid point clan: %v", err)
	}
	if err := resp.ResolveSein(); err != nil {
		return types.PointResp{}, fmt.Errorf("invalid point sein: %v", err)
	}
	return resp, nil
}

func Wallet(point, masterTicket, passphrase string, life int) (keygen.Wallet, error) {
	wallet, _, _, err := getWalletAndPoint(point, masterTicket, passphrase, life, false)
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

func SubmitSignedL1Tx(signedTxHex string) (*ethTypes.Transaction, error) {
	client, err := ethclient.Dial(EthProvider)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("failed to connect to Ethereum: %v", err)
	}
	signedTxData := common.FromHex(signedTxHex)
	var tx ethTypes.Transaction
	if err := tx.UnmarshalBinary(signedTxData); err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("unmarshal tx: %w", err)
	}
	return &tx, client.SendTransaction(context.Background(), &tx)
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
	if err := getGas(ctx, client, auth, 5); err != nil {
		return &ethTypes.Transaction{}, err
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

func l1Escape(patp string, sponsor string, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	_, pointInt, err := types.ValidateAndNormalizePatp(patp)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	_, sponsorInt, err := types.ValidateAndNormalizePatp(sponsor)
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
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("failed to create transactor: %v", err)
	}
	if err := getGas(ctx, client, auth, 5); err != nil {
		return &ethTypes.Transaction{}, err
	}
	eclipticAddress := common.HexToAddress(EclipticContract)
	eclip, err := ecliptic.NewEcliptic(eclipticAddress, client)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("failed to bind Ecliptic contract: %v", err)
	}
	session := &ecliptic.EclipticSession{
		Contract:     eclip,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From: auth.From,
		},
	}
	allowed, err := session.CanEscapeTo(pointInt, sponsorInt)
	if allowed {
		return session.Escape(pointInt, sponsorInt)
	} else {
		return &ethTypes.Transaction{}, fmt.Errorf("illegal escape: %v", err)
	}
}

func l1Adopt(adoptee string, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	_, adopteeInt, err := types.ValidateAndNormalizePatp(adoptee)
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
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("failed to create transactor: %v", err)
	}
	if err := getGas(ctx, client, auth, 5); err != nil {
		return &ethTypes.Transaction{}, err
	}
	eclipticAddress := common.HexToAddress(EclipticContract)
	eclip, err := ecliptic.NewEcliptic(eclipticAddress, client)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("failed to bind Ecliptic contract: %v", err)
	}
	session := &ecliptic.EclipticSession{
		Contract:     eclip,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From: auth.From,
		},
	}
	return session.Adopt(adopteeInt)
}

func Ecliptic() (*ecliptic.Ecliptic, error) {
	client, err := ethclient.Dial(EthProvider)
	if err != nil {
		return &ecliptic.Ecliptic{}, fmt.Errorf("failed to connect to Ethereum: %v", err)
	}
	eclip, err := ecliptic.NewEcliptic(common.HexToAddress(EclipticContract), client)
	if err != nil {
		return &ecliptic.Ecliptic{}, fmt.Errorf("failed to bind Ecliptic contract: %v", err)
	}
	return eclip, nil
}

func l1CancelEscape(point string, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	_, pointInt, err := types.ValidateAndNormalizePatp(point)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	if EthProvider == "" {
		return nil, fmt.Errorf("must set ETH_PROVIDER (ex: https://mainnet.infura.io/v3/<your-key>)")
	}
	client, err := ethclient.Dial(EthProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum: %v", err)
	}
	ctx := context.Background()
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}
	if err := getGas(ctx, client, auth, 5); err != nil {
		return nil, err
	}
	eclipticAddress := common.HexToAddress(EclipticContract)
	eclip, err := ecliptic.NewEcliptic(eclipticAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind Ecliptic contract: %v", err)
	}
	session := &ecliptic.EclipticSession{
		Contract:     eclip,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From: auth.From,
		},
	}
	return session.CancelEscape(pointInt)
}

// Transfer Ownership (L1)
func l1TransferOwnership(newOwner common.Address, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	if EthProvider == "" {
		return nil, fmt.Errorf("must set ETH_PROVIDER")
	}
	client, err := ethclient.Dial(EthProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum: %v", err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}
	if err := getGas(context.Background(), client, auth, 5); err != nil {
		return nil, err
	}
	eclipticAddress := common.HexToAddress(EclipticContract)
	eclip, err := ecliptic.NewEcliptic(eclipticAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind Ecliptic contract: %v", err)
	}
	session := &ecliptic.EclipticSession{
		Contract:     eclip,
		TransactOpts: *auth,
	}
	return session.TransferOwnership(newOwner)
}

// Set Management Proxy (L1)
func l1SetManagementProxy(pointStr string, manager common.Address, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	_, point, err := types.ValidateAndNormalizePatp(pointStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	client, err := ethclient.Dial(EthProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum: %v", err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}
	if err := getGas(context.Background(), client, auth, 5); err != nil {
		return nil, err
	}
	eclipticAddress := common.HexToAddress(EclipticContract)
	eclip, err := ecliptic.NewEcliptic(eclipticAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind Ecliptic contract: %v", err)
	}
	session := &ecliptic.EclipticSession{
		Contract:     eclip,
		TransactOpts: *auth,
	}
	return session.SetManagementProxy(uint32(point), manager)
}

// Set Transfer Proxy (L1)
func l1SetTransferProxy(pointStr string, proxy common.Address, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	_, point, err := types.ValidateAndNormalizePatp(pointStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	client, err := ethclient.Dial(EthProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum: %v", err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}
	if err := getGas(context.Background(), client, auth, 5); err != nil {
		return nil, err
	}
	eclipticAddress := common.HexToAddress(EclipticContract)
	eclip, err := ecliptic.NewEcliptic(eclipticAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind Ecliptic contract: %v", err)
	}
	session := &ecliptic.EclipticSession{
		Contract:     eclip,
		TransactOpts: *auth,
	}
	return session.SetTransferProxy(uint32(point), proxy)
}

// Set Spawn Proxy (L1)
func l1SetSpawnProxy(prefixStr string, proxy common.Address, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	prefix, err := strconv.ParseUint(prefixStr, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("invalid prefix: %v", err)
	}
	client, err := ethclient.Dial(EthProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum: %v", err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}
	if err := getGas(context.Background(), client, auth, 5); err != nil {
		return nil, err
	}
	eclipticAddress := common.HexToAddress(EclipticContract)
	eclip, err := ecliptic.NewEcliptic(eclipticAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind Ecliptic contract: %v", err)
	}
	session := &ecliptic.EclipticSession{
		Contract:     eclip,
		TransactOpts: *auth,
	}
	return session.SetSpawnProxy(uint16(prefix), proxy)
}

// Set Voting Proxy (L1)
func l1SetVotingProxy(galaxyStr string, voter common.Address, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	galaxy, err := strconv.ParseUint(galaxyStr, 10, 8)
	if err != nil {
		return nil, fmt.Errorf("invalid galaxy: %v", err)
	}
	client, err := ethclient.Dial(EthProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum: %v", err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}
	if err := getGas(context.Background(), client, auth, 5); err != nil {
		return nil, err
	}
	eclipticAddress := common.HexToAddress(EclipticContract)
	eclip, err := ecliptic.NewEcliptic(eclipticAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind Ecliptic contract: %v", err)
	}
	session := &ecliptic.EclipticSession{
		Contract:     eclip,
		TransactOpts: *auth,
	}
	return session.SetVotingProxy(uint8(galaxy), voter)
}

// Generic L2 Transaction Handler for Proxies
func ProxyTransaction(point, masterTicket, passphrase, target string,
	operation func(context.Context, string, string, string, *ecdsa.PrivateKey) (*types.Transaction, error)) (types.Transaction, error) {
	return l2Transaction(point, masterTicket, passphrase, target, operation)
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

func getGas(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts, maxRetries int) error {
	for i := 0; i < maxRetries; i++ {
		if auth.GasPrice == nil || auth.GasPrice.Cmp(big.NewInt(0)) == 0 {
			gasPrice, err := client.SuggestGasPrice(ctx)
			if err == nil {
				auth.GasPrice = gasPrice
				return nil
			}
			time.Sleep(time.Second * 2)
			continue
		}
		return nil
	}
	return fmt.Errorf("failed to fetch gas price after %d retries", maxRetries)
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
func l2Transaction(point, masterTicket, passphrase, target string,
	operation func(context.Context, string, string, string, *ecdsa.PrivateKey) (*types.Transaction, error)) (types.Transaction, error) {
	masterTicket = strings.TrimPrefix(masterTicket, "~")
	privKey, derivedPubkey, _, _, _, err := ValidateKey(point, masterTicket, passphrase, "", false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	patp, _, err := types.ValidateAndNormalizePatp(point)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	tx, err := operation(ctx, patp, target, derivedPubkey, privKey)
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
	if err := ticketValidation(masterTicket); err != nil {
		return keygen.Wallet{}, nil, "", err
	}
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
	if wallet.Network.Keys.Crypt.Public != pointKey && adjustLife {
		return keygen.Wallet{}, nil, "", fmt.Errorf("%w: expected 0x%s, got %s",
			ErrKeyMismatch, wallet.Network.Keys.Crypt.Public, pInfo.Network.Keys.Crypt)
	}
	return wallet, pInfo, patp, nil
}

// validates and returns a keypair from either master ticket or eth key input
func ValidateKey(point, input, passphrase, seed string, genkey bool) (*ecdsa.PrivateKey, string, *types.Point, types.NetworkKeys, string, error) {
	var ethKey *ecdsa.PrivateKey
	var derivedPubkey string
	var ownerPubkey string
	var pointInfo *types.Point
	var networkKeys types.NetworkKeys
	var authType string
	if strings.Contains(input, "-") {
		wallet, pointInfo, _, err := getWalletAndPoint(point, input, passphrase, 0, false)
		if err != nil {
			return ethKey, derivedPubkey, pointInfo, networkKeys, authType, err
		}
		if err := ticketValidation(input); err != nil {
			return ethKey, derivedPubkey, pointInfo, networkKeys, authType, err
		}
		ethKey, err = crypto.HexToECDSA(wallet.Ownership.Keys.Private)
		if err != nil {
			return ethKey, derivedPubkey, pointInfo, networkKeys, authType, fmt.Errorf("failed to get private key: %v", err)
		}
		authType = "ticket"
		derivedPubkey = wallet.Ownership.Keys.Address
		ownerPubkey = pointInfo.Ownership.Owner.Address
		if genkey {
			networkKeys = types.NetworkKeys{
				Crypt: types.KeyPair{
					Private: wallet.Network.Keys.Crypt.Private,
					Public:  wallet.Network.Keys.Crypt.Public,
				},
				Auth: types.KeyPair{
					Private: wallet.Network.Keys.Auth.Private,
					Public:  wallet.Network.Keys.Auth.Public,
				},
			}
		}
	} else if len(input) > 63 {
		pResp, err := Point(point)
		pointInfo = pResp.Point
		if err != nil {
			return ethKey, derivedPubkey, pointInfo, networkKeys, authType, fmt.Errorf("failed to get point: %v", err)
		}
		ethKey, err = crypto.HexToECDSA(input)
		if err != nil {
			return ethKey, derivedPubkey, pointInfo, networkKeys, authType, fmt.Errorf("failed to get private key: %v", err)
		}
		publicKeyECDSA := ethKey.Public().(*ecdsa.PublicKey)
		derivedPubkey = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		ownerPubkey = pointInfo.Ownership.Owner.Address
		authType = "privkey"
		if genkey {
			networkKeys, err = GenerateNetworkKeysFromSeed(seed)
			if err != nil {
				return ethKey, derivedPubkey, pointInfo, networkKeys, authType, fmt.Errorf("failed to generate network keys from seed: %v", err)
			}
		}
	} else {
		return ethKey, derivedPubkey, pointInfo, networkKeys, authType, fmt.Errorf("invalid private key (too short)")
	}
	if !strings.EqualFold(strings.TrimPrefix(strings.ToLower(derivedPubkey), "0x"), strings.TrimPrefix(strings.ToLower(ownerPubkey), "0x")) {
		return ethKey, derivedPubkey, pointInfo, networkKeys, authType, fmt.Errorf("%w: PKI %s/provided %s", ErrKeyMismatch, ownerPubkey, derivedPubkey)
	}
	return ethKey, derivedPubkey, pointInfo, networkKeys, authType, nil
}

func ticketValidation(input string) error {
	patq := input
	if !strings.HasPrefix(input, "~") {
		patq = "~" + input
	}
	if len(strings.TrimPrefix(input, "~")) < 27 {
		return fmt.Errorf("invalid ticket text (too short)")
	}
	if !co.IsValidPatq(patq) {
		return fmt.Errorf("invalid ticket text (not valid @q)")
	}
	return nil
}

func GenerateNetworkKeysFromSeed(seedHex string) (types.NetworkKeys, error) {
	seedHex = strings.TrimPrefix(seedHex, "0x")
	seedBytes, err := hex.DecodeString(seedHex)
	if err != nil {
		return types.NetworkKeys{}, fmt.Errorf("invalid seed hex: %v", err)
	}
	if len(seedBytes) != 32 {
		return types.NetworkKeys{}, fmt.Errorf("seed must be 32 bytes (64 hex chars), got %d bytes", len(seedBytes))
	}
	var seed [32]byte
	copy(seed[:], seedBytes)
	reversed := make([]byte, len(seed))
	for i, v := range seed {
		reversed[len(seed)-i-1] = v
	}
	hash := sha512.Sum512(reversed)
	cryptSeed := hash[32:]
	authSeed := hash[:32]
	cryptPriv := ed25519.NewKeyFromSeed(cryptSeed)
	cryptPub := cryptPriv.Public().(ed25519.PublicKey)
	authPriv := ed25519.NewKeyFromSeed(authSeed)
	authPub := authPriv.Public().(ed25519.PublicKey)
	cryptPrivReversed := reverseBytes(cryptSeed)
	cryptPubReversed := reverseBytes(cryptPub)
	authPrivReversed := reverseBytes(authSeed)
	authPubReversed := reverseBytes(authPub)
	return types.NetworkKeys{
		Crypt: types.KeyPair{
			Private: hex.EncodeToString(cryptPrivReversed),
			Public:  hex.EncodeToString(cryptPubReversed),
		},
		Auth: types.KeyPair{
			Private: hex.EncodeToString(authPrivReversed),
			Public:  hex.EncodeToString(authPubReversed),
		},
	}, nil
}

func reverseBytes(input []byte) []byte {
	reversed := make([]byte, len(input))
	for i, v := range input {
		reversed[len(input)-i-1] = v
	}
	return reversed
}
