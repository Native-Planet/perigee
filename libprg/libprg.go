package libprg

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	azimuth "github.com/Native-Planet/perigee/abi/azimuthcontract"
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
)

type EclipticSession struct {
	Contract     *ecliptic.Ecliptic
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

/*
Note: for passphrases and life values, just set default values if
you don't have a specific reason to use them
*/

func Escape(ctx context.Context, point, sponsor, masterTicket, passphrase string) (interface{}, error) {
	privateKey, _, pointInfo, _, _, err := ValidateKey(ctx, point, masterTicket, passphrase, "", false, false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return l2Transaction(ctx, point, masterTicket, passphrase, sponsor, roller.Client.Escape)
	} else {
		return l1Escape(ctx, point, sponsor, privateKey)
	}
}

func CancelEscape(ctx context.Context, point, sponsor, masterTicket, passphrase string) (interface{}, error) {
	privateKey, _, pointInfo, _, _, err := ValidateKey(ctx, point, masterTicket, passphrase, "", false, false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return l2Transaction(ctx, point, masterTicket, passphrase, sponsor, roller.Client.CancelEscape)
	} else {
		return l1CancelEscape(ctx, point, privateKey)
	}
}

// Spawn submits a child spawn signed by the parent owner or spawn proxy.
func Spawn(ctx context.Context, parentPoint, masterTicket, passphrase, spawnPoint, owner string) (interface{}, error) {
	privateKey, derivedPubkey, pointInfo, _, _, err := validateKey(ctx, parentPoint, masterTicket, passphrase, "", false, keyProxyAccess{spawn: true})
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	parentPatp, parentInt, err := types.ValidateAndNormalizePatp(parentPoint)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	spawnPatp, spawnInt, err := types.ValidateAndNormalizePatp(spawnPoint)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	if !isChild(parentInt, spawnInt) {
		return types.Transaction{}, fmt.Errorf("%w: %s is not a child of %s", ErrInvalidPoint, spawnPatp, parentPatp)
	}
	normalizedOwner, err := roller.ValidateAddress(owner, false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("invalid owner address: %v", err)
	}
	l2Only := pointInfo.Dominion == "l2" || pointInfo.Dominion == "spawn" ||
		strings.EqualFold(pointInfo.Ownership.TransferProxy.Address, "0x0000000000000000000000000000000000000000")

	if l2Only {
		tx, err := roller.Client.Spawn(ctx, parentPatp, spawnPatp, normalizedOwner, derivedPubkey, privateKey)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("%w: %v", ErrRollerOperation, err)
		}
		return *tx, nil
	}

	return l1Spawn(ctx, spawnPatp, normalizedOwner, derivedPubkey, privateKey)
}

// SpawnRandomOwner generates a child ticket and spawns to its ownership address.
// An empty spawnPoint selects an available child. For L2, accept claims ownership
// and autoKey configures network keys. The result retains the ticket and completed
// transactions when a subsequent operation fails.
func SpawnRandomOwner(ctx context.Context, parentPoint, masterTicket, passphrase, spawnPoint string, accept bool, autoKey bool) (types.SpawnResult, error) {
	var patp string
	var pointInt uint32
	var err error

	if spawnPoint == "" {
		patp, pointInt, err = pickRandomChild(ctx, parentPoint, nil)
		if err != nil {
			return types.SpawnResult{}, fmt.Errorf("pick child: %v", err)
		}
	} else {
		patp, pointInt, err = types.ValidateAndNormalizePatp(spawnPoint)
		if err != nil {
			return types.SpawnResult{}, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
		}
	}
	ticket, err := generateTicket(pointInt)
	if err != nil {
		return types.SpawnResult{}, fmt.Errorf("generate ticket: %v", err)
	}
	trimmedTicket := strings.TrimPrefix(ticket, "~")
	wallet := keygen.GenerateWallet(trimmedTicket, uint32(pointInt), "", 0, true)
	ownerAddr := wallet.Ownership.Keys.Address
	if !strings.HasPrefix(strings.ToLower(ownerAddr), "0x") {
		ownerAddr = "0x" + ownerAddr
	}

	parentInfo, err := Point(ctx, parentPoint)
	if err != nil {
		return types.SpawnResult{}, fmt.Errorf("get parent info: %v", err)
	}
	l2Only := parentInfo.Point.Dominion == "l2" ||
		parentInfo.Point.Dominion == "spawn" ||
		strings.EqualFold(parentInfo.Point.Ownership.TransferProxy.Address, "0x0000000000000000000000000000000000000000")

	result := types.SpawnResult{
		Ship: pointInt, Patp: patp, ParentPatp: parentInfo.PatpName,
		Ticket: ticket, Owner: ownerAddr, Dominion: parentInfo.Point.Dominion,
		SpawnProxy:  parentInfo.Point.Ownership.SpawnProxy.Address,
		ManageProxy: parentInfo.Point.Ownership.ManagementProxy.Address,
	}
	spawnTx, err := Spawn(ctx, parentPoint, masterTicket, passphrase, patp, ownerAddr)
	if err != nil {
		return result, fmt.Errorf("spawn: %v", err)
	}

	result.SpawnTx = spawnTx
	if !l2Only {
		if tx, ok := spawnTx.(*ethTypes.Transaction); ok {
			client, err := ethclient.DialContext(ctx, EthProvider)
			if err != nil {
				return result, fmt.Errorf("failed to connect to Ethereum: %v", err)
			}
			defer client.Close()
			if _, err := waitForReceipt(ctx, client, tx); err != nil {
				return result, fmt.Errorf("waiting for spawn receipt: %v", err)
			}
			childPriv, err := crypto.HexToECDSA(wallet.Ownership.Keys.Private)
			if err != nil {
				return result, fmt.Errorf("child private key: %v", err)
			}
			actTx, err := l1Activate(ctx, patp, childPriv, client)
			if err != nil {
				return result, fmt.Errorf("activate: %v", err)
			}
			result.ActivateTx = actTx
		}
	} else if accept || autoKey {
		childPriv, err := crypto.HexToECDSA(wallet.Ownership.Keys.Private)
		if err != nil {
			return result, fmt.Errorf("child private key: %v", err)
		}
		if accept {
			tx, err := roller.Client.TransferPoint(ctx, patp, true, ownerAddr, wallet.Ownership.Keys.Address, childPriv)
			if err != nil {
				return result, fmt.Errorf("accept transfer: %v", err)
			}
			result.AcceptTx = *tx
		}
		if autoKey {
			tx, err := roller.Client.ConfigureKeys(ctx, patp,
				"0x"+wallet.Network.Keys.Crypt.Public,
				"0x"+wallet.Network.Keys.Auth.Public,
				false, wallet.Ownership.Keys.Address, childPriv)
			if err != nil {
				return result, fmt.Errorf("configure child keys: %v", err)
			}
			result.KeysTx = *tx
			result.NetworkKeys = &types.NetworkKeys{
				Crypt: types.KeyPair{
					Private: wallet.Network.Keys.Crypt.Private,
					Public:  wallet.Network.Keys.Crypt.Public,
				},
				Auth: types.KeyPair{
					Private: wallet.Network.Keys.Auth.Private,
					Public:  wallet.Network.Keys.Auth.Public,
				},
			}
			keyfile, err := roller.Keyfile(wallet.Network.Keys.Crypt.Private, wallet.Network.Keys.Auth.Private, patp, 1)
			if err != nil {
				return result, fmt.Errorf("generate child keyfile: %v", err)
			}
			result.Keyfile = keyfile
		}
	}

	return result, nil
}

// PickRandomChildExcluding selects an unspawned child outside the excluded set.
func PickRandomChildExcluding(ctx context.Context, parentPatp string, excluded []string) (string, uint32, error) {
	excludedSet := make(map[string]struct{}, len(excluded))
	for _, patp := range excluded {
		if patp == "" {
			continue
		}
		normalized, _, err := types.ValidateAndNormalizePatp(patp)
		if err == nil {
			excludedSet[strings.ToLower(normalized)] = struct{}{}
		} else {
			excludedSet[strings.ToLower(strings.TrimSpace(patp))] = struct{}{}
		}
	}
	return pickRandomChild(ctx, parentPatp, excludedSet)
}

func Adopt(ctx context.Context, point, adoptee, masterTicket, passphrase string) (interface{}, error) {
	access := keyProxyAccess{management: true}
	privateKey, _, pointInfo, _, _, err := validateKey(ctx, point, masterTicket, passphrase, "", false, access)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return l2TransactionWithAccess(ctx, point, masterTicket, passphrase, adoptee, access, roller.Client.Adopt)
	} else {
		return l1Adopt(ctx, adoptee, privateKey)
	}
}

func Breach(ctx context.Context, point, ticket, passphrase, seed string) (interface{}, error) {
	var wallet keygen.Wallet
	var pointInfo *types.Point
	var patp string
	var err error
	if seed == "" {
		wallet, pointInfo, patp, err = getWalletAndPoint(ctx, point, ticket, passphrase, 0, false)
		if err != nil {
			return types.Transaction{}, err
		}
	} else {
		// if providing a seed, network keys are generated from it instead of a wallet
		pointRes, err := Point(ctx, point)
		if err != nil {
			return types.Transaction{}, err
		}
		pointInfo = pointRes.Point
		patp = pointRes.PatpName
	}
	if pointInfo.Dominion == "l2" {
		privKey, derivedPubkey, _, networkKeys, _, err := ValidateKey(ctx, point, ticket, passphrase, seed, true, false)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
		}
		keysTx, err := roller.Client.ConfigureKeys(ctx, patp,
			"0x"+networkKeys.Crypt.Public,
			"0x"+networkKeys.Auth.Public,
			true, derivedPubkey, privKey)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("%w: %v", ErrRollerOperation, err)
		}
		return *keysTx, nil
	} else {
		if seed != "" {
			return types.Transaction{}, fmt.Errorf("breach with a private key and seed is only supported for l2 points; %s is %s", patp, pointInfo.Dominion)
		}
		tx, err := setL1NetworkKeys(ctx, point, pointInfo, wallet, true)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
		}
		ethClient, err := ethclient.DialContext(ctx, EthProvider)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("failed to connect to Ethereum node: %v", err)
		}
		defer ethClient.Close()
		receipt, err := waitForReceipt(ctx, ethClient, tx)
		if err != nil {
			return types.Transaction{}, fmt.Errorf("waiting for keys receipt: %v", err)
		}
		return receipt, nil
	}
}

func TransferOwnership(ctx context.Context, point, masterTicket, passphrase, newOwner string, reset bool) (interface{}, error) {
	privateKey, derivedPubkey, pointInfo, _, _, err := ValidateKey(ctx, point, masterTicket, passphrase, "", false, false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		tx, err := roller.Client.TransferPoint(
			ctx,
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
		return l1TransferOwnership(ctx, common.HexToAddress(newOwner), privateKey)
	}
}

func SetManagementProxy(ctx context.Context, point, masterTicket, passphrase, proxy string) (interface{}, error) {
	privateKey, _, pointInfo, _, _, err := ValidateKey(ctx, point, masterTicket, passphrase, "", false, false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return l2Transaction(ctx, point, masterTicket, passphrase, proxy, roller.Client.SetManagementProxy)
	} else {
		return l1SetManagementProxy(ctx, point, common.HexToAddress(proxy), privateKey)
	}
}

func SetTransferProxy(ctx context.Context, point, masterTicket, passphrase, proxy string) (interface{}, error) {
	privateKey, _, pointInfo, _, _, err := ValidateKey(ctx, point, masterTicket, passphrase, "", false, false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return l2Transaction(ctx, point, masterTicket, passphrase, proxy, roller.Client.SetTransferProxy)
	} else {
		return l1SetTransferProxy(ctx, point, common.HexToAddress(proxy), privateKey)
	}
}

func SetSpawnProxy(ctx context.Context, point, masterTicket, passphrase, proxy string) (interface{}, error) {
	privateKey, _, pointInfo, _, _, err := ValidateKey(ctx, point, masterTicket, passphrase, "", false, false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return l2Transaction(ctx, point, masterTicket, passphrase, proxy, roller.Client.SetSpawnProxy)
	} else {
		return l1SetSpawnProxy(ctx, point, common.HexToAddress(proxy), privateKey)
	}
}

func SetVotingProxy(ctx context.Context, galaxy, masterTicket, passphrase, voter string) (interface{}, error) {
	privateKey, _, pointInfo, _, _, err := ValidateKey(ctx, galaxy, masterTicket, passphrase, "", false, false)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("%w: %v", ErrKeyMaterial, err)
	}
	if pointInfo.Dominion == "l2" {
		return types.Transaction{}, fmt.Errorf("voting proxies only exist on l1")
	} else {
		return l1SetVotingProxy(ctx, galaxy, common.HexToAddress(voter), privateKey)
	}
}

func Pending(ctx context.Context, addr string) ([]types.PendingTx, error) {
	if addr == "" {
		return roller.Client.GetAllPending(ctx)
	}
	if !strings.HasPrefix(addr, "0x") {
		_, pInfo, err := validatePointAndGetInfo(ctx, addr)
		if err != nil {
			return nil, err
		}
		addr = pInfo.Ownership.Owner.Address
	}
	return roller.Client.GetPendingByAddress(ctx, addr)
}

func Point(ctx context.Context, point string) (types.PointResp, error) {
	patp, pInfo, err := validatePointAndGetInfo(ctx, point)
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

func Wallet(ctx context.Context, point, masterTicket, passphrase string, life int) (keygen.Wallet, error) {
	wallet, _, _, err := getWalletAndPoint(ctx, point, masterTicket, passphrase, life, false)
	if err != nil {
		return keygen.Wallet{}, err
	}
	return wallet, nil
}

func Keyfile(ctx context.Context, point, masterTicket, passphrase string, life int) (string, error) {
	wallet, pInfo, patp, err := getWalletAndPoint(ctx, point, masterTicket, passphrase, life, true)
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

// newTransactor builds a keyed transactor for the client's chain with a gas
// price set, bound to ctx.
func newTransactor(ctx context.Context, client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}
	auth.Context = ctx
	if err := getGas(ctx, client, auth, 5); err != nil {
		return nil, err
	}
	return auth, nil
}

// newEclipticSession dials EthProvider and builds a transacting Ecliptic
// session. The caller must Close() the returned client.
func newEclipticSession(ctx context.Context, privateKey *ecdsa.PrivateKey) (*ecliptic.EclipticSession, *ethclient.Client, error) {
	if EthProvider == "" {
		return nil, nil, fmt.Errorf("must set ETH_PROVIDER (ex: https://mainnet.infura.io/v3/<your-key>)")
	}
	client, err := ethclient.DialContext(ctx, EthProvider)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to Ethereum: %v", err)
	}
	auth, err := newTransactor(ctx, client, privateKey)
	if err != nil {
		client.Close()
		return nil, nil, err
	}
	eclip, err := ecliptic.NewEcliptic(common.HexToAddress(EclipticContract), client)
	if err != nil {
		client.Close()
		return nil, nil, fmt.Errorf("failed to bind Ecliptic contract: %v", err)
	}
	session := &ecliptic.EclipticSession{
		Contract:     eclip,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
	return session, client, nil
}

func setL1NetworkKeys(ctx context.Context, patp string, point *types.Point, wallet keygen.Wallet, breach bool) (*ethTypes.Transaction, error) {
	patp, pointInt, err := types.ValidateAndNormalizePatp(patp)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	privateKey, err := crypto.HexToECDSA(wallet.Ownership.Keys.Private)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("failed to get private key: %v", err)
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
	session, client, err := newEclipticSession(ctx, privateKey)
	if err != nil {
		return &ethTypes.Transaction{}, err
	}
	defer client.Close()
	return session.ConfigureKeys(pointInt, publicCrypt, publicAuth, CRYPTO_SUITE_VERSION, breach)
}

func l1Escape(ctx context.Context, patp string, sponsor string, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	_, pointInt, err := types.ValidateAndNormalizePatp(patp)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	_, sponsorInt, err := types.ValidateAndNormalizePatp(sponsor)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	session, client, err := newEclipticSession(ctx, privateKey)
	if err != nil {
		return &ethTypes.Transaction{}, err
	}
	defer client.Close()
	allowed, err := session.CanEscapeTo(pointInt, sponsorInt)
	if allowed {
		return session.Escape(pointInt, sponsorInt)
	} else {
		return &ethTypes.Transaction{}, fmt.Errorf("illegal escape: %v", err)
	}
}

func l1Adopt(ctx context.Context, adoptee string, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	_, adopteeInt, err := types.ValidateAndNormalizePatp(adoptee)
	if err != nil {
		return &ethTypes.Transaction{}, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	session, client, err := newEclipticSession(ctx, privateKey)
	if err != nil {
		return &ethTypes.Transaction{}, err
	}
	defer client.Close()
	return session.Adopt(adopteeInt)
}

func l1CancelEscape(ctx context.Context, point string, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	_, pointInt, err := types.ValidateAndNormalizePatp(point)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	session, client, err := newEclipticSession(ctx, privateKey)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return session.CancelEscape(pointInt)
}

func l1Spawn(ctx context.Context, spawnPoint, owner string, signer string, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	_, spawnInt, err := types.ValidateAndNormalizePatp(spawnPoint)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	session, client, err := newEclipticSession(ctx, privateKey)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	// best-effort preflight checks against Azimuth
	azAddress, err := session.Contract.Azimuth(&session.CallOpts)
	if err == nil {
		if az, err := azimuth.NewAzimuth(azAddress, client); err == nil {
			if live, err := az.IsLive(&session.CallOpts, spawnInt); err == nil && live {
				return nil, fmt.Errorf("child %s already active", spawnPoint)
			}
			if canSpawn, err := az.CanSpawnAs(&session.CallOpts, spawnInt, common.HexToAddress(signer)); err == nil && !canSpawn {
				return nil, fmt.Errorf("signer %s cannot spawn %s", signer, spawnPoint)
			}
		}
	}
	return session.Spawn(spawnInt, common.HexToAddress(owner))
}

func l1Activate(ctx context.Context, point string, privateKey *ecdsa.PrivateKey, client *ethclient.Client) (*ethTypes.Transaction, error) {
	_, pointInt, err := types.ValidateAndNormalizePatp(point)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	auth, err := newTransactor(ctx, client, privateKey)
	if err != nil {
		return nil, err
	}
	eclip, err := ecliptic.NewEcliptic(common.HexToAddress(EclipticContract), client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind Ecliptic contract: %v", err)
	}
	azAddress, err := eclip.Azimuth(&bind.CallOpts{From: auth.From, Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to load azimuth address: %v", err)
	}
	az, err := azimuth.NewAzimuth(azAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind Azimuth contract: %v", err)
	}
	session := &azimuth.AzimuthSession{
		Contract:     az,
		TransactOpts: *auth,
	}
	return session.ActivatePoint(pointInt)
}

// Transfer Ownership (L1)
func l1TransferOwnership(ctx context.Context, newOwner common.Address, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	session, client, err := newEclipticSession(ctx, privateKey)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return session.TransferOwnership(newOwner)
}

// Set Management Proxy (L1)
func l1SetManagementProxy(ctx context.Context, pointStr string, manager common.Address, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	_, point, err := types.ValidateAndNormalizePatp(pointStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	session, client, err := newEclipticSession(ctx, privateKey)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return session.SetManagementProxy(uint32(point), manager)
}

// Set Transfer Proxy (L1)
func l1SetTransferProxy(ctx context.Context, pointStr string, proxy common.Address, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	_, point, err := types.ValidateAndNormalizePatp(pointStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	session, client, err := newEclipticSession(ctx, privateKey)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return session.SetTransferProxy(uint32(point), proxy)
}

// Set Spawn Proxy (L1)
func l1SetSpawnProxy(ctx context.Context, prefixStr string, proxy common.Address, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	prefix, err := strconv.ParseUint(prefixStr, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("invalid prefix: %v", err)
	}
	session, client, err := newEclipticSession(ctx, privateKey)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return session.SetSpawnProxy(uint16(prefix), proxy)
}

// Set Voting Proxy (L1)
func l1SetVotingProxy(ctx context.Context, galaxyStr string, voter common.Address, privateKey *ecdsa.PrivateKey) (*ethTypes.Transaction, error) {
	galaxy, err := strconv.ParseUint(galaxyStr, 10, 8)
	if err != nil {
		return nil, fmt.Errorf("invalid galaxy: %v", err)
	}
	session, client, err := newEclipticSession(ctx, privateKey)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return session.SetVotingProxy(uint8(galaxy), voter)
}

// Generic L2 Transaction Handler for Proxies
func ProxyTransaction(ctx context.Context, point, masterTicket, passphrase, target string,
	operation func(context.Context, string, string, string, *ecdsa.PrivateKey) (*types.Transaction, error)) (types.Transaction, error) {
	return l2Transaction(ctx, point, masterTicket, passphrase, target, operation)
}

func Wait(ctx context.Context, duration time.Duration, keysTx string) error {
	if duration == 0 {
		return nil
	}
	waitTime, err := roller.Client.WhenNextBatch(ctx)
	if err != nil {
		return fmt.Errorf("getting batch info: %v", err)
	}
	if waitTime > duration {
		waitTime = duration
	}
	if waitTime <= 0 {
		waitTime = 10 * time.Second
	}
	deadline := time.Now().Add(duration)
	ticker := time.NewTicker(waitTime)
	defer ticker.Stop()

	for {
		pending, err := Pending(ctx, "")
		if err != nil {
			fmt.Printf("Error getting pending ships: %v\n", err)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(10 * time.Second):
			}
			if time.Now().After(deadline) {
				return fmt.Errorf("timeout waiting for transaction after %v", duration)
			}
			continue
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
		case <-ctx.Done():
			return ctx.Err()
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

func ticketEntropyBytes(point uint32) int {
	switch {
	case point < 256:
		return 48 // galaxies
	case point < 65536:
		return 16 // stars
	default:
		return 8 // planets
	}
}

func generateTicket(point uint32) (string, error) {
	bytesLen := ticketEntropyBytes(point)
	buf := make([]byte, bytesLen)
	if _, err := rand.Read(buf); err != nil {
		return "", fmt.Errorf("generate entropy: %v", err)
	}
	return co.Hex2Patq(hex.EncodeToString(buf))
}

func pickRandomChild(ctx context.Context, parentPatp string, excludedSet map[string]struct{}) (string, uint32, error) {
	parentPatp, parentInt, err := types.ValidateAndNormalizePatp(parentPatp)
	if err != nil {
		return "", 0, fmt.Errorf("%w: %v", ErrInvalidPoint, err)
	}
	var childCount int
	var shift uint
	switch {
	case parentInt < 256:
		childCount = 256
		shift = 8
	case parentInt < 65536:
		childCount = 65536
		shift = 16
	default:
		return "", 0, fmt.Errorf("point %s cannot spawn children", parentPatp)
	}

	unspawned, err := roller.Client.GetUnspawned(ctx, parentPatp)
	if err == nil && len(unspawned) > 0 {
		eligible := make([]types.ShipInfo, 0, len(unspawned))
		for _, ship := range unspawned {
			name, point, err := types.ValidateAndNormalizePatp(ship.Name)
			if err != nil || !isChild(parentInt, point) {
				return "", 0, fmt.Errorf("invalid child %q for %s", ship.Name, parentPatp)
			}
			ship.Name = name
			if _, excluded := excludedSet[strings.ToLower(ship.Name)]; excluded {
				continue
			}
			eligible = append(eligible, ship)
		}
		if len(eligible) > 0 {
			idx, err := randInt(len(eligible))
			if err != nil {
				return "", 0, fmt.Errorf("choose eligible child: %v", err)
			}
			childPatp := eligible[idx].Name
			_, childInt, err := types.ValidateAndNormalizePatp(childPatp)
			if err != nil {
				return "", 0, fmt.Errorf("normalize child: %v", err)
			}
			return childPatp, childInt, nil
		}
	}

	spawned, err := roller.Client.GetSpawned(ctx, parentPatp)
	if err != nil {
		return "", 0, fmt.Errorf("get spawned: %v", err)
	}
	spawnedSet := make(map[string]struct{}, len(spawned))
	for _, s := range spawned {
		name, _, err := types.ValidateAndNormalizePatp(s.Name)
		if err != nil {
			return "", 0, fmt.Errorf("normalize spawned child: %v", err)
		}
		spawnedSet[strings.ToLower(name)] = struct{}{}
	}

	// Index zero is the parent itself. Visit every child once from a random start.
	childCount--
	start, err := randInt(childCount)
	if err != nil {
		return "", 0, fmt.Errorf("random child: %v", err)
	}
	for i := 0; i < childCount; i++ {
		if err := ctx.Err(); err != nil {
			return "", 0, err
		}
		idx := (start+i)%childCount + 1
		childInt := uint32(idx)<<shift | parentInt
		childPatp, err := co.Point2Patp(big.NewInt(int64(childInt)))
		if err != nil {
			return "", 0, fmt.Errorf("convert child patp: %v", err)
		}
		if _, exists := spawnedSet[strings.ToLower(childPatp)]; exists {
			continue
		}
		if _, excluded := excludedSet[strings.ToLower(childPatp)]; excluded {
			continue
		}
		return childPatp, childInt, nil
	}

	return "", 0, fmt.Errorf("no unspawned children available for %s", parentPatp)
}

func isChild(parent, child uint32) bool {
	switch {
	case parent < 256:
		return child >= 256 && child < 65536 && child&0xff == parent
	case parent < 65536:
		return child >= 65536 && child&0xffff == parent
	default:
		return false
	}
}

func randInt(max int) (int, error) {
	if max <= 0 {
		return 0, fmt.Errorf("invalid max for random int")
	}
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(nBig.Int64()), nil
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
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
}

func validatePointAndGetInfo(ctx context.Context, point string) (string, *types.Point, error) {
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
func l2Transaction(ctx context.Context, point, masterTicket, passphrase, target string,
	operation func(context.Context, string, string, string, *ecdsa.PrivateKey) (*types.Transaction, error)) (types.Transaction, error) {
	return l2TransactionWithAccess(ctx, point, masterTicket, passphrase, target, keyProxyAccess{}, operation)
}

func l2TransactionWithAccess(ctx context.Context, point, masterTicket, passphrase, target string,
	access keyProxyAccess,
	operation func(context.Context, string, string, string, *ecdsa.PrivateKey) (*types.Transaction, error)) (types.Transaction, error) {
	masterTicket = strings.TrimPrefix(masterTicket, "~")
	privKey, derivedPubkey, _, _, _, err := validateKey(ctx, point, masterTicket, passphrase, "", false, access)
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
func getWalletAndPoint(ctx context.Context, point, masterTicket, passphrase string, life int, adjustLife bool) (keygen.Wallet, *types.Point, string, error) {
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

type keyProxyAccess struct {
	management bool
	spawn      bool
}

// ValidateKey derives and validates a keypair from a master ticket or Ethereum key.
// allowProxy also accepts the point's management and spawn proxy addresses.
func ValidateKey(ctx context.Context, point, input, passphrase, seed string, genkey bool, allowProxy bool) (*ecdsa.PrivateKey, string, *types.Point, types.NetworkKeys, string, error) {
	access := keyProxyAccess{
		management: allowProxy,
		spawn:      allowProxy,
	}
	return validateKey(ctx, point, input, passphrase, seed, genkey, access)
}

func validateKey(ctx context.Context, point, input, passphrase, seed string, genkey bool, access keyProxyAccess) (*ecdsa.PrivateKey, string, *types.Point, types.NetworkKeys, string, error) {
	var ethKey *ecdsa.PrivateKey
	var derivedPubkey string
	var ownerPubkey string
	var pointInfo *types.Point
	var networkKeys types.NetworkKeys
	var authType string
	if strings.Contains(input, "-") {
		wallet, pInfo, _, err := getWalletAndPoint(ctx, point, input, passphrase, 0, false)
		if err != nil {
			return ethKey, derivedPubkey, pInfo, networkKeys, authType, err
		}
		pointInfo = pInfo
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
		pResp, err := Point(ctx, point)
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
	allowed := []string{ownerPubkey}
	if access.management {
		allowed = append(allowed, pointInfo.Ownership.ManagementProxy.Address)
	}
	if access.spawn {
		allowed = append(allowed, pointInfo.Ownership.SpawnProxy.Address)
	}
	match := false
	for _, a := range allowed {
		if strings.EqualFold(strings.TrimPrefix(strings.ToLower(derivedPubkey), "0x"), strings.TrimPrefix(strings.ToLower(a), "0x")) {
			match = true
			break
		}
	}
	if !match {
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

// +code
// step is for incrementing the code revision
func GenerateCode(ctx context.Context, point, ticket, passphrase string, life, step int) (string, error) {
	wallet, _, _, err := getWalletAndPoint(ctx, point, ticket, passphrase, life, true)
	if err != nil {
		return "", err
	}
	ringHex := wallet.Network.Keys.Crypt.Private + wallet.Network.Keys.Auth.Private + "42"
	ring, err := hex2buf(ringHex)
	if err != nil {
		return "", err
	}
	bsalt, _ := new(big.Int).SetString("73736170", 16)
	esalt := new(big.Int).SetInt64(int64(step))
	saltInt := new(big.Int).Add(bsalt, esalt)
	salt, err := hex2buf(saltInt.Text(16))
	if err != nil {
		return "", err
	}
	hash := shax(ring)
	result := shaf(hash, salt)
	half := result[:len(result)/2]
	hexHalf := buf2hex(half)
	patp, err := co.Hex2Patp(hexHalf)
	if err != nil {
		return "", err
	}
	return patp[1:], nil
}

func shas(buf, salt []byte) []byte {
	bufHash := shax(buf)
	paddedSalt := make([]byte, len(bufHash))
	copy(paddedSalt, salt)
	xorred := xor(paddedSalt, bufHash)
	return shax(xorred)
}

func shaf(buf, salt []byte) []byte {
	result := shas(buf, salt)
	halfway := len(result) / 2
	front := result[:halfway]
	back := result[halfway:]
	return xor(front, back)
}

func shax(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func xor(a, b []byte) []byte {
	length := len(a)
	if len(b) < length {
		length = len(b)
	}
	result := make([]byte, len(b))
	for i := 0; i < length; i++ {
		result[i] = a[i] ^ b[i]
	}
	for i := length; i < len(b); i++ {
		result[i] = b[i]
	}
	return result
}

func hex2buf(hexStr string) ([]byte, error) {
	buf, err := hex.DecodeString(hexStr)
	if err != nil {
		return []byte{}, fmt.Errorf("Invalid hex string: %s", hexStr)
	}
	reversed := make([]byte, len(buf))
	for i := 0; i < len(buf); i++ {
		reversed[i] = buf[len(buf)-i-1]
	}
	return reversed, nil
}

func buf2hex(buf []byte) string {
	reversed := make([]byte, len(buf))
	for i := 0; i < len(buf); i++ {
		reversed[i] = buf[len(buf)-i-1]
	}
	return hex.EncodeToString(reversed)
}
