package wallet

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"perigee/types"

	"github.com/deelawn/urbit-gob/co"
	"github.com/pkova/argon2"
)

// thank you pyry
func Argon2u(entropy []byte, ship uint32) ([]byte, error) {
	salt := fmt.Sprintf("urbitkeygen%d", ship)
	return argon2.UKey(entropy, []byte(salt), 1, 512000, 4, 32), nil
}

func GenerateWallet(config types.WalletConfig) (*types.Wallet, error) {
	if config.Ticket == "" {
		return nil, errors.New("no ticket provided")
	}
	if config.Ship == 0 {
		return nil, errors.New("no ship provided")
	}
	ticketHex, err := co.Patq2Hex(config.Ticket)
	if err != nil {
		return nil, fmt.Errorf("failed to convert ticket to hex: %v", err)
	}
	ticketBytes, err := HexToBytes(ticketHex)
	if err != nil {
		return nil, fmt.Errorf("invalid ticket: %v", err)
	}
	patp, err := co.Point2Patp(new(big.Int).SetUint64(uint64(config.Ship)))
	if err != nil {
		return nil, fmt.Errorf("invalid ship: %v", err)
	}
	tier, err := co.Clan(patp)
	if err != nil {
		return nil, fmt.Errorf("invalid ship: %v", err)
	}
	masterSeed, err := Argon2u(ticketBytes, config.Ship)
	if err != nil {
		return nil, fmt.Errorf("failed to generate master seed: %v", err)
	}

	wallet := &types.Wallet{
		Meta: types.WalletMeta{
			Generator: struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			}{
				Name:    "perigee-go",
				Version: "0.1.0",
			},
			Spec:       "UP8",
			Ship:       config.Ship,
			Patp:       patp,
			Tier:       tier,
			Passphrase: config.Passphrase,
		},
		Ticket: config.Ticket,
	}

	if err := wallet.GenerateNodes(masterSeed); err != nil {
		return nil, fmt.Errorf("failed to generate nodes: %v", err)
	}
	if config.Boot {
		if err := wallet.GenerateNetworkKeys(config.Revision); err != nil {
			return nil, fmt.Errorf("failed to generate network keys: %v", err)
		}
	}
	return wallet, nil
}

func HexToBytes(hexInput string) ([]byte, error) {
	cleanHex := StripHexPrefix(hexInput)
	if len(cleanHex)%2 == 1 {
		cleanHex = "0" + cleanHex
	}
	buf, err := hex.DecodeString(cleanHex)
	if err != nil {
		return nil, err
	}
	if len(cleanHex) != 0 && len(buf) == 0 {
		return nil, errors.New("invalid hex string: " + hexInput)
	}
	return buf, nil
}

func StripHexPrefix(hex string) string {
	if len(hex) >= 2 && hex[0:2] == "0x" {
		return hex[2:]
	}
	return hex
}

func AddHexPrefix(hex string) string {
	if len(hex) >= 2 && hex[0:2] == "0x" {
		return hex
	}
	return "0x" + hex
}
