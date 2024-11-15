package types

import (
	"crypto/ed25519"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/deelawn/urbit-gob/co"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

type ChildSeedType string

const (
	GalaxyMin = 0x00000000
	GalaxyMax = 0x000000ff
	PlanetMin = 0x00010000
	PlanetMax = 0xffffffff
)

func IsGalaxy(ship uint32) bool {
	return ship >= GalaxyMin && ship <= GalaxyMax
}

func IsPlanet(ship uint32) bool {
	return ship >= PlanetMin && ship <= PlanetMax
}

const (
	Ownership                ChildSeedType = "ownership"
	Transfer                 ChildSeedType = "transfer"
	Spawn                    ChildSeedType = "spawn"
	Voting                   ChildSeedType = "voting"
	Management               ChildSeedType = "management"
	Network                  ChildSeedType = "network"
	DerivationPath                         = "m/44'/60'/0'/0/0"
	NetworkKeyCurveParameter               = "42"
)

type KeyPair struct {
	Public  string `json:"public"`
	Private string `json:"private"`
	Chain   string `json:"chain,omitempty"`
	Address string `json:"address,omitempty"`
}

type NetworkKeys struct {
	Crypt KeyPair `json:"crypt"`
	Auth  KeyPair `json:"auth"`
}

type Node struct {
	Type           ChildSeedType `json:"type"`
	Seed           string        `json:"seed"`
	Keys           KeyPair       `json:"keys"`
	DerivationPath string        `json:"derivationPath"`
}

type NetworkInfo struct {
	Type ChildSeedType `json:"type"`
	Seed string        `json:"seed"`
	Keys NetworkKeys   `json:"keys"`
}

type WalletConfig struct {
	Ticket     string `json:"ticket"`
	Ship       uint32 `json:"ship"`
	Passphrase string `json:"passphrase,omitempty"`
	Revision   uint32 `json:"revision,omitempty"`
	Boot       bool   `json:"boot,omitempty"`
}

type WalletMeta struct {
	Generator struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"generator"`
	Spec       string `json:"spec"`
	Ship       uint32 `json:"ship"`
	Patp       string `json:"patp"`
	Tier       string `json:"tier"`
	Passphrase string `json:"passphrase,omitempty"`
}

// urbit HD wallet
type Wallet struct {
	Meta       WalletMeta  `json:"meta"`
	Ticket     string      `json:"ticket"`
	Shards     []string    `json:"shards"`
	Ownership  Node        `json:"ownership"`
	Transfer   Node        `json:"transfer"`
	Spawn      Node        `json:"spawn,omitempty"`
	Voting     Node        `json:"voting,omitempty"`
	Management Node        `json:"management"`
	Network    NetworkInfo `json:"network,omitempty"`
}

func deriveNetworkKeys(seedHex string) (*NetworkKeys, error) {
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		return nil, err
	}
	for i, j := 0, len(seed)-1; i < j; i, j = i+1, j-1 {
		seed[i], seed[j] = seed[j], seed[i]
	}
	h := sha512.Sum512(seed)
	cryptSeed := h[32:]
	authSeed := h[:32]
	cryptPrivKey := ed25519.NewKeyFromSeed(cryptSeed)
	authPrivKey := ed25519.NewKeyFromSeed(authSeed)
	cryptPubKey := cryptPrivKey[32:]
	authPubKey := authPrivKey[32:]
	reverseBytes(cryptPrivKey[:32])
	reverseBytes(cryptPubKey)
	reverseBytes(authPrivKey[:32])
	reverseBytes(authPubKey)

	return &NetworkKeys{
		Crypt: KeyPair{
			Private: hex.EncodeToString(cryptPrivKey[:32]),
			Public:  hex.EncodeToString(cryptPubKey),
		},
		Auth: KeyPair{
			Private: hex.EncodeToString(authPrivKey[:32]),
			Public:  hex.EncodeToString(authPubKey),
		},
	}, nil
}
func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func (w *Wallet) GenerateNetworkKeys(revision uint32) error {
	networkSeed, err := deriveNetworkSeed(w.Management.Seed, w.Meta.Passphrase, revision)
	if err != nil {
		return err
	}
	seedHex := hex.EncodeToString(networkSeed)
	networkKeys, err := deriveNetworkKeys(seedHex)
	if err != nil {
		return err
	}
	w.Network = NetworkInfo{
		Type: Network,
		Seed: seedHex,
		Keys: *networkKeys,
	}
	return nil
}

func (w *Wallet) GenerateNodes(masterSeed []byte) error {
	if len(w.Shards) == 0 {
		shards, err := ShardTicket(w.Ticket)
		if err != nil {
			return fmt.Errorf("failed to generate ticket shards: %v", err)
		}
		w.Shards = shards
	}
	ownershipSeed, err := DeriveNodeSeed(masterSeed, Ownership)
	if err != nil {
		return err
	}
	ownershipKeys, err := DeriveNodeKeys(ownershipSeed, DerivationPath, w.Meta.Passphrase)
	if err != nil {
		return err
	}
	w.Ownership = Node{
		Type:           Ownership,
		Seed:           ownershipSeed,
		Keys:           *ownershipKeys,
		DerivationPath: DerivationPath,
	}
	transferSeed, err := DeriveNodeSeed(masterSeed, Transfer)
	if err != nil {
		return err
	}
	transferKeys, err := DeriveNodeKeys(transferSeed, DerivationPath, w.Meta.Passphrase)
	if err != nil {
		return err
	}
	w.Transfer = Node{
		Type:           Transfer,
		Seed:           transferSeed,
		Keys:           *transferKeys,
		DerivationPath: DerivationPath,
	}
	managementSeed, err := DeriveNodeSeed(masterSeed, Management)
	if err != nil {
		return err
	}
	managementKeys, err := DeriveNodeKeys(managementSeed, DerivationPath, w.Meta.Passphrase)
	if err != nil {
		return err
	}
	w.Management = Node{
		Type:           Management,
		Seed:           managementSeed,
		Keys:           *managementKeys,
		DerivationPath: DerivationPath,
	}
	if !IsPlanet(w.Meta.Ship) {
		spawnSeed, err := DeriveNodeSeed(masterSeed, Spawn)
		if err != nil {
			return err
		}
		spawnKeys, err := DeriveNodeKeys(spawnSeed, DerivationPath, w.Meta.Passphrase)
		if err != nil {
			return err
		}
		w.Spawn = Node{
			Type:           Spawn,
			Seed:           spawnSeed,
			Keys:           *spawnKeys,
			DerivationPath: DerivationPath,
		}
	}
	if IsGalaxy(w.Meta.Ship) {
		votingSeed, err := DeriveNodeSeed(masterSeed, Voting)
		if err != nil {
			return err
		}
		votingKeys, err := DeriveNodeKeys(votingSeed, DerivationPath, w.Meta.Passphrase)
		if err != nil {
			return err
		}
		w.Voting = Node{
			Type:           Voting,
			Seed:           votingSeed,
			Keys:           *votingKeys,
			DerivationPath: DerivationPath,
		}
	}

	return nil
}

// derive a network seed from management mnemonic
func deriveNetworkSeed(mnemonic, passphrase string, revision uint32) ([]byte, error) {
	seed := bip39.NewSeed(mnemonic, passphrase)
	data := append(seed, []byte(Network)...)
	data = append(data, []byte(fmt.Sprintf("%d", revision))...)
	hash := sha256.Sum256(data)
	if revision != 0 {
		hash = sha256.Sum256(hash[:])
	}

	return hash[:], nil
}

// converts bip44 path string to slice of child indices
func parsePath(path string) ([]uint32, error) {
	if !strings.HasPrefix(path, "m/") {
		return nil, fmt.Errorf("invalid derivation path: %s", path)
	}
	components := strings.Split(path[2:], "/")
	indices := make([]uint32, len(components))
	for i, component := range components {
		var value uint32
		hardened := strings.HasSuffix(component, "'")
		if hardened {
			component = strings.TrimSuffix(component, "'")
		}
		_, err := fmt.Sscanf(component, "%d", &value)
		if err != nil {
			return nil, fmt.Errorf("invalid path component: %s", component)
		}
		if hardened {
			value += hdkeychain.HardenedKeyStart
		}
		indices[i] = value
	}

	return indices, nil
}

// derive bip39
func DeriveNodeSeed(master []byte, seedType ChildSeedType) (string, error) {
	hash := sha256.Sum256(append(master, []byte(seedType)...))
	return bip39.NewMnemonic(hash[:])
}

// derive secp256k1
func DeriveNodeKeys(mnemonic, derivationPath, passphrase string) (*KeyPair, error) {
	seed := bip39.NewSeed(mnemonic, passphrase)
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}
	indices, err := parsePath(derivationPath)
	if err != nil {
		return nil, err
	}
	key := masterKey
	for _, index := range indices {
		key, err = key.Derive(index)
		if err != nil {
			return nil, err
		}
	}
	privKey, err := key.ECPrivKey()
	if err != nil {
		return nil, err
	}
	pubKey, err := key.ECPubKey()
	if err != nil {
		return nil, err
	}
	privKeyHex := hex.EncodeToString(privKey.Serialize())
	pubKeyHex := hex.EncodeToString(pubKey.SerializeCompressed())
	ethPubKey := pubKey.ToECDSA()
	address := crypto.PubkeyToAddress(*ethPubKey)
	return &KeyPair{
		Private: privKeyHex,
		Public:  pubKeyHex,
		Chain:   hex.EncodeToString(key.ChainCode()),
		Address: address.Hex(),
	}, nil
}

func ShardTicket(ticket string) ([]string, error) {
	ticketHex, err := co.Patq2Hex(ticket)
	if err != nil {
		return nil, fmt.Errorf("failed to convert ticket to hex: %v", err)
	}
	ticketBytes, err := hex.DecodeString(ticketHex)
	if err != nil {
		return nil, fmt.Errorf("invalid hex ticket: %v", err)
	}
	if len(ticketBytes) != 48 {
		return []string{ticket}, nil
	}
	shard0 := ticketBytes[:32]
	shard1 := ticketBytes[16:]
	shard2 := append(ticketBytes[:16], ticketBytes[32:]...)
	shards := make([]string, 3)
	for i, shard := range [][]byte{shard0, shard1, shard2} {
		shardHex := hex.EncodeToString(shard)
		shardPatq, err := co.Hex2Patq(shardHex)
		if err != nil {
			return nil, fmt.Errorf("failed to convert shard %d to @q: %v", i, err)
		}
		shards[i] = shardPatq
	}
	verified, err := verifyShardsCanCombine(shards, ticket)
	if err != nil {
		return nil, fmt.Errorf("shard verification failed: %v", err)
	}
	if !verified {
		return nil, errors.New("produced invalid shards")
	}

	return shards, nil
}

func verifyShardsCanCombine(shards []string, originalTicket string) (bool, error) {
	combinations := [][3]string{
		{shards[0], shards[1], ""},
		{shards[0], "", shards[2]},
		{"", shards[1], shards[2]},
	}
	for _, combo := range combinations {
		result, err := CombineShards(combo[:])
		if err != nil {
			return false, err
		}
		if result != originalTicket {
			return false, nil
		}
	}
	return true, nil
}

// expects an array of three strings, with an empty string in place of the missing shard
func CombineShards(shards []string) (string, error) {
	if len(shards) != 3 {
		return "", errors.New("combine: need exactly three shard positions")
	}
	emptyCount := 0
	for _, shard := range shards {
		if shard == "" {
			emptyCount++
		}
	}
	if emptyCount != 1 {
		return "", errors.New("combine: need exactly two shards")
	}
	s0, s1, s2 := shards[0], shards[1], shards[2]
	var h0, h1, h2 string
	var err error
	if s0 != "" {
		h0, err = co.Patq2Hex(s0)
		if err != nil {
			return "", fmt.Errorf("failed to convert shard 0 to hex: %v", err)
		}
	}
	if s1 != "" {
		h1, err = co.Patq2Hex(s1)
		if err != nil {
			return "", fmt.Errorf("failed to convert shard 1 to hex: %v", err)
		}
	}
	if s2 != "" {
		h2, err = co.Patq2Hex(s2)
		if err != nil {
			return "", fmt.Errorf("failed to convert shard 2 to hex: %v", err)
		}
	}
	var resultHex string
	switch {
	case s0 != "" && s1 != "":
		resultHex = h0[:64] + h1
	case s0 != "" && s2 != "":
		resultHex = h0 + h2[64:]
	case s1 != "" && s2 != "":
		resultHex = h2[:64] + h1
	default:
		return "", errors.New("invalid shard combination")
	}
	result, err := co.Hex2Patq(resultHex)
	if err != nil {
		return "", fmt.Errorf("failed to convert result to @q: %v", err)
	}
	return result, nil
}
