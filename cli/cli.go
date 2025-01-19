package cli

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Native-Planet/perigee/roller"
	"github.com/Native-Planet/perigee/types"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nathanlever/keygen"
	"github.com/spf13/cobra"
)

var (
	ctx    = context.Background()
	apiURL = os.Getenv("API_URL")
)

var GetPointCmd = &cobra.Command{
	Use:   "get-point",
	Short: "Retrieve point information",
	RunE: func(cmd *cobra.Command, args []string) error {
		point, _ := cmd.Flags().GetString("point")
		if point == "" {
			return fmt.Errorf("point is required")
		}
		patp, _, err := types.ValidateAndNormalizePatp(point)
		if err != nil {
			return fmt.Errorf("invalid point")
		}
		pInfo, err := roller.Client.GetPoint(ctx, patp)
		if err != nil {
			return fmt.Errorf("error getting point: %v", err)
		}
		resp := types.PointResp{
			Point:    pInfo,
			PatpName: patp,
		}
		if err := resp.Point.ResolveSponsorPatp(); err != nil {
			return fmt.Errorf("invalid sponsor point")
		}
		jsonData, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshaling response: %v", err)
		}
		fmt.Println(string(jsonData))
		return nil
	},
}

var GetPendingCmd = &cobra.Command{
	Use:   "get-pending",
	Short: "List pending rollup queue",
	RunE: func(cmd *cobra.Command, args []string) error {
		point, _ := cmd.Flags().GetString("point")
		address, _ := cmd.Flags().GetString("address")
		if point == "" && address == "" {
			pending, err := roller.Client.GetAllPending(ctx)
			if err != nil {
				return fmt.Errorf("error getting pending ships: %v", err)
			}
			jsonData, err := json.MarshalIndent(pending, "", "  ")
			if err != nil {
				return fmt.Errorf("error marshaling response: %v", err)
			}
			fmt.Println(string(jsonData))
			return nil
		} else if point != "" {
			patp, _, err := types.ValidateAndNormalizePatp(point)
			if err != nil {
				return fmt.Errorf("invalid point")
			}
			pInfo, err := roller.Client.GetPoint(ctx, patp)
			if err != nil {
				return fmt.Errorf("error getting point: %v", err)
			}
			pending, err := roller.Client.GetPendingByAddress(ctx, pInfo.Ownership.Owner.Address)
			if err != nil {
				return fmt.Errorf("error getting pending ships: %v", err)
			}
			jsonData, err := json.MarshalIndent(pending, "", "  ")
			if err != nil {
				return fmt.Errorf("error marshaling response: %v", err)
			}
			fmt.Println(string(jsonData))
			return nil
		} else if address != "" {
			pending, err := roller.Client.GetPendingByAddress(ctx, address)
			if err != nil {
				return fmt.Errorf("error getting pending ships: %v", err)
			}
			jsonData, err := json.MarshalIndent(pending, "", "  ")
			if err != nil {
				return fmt.Errorf("error marshaling response: %v", err)
			}
			fmt.Println(string(jsonData))
		}
		return nil
	},
}

var ModBreachCmd = &cobra.Command{
	Use:   "breach",
	Short: "Breach a point",
	RunE: func(cmd *cobra.Command, args []string) error {
		point, err := cmd.Flags().GetString("point")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		if point == "" {
			return fmt.Errorf("point is required")
		}
		masterTicket, err := cmd.Flags().GetString("master-ticket")
		if err != nil {
			return fmt.Errorf("error getting master-ticket flag: %v", err)
		}
		if masterTicket == "" {
			return fmt.Errorf("master-ticket is required")
		}
		passphrase, err := cmd.Flags().GetString("passphrase")
		if err != nil {
			return fmt.Errorf("error getting passphrase flag: %v", err)
		}
		patp, pointInt, err := types.ValidateAndNormalizePatp(point)
		if err != nil {
			return fmt.Errorf("invalid point: %v", err)
		}
		pInfo, err := roller.Client.GetPoint(ctx, patp)
		if err != nil {
			return fmt.Errorf("error getting point: %v", err)
		}
		currentRevision := fmt.Sprintf("%v", pInfo.Network.Keys.Life)
		rev, err := strconv.Atoi(currentRevision)
		if err != nil {
			return fmt.Errorf("invalid life: %v", err)
		}
		wallet := keygen.GenerateWallet(masterTicket, uint32(pointInt), passphrase, uint(rev), true)
		privKey, err := crypto.HexToECDSA(wallet.Ownership.Keys.Private)
		if err != nil {
			return fmt.Errorf("invalid key material: %v", err)
		}
		keysTx, err := roller.Client.ConfigureKeys(ctx, patp, "0x"+wallet.Network.Keys.Crypt.Public, "0x"+wallet.Network.Keys.Auth.Public, true, wallet.Ownership.Keys.Address, privKey)
		if err != nil {
			return fmt.Errorf("error processing breach: %v", err)
		}
		jsonData, err := json.MarshalIndent(keysTx, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshaling response: %v", err)
		}
		fmt.Println(string(jsonData))
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		var deadline time.Time
		if cmd.Flags().Changed("wait") {
			duration, err := cmd.Flags().GetDuration("wait")
			if err != nil {
				return fmt.Errorf("invalid duration: %v", err)
			}
			deadline = time.Now().Add(duration)
		}
		for {
			pending, err := roller.Client.GetPendingByAddress(ctx, pInfo.Ownership.Owner.Address)
			if err != nil {
				return fmt.Errorf("error getting pending ships: %v", err)
			}
			found := false
			for _, tx := range pending {
				if tx.RawTx.Sig == keysTx.Signature {
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Transaction committed to chain")
				return nil
			}
			if cmd.Flags().Changed("wait") {
				if time.Now().After(deadline) {
					return fmt.Errorf("timeout after %v", time.Since(deadline.Add(-deadline.Sub(time.Now()))))
				}
				select {
				case <-ticker.C:
					continue
				case <-time.After(deadline.Sub(time.Now())):
					return fmt.Errorf("timeout after %v", time.Since(deadline.Add(-deadline.Sub(time.Now()))))
				}
			} else {
				<-ticker.C
			}
		}
	},
}

var ModEscapeCmd = &cobra.Command{
	Use:   "escape",
	Short: "Escape from a sponsor",
	RunE: func(cmd *cobra.Command, args []string) error {
		var privKey *ecdsa.PrivateKey
		var addr string
		point, err := cmd.Flags().GetString("point")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		sponsorArg, err := cmd.Flags().GetString("sponsor")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		if point == "" {
			return fmt.Errorf("point is required")
		}
		if sponsorArg == "" {
			return fmt.Errorf("adoptee is required")
		}
		masterTicket, err := cmd.Flags().GetString("master-ticket")
		if err != nil {
			return fmt.Errorf("error getting master-ticket flag: %v", err)
		}
		ethKey, err := cmd.Flags().GetString("private-key")
		if err != nil {
			return fmt.Errorf("error getting private-key flag: %v", err)
		}
		if masterTicket == "" && ethKey == "" {
			return fmt.Errorf("master-ticket is required")
		}
		passphrase, err := cmd.Flags().GetString("passphrase")
		if err != nil {
			return fmt.Errorf("error getting passphrase flag: %v", err)
		}
		patp, pointInt, err := types.ValidateAndNormalizePatp(point)
		if err != nil {
			return fmt.Errorf("invalid point: %v", err)
		}
		sponsor, _, err := types.ValidateAndNormalizePatp(sponsorArg)
		if err != nil {
			return fmt.Errorf("invalid point: %v", err)
		}
		pInfo, err := roller.Client.GetPoint(ctx, patp)
		if err != nil {
			return fmt.Errorf("error getting point: %v", err)
		}
		currentRevision := fmt.Sprintf("%v", pInfo.Network.Keys.Life)
		rev, err := strconv.Atoi(currentRevision)
		if err != nil {
			return fmt.Errorf("invalid life: %v", err)
		}
		if masterTicket != "" {
			wallet := keygen.GenerateWallet(masterTicket, uint32(pointInt), passphrase, uint(rev), true)
			privKey, err = crypto.HexToECDSA(wallet.Ownership.Keys.Private)
			if err != nil {
				return fmt.Errorf("invalid key material: %v", err)
			}
			addr = wallet.Ownership.Keys.Address
		} else if ethKey != "" {
			hexKey := strings.TrimPrefix(ethKey, "0x")
			privateKey, err := crypto.HexToECDSA(hexKey)
			if err != nil {
				return fmt.Errorf("invalid private key: %v", err)
			}
			addr = fmt.Sprintf("%v", crypto.PubkeyToAddress(privateKey.PublicKey))
		}
		keysTx, err := roller.Client.Escape(ctx, patp, sponsor, addr, privKey)
		if err != nil {
			return fmt.Errorf("error processing breach: %v", err)
		}
		jsonData, err := json.MarshalIndent(keysTx, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshaling response: %v", err)
		}
		fmt.Println(string(jsonData))
		return nil
	},
}

var ModCancelEscapeCmd = &cobra.Command{
	Use:   "cancel-escape",
	Short: "Cancel an escape from a sponsor",
	RunE: func(cmd *cobra.Command, args []string) error {
		var privKey *ecdsa.PrivateKey
		var addr string
		point, err := cmd.Flags().GetString("point")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		sponsorArg, err := cmd.Flags().GetString("sponsor")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		if point == "" {
			return fmt.Errorf("point is required")
		}
		if sponsorArg == "" {
			return fmt.Errorf("sponsor is required")
		}
		masterTicket, err := cmd.Flags().GetString("master-ticket")
		if err != nil {
			return fmt.Errorf("error getting master-ticket flag: %v", err)
		}
		ethKey, err := cmd.Flags().GetString("private-key")
		if err != nil {
			return fmt.Errorf("error getting private-key flag: %v", err)
		}
		if masterTicket == "" && ethKey == "" {
			return fmt.Errorf("master-ticket is required")
		}
		passphrase, err := cmd.Flags().GetString("passphrase")
		if err != nil {
			return fmt.Errorf("error getting passphrase flag: %v", err)
		}
		patp, pointInt, err := types.ValidateAndNormalizePatp(point)
		if err != nil {
			return fmt.Errorf("invalid point: %v", err)
		}
		sponsor, _, err := types.ValidateAndNormalizePatp(sponsorArg)
		if err != nil {
			return fmt.Errorf("invalid point: %v", err)
		}
		pInfo, err := roller.Client.GetPoint(ctx, patp)
		if err != nil {
			return fmt.Errorf("error getting point: %v", err)
		}
		currentRevision := fmt.Sprintf("%v", pInfo.Network.Keys.Life)
		rev, err := strconv.Atoi(currentRevision)
		if err != nil {
			return fmt.Errorf("invalid life: %v", err)
		}
		if masterTicket != "" {
			wallet := keygen.GenerateWallet(masterTicket, uint32(pointInt), passphrase, uint(rev), true)
			privKey, err = crypto.HexToECDSA(wallet.Ownership.Keys.Private)
			if err != nil {
				return fmt.Errorf("invalid key material: %v", err)
			}
			addr = wallet.Ownership.Keys.Address
		} else if ethKey != "" {
			hexKey := strings.TrimPrefix(ethKey, "0x")
			privKey, err = crypto.HexToECDSA(hexKey)
			if err != nil {
				return fmt.Errorf("invalid private key: %v", err)
			}
			addr = fmt.Sprintf("%v", crypto.PubkeyToAddress(privKey.PublicKey))
		}
		keysTx, err := roller.Client.CancelEscape(ctx, patp, sponsor, addr, privKey)
		if err != nil {
			return fmt.Errorf("error processing breach: %v", err)
		}
		jsonData, err := json.MarshalIndent(keysTx, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshaling response: %v", err)
		}
		fmt.Println(string(jsonData))
		return nil
	},
}

var ModAdoptCmd = &cobra.Command{
	Use:   "adopt",
	Short: "Adopt a point",
	RunE: func(cmd *cobra.Command, args []string) error {
		var privKey *ecdsa.PrivateKey
		var addr string
		point, err := cmd.Flags().GetString("point")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		sponsorArg, err := cmd.Flags().GetString("sponsor")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		if point == "" {
			return fmt.Errorf("point is required")
		}
		if sponsorArg == "" {
			return fmt.Errorf("point is required")
		}
		masterTicket, err := cmd.Flags().GetString("master-ticket")
		if err != nil {
			return fmt.Errorf("error getting master-ticket flag: %v", err)
		}
		ethKey, err := cmd.Flags().GetString("private-key")
		if err != nil {
			return fmt.Errorf("error getting private-key flag: %v", err)
		}
		if masterTicket == "" && ethKey == "" {
			return fmt.Errorf("master-ticket is required")
		}
		passphrase, err := cmd.Flags().GetString("passphrase")
		if err != nil {
			return fmt.Errorf("error getting passphrase flag: %v", err)
		}
		patp, pointInt, err := types.ValidateAndNormalizePatp(point)
		if err != nil {
			return fmt.Errorf("invalid point: %v", err)
		}
		sponsor, _, err := types.ValidateAndNormalizePatp(sponsorArg)
		if err != nil {
			return fmt.Errorf("invalid point: %v", err)
		}
		pInfo, err := roller.Client.GetPoint(ctx, patp)
		if err != nil {
			return fmt.Errorf("error getting point: %v", err)
		}
		currentRevision := fmt.Sprintf("%v", pInfo.Network.Keys.Life)
		rev, err := strconv.Atoi(currentRevision)
		if err != nil {
			return fmt.Errorf("invalid life: %v", err)
		}
		if masterTicket != "" {
			wallet := keygen.GenerateWallet(masterTicket, uint32(pointInt), passphrase, uint(rev), true)
			privKey, err = crypto.HexToECDSA(wallet.Ownership.Keys.Private)
			if err != nil {
				return fmt.Errorf("invalid key material: %v", err)
			}
			addr = wallet.Ownership.Keys.Address
		} else if ethKey != "" {
			hexKey := strings.TrimPrefix(ethKey, "0x")
			privKey, err = crypto.HexToECDSA(hexKey)
			if err != nil {
				return fmt.Errorf("invalid private key: %v", err)
			}
			addr = fmt.Sprintf("%v", crypto.PubkeyToAddress(privKey.PublicKey))
		}
		keysTx, err := roller.Client.Escape(ctx, patp, sponsor, addr, privKey)
		if err != nil {
			return fmt.Errorf("error processing breach: %v", err)
		}
		jsonData, err := json.MarshalIndent(keysTx, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshaling response: %v", err)
		}
		fmt.Println(string(jsonData))
		return nil
	},
}

var GetWalletCmd = &cobra.Command{
	Use:   "get-wallet",
	Short: "Generate a wallet from master ticket",
	RunE: func(cmd *cobra.Command, args []string) error {
		point, _ := cmd.Flags().GetString("point")
		if point == "" {
			return fmt.Errorf("point is required")
		}
		masterTicket, _ := cmd.Flags().GetString("master-ticket")
		if masterTicket == "" {
			return fmt.Errorf("master-ticket is required")
		}
		output, _ := cmd.Flags().GetString("output-dir")
		patp, pointInt, err := types.ValidateAndNormalizePatp(point)
		if err != nil {
			return fmt.Errorf("invalid point")
		}
		var revision string
		life, _ := cmd.Flags().GetString("life")
		if life != "" {
			revision = life
		} else {
			pInfo, err := roller.Client.GetPoint(ctx, patp)
			if err != nil {
				return fmt.Errorf("Error getting point: %v", err)
			}
			revision = fmt.Sprintf("%v", pInfo.Network.Keys.Life)
		}
		rev, err := strconv.Atoi(revision)
		walletData := keygen.GenerateWallet(masterTicket, uint32(pointInt), "", uint(rev), true)
		jsonData, err := json.MarshalIndent(types.WalletResp{Wallet: walletData}, "", "  ")
		if err != nil {
			return fmt.Errorf("Error marshaling response: %v", err)
		}
		fmt.Println(string(jsonData))
		pwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("error getting current directory: %v", err)
		}
		var outDir string
		if output != "" {
			outDir, err = validatePath(output)
			if err != nil {
				return fmt.Errorf("error validating output directory: %v", err)
			}
		} else {
			outDir = filepath.Join(pwd, "out")
			err = os.MkdirAll(outDir, 0755)
			if err != nil {
				return fmt.Errorf("error creating output directory: %v", err)
			}
		}
		p := strings.ReplaceAll(patp, "~", "")
		filename := filepath.Join(outDir, fmt.Sprintf("%s-wallet-%v.json", p, rev))
		err = os.WriteFile(filename, []byte(jsonData), 0600)
		if err != nil {
			return fmt.Errorf("error writing keyfile to disk: %v", err)
		}
		fmt.Printf("Wallet written to %s\n", filename)
		return nil
	},
}

var GetKeyfileCmd = &cobra.Command{
	Use:   "get-keyfile",
	Short: "Generate a keyfile from master ticket",
	RunE: func(cmd *cobra.Command, args []string) error {
		point, _ := cmd.Flags().GetString("point")
		if point == "" {
			return fmt.Errorf("point is required")
		}
		masterTicket, _ := cmd.Flags().GetString("master-ticket")
		if masterTicket == "" {
			return fmt.Errorf("master-ticket is required")
		}
		output, _ := cmd.Flags().GetString("output-dir")
		patp, pointInt, err := types.ValidateAndNormalizePatp(point)
		if err != nil {
			return fmt.Errorf("invalid point")
		}
		pInfo, err := roller.Client.GetPoint(ctx, patp)
		if err != nil {
			return fmt.Errorf("Error getting point: %v", err)
		}
		var rev string
		var lifeInt int
		life, _ := cmd.Flags().GetString("life")
		if life != "" {
			lifeInt, err = strconv.Atoi(rev)
			if err != nil {
				return fmt.Errorf("invalid life value: %v", err)
			}
		} else {
			rev = fmt.Sprintf("%v", pInfo.Network.Keys.Life)
			lifeInt, err = strconv.Atoi(rev)
			if err != nil {
				return fmt.Errorf("invalid life value: %v", err)
			}
		}
		lifeInt -= 1
		wallet := keygen.GenerateWallet(masterTicket, uint32(pointInt), "", uint(lifeInt), true)
		pointKey := strings.TrimPrefix(pInfo.Network.Keys.Crypt, "0x")
		if wallet.Network.Keys.Crypt.Public != pointKey {
			return fmt.Errorf("could not generate public key matching PKI; 0x%s / %s", wallet.Network.Keys.Crypt.Public, pInfo.Network.Keys.Crypt)
		}
		lifeInt += 1
		keyfile, err := roller.Keyfile(wallet.Network.Keys.Crypt.Private, wallet.Network.Keys.Auth.Private, patp, lifeInt)
		if err != nil {
			return fmt.Errorf("error generating keyfile: %v", err)
		}
		fmt.Println(keyfile)
		var outDir string
		if output != "" {
			outDir, err = validatePath(output)
			if err != nil {
				return fmt.Errorf("error validating output directory: %v", err)
			}
		} else {
			pwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("error getting current directory: %v", err)
			}
			outDir = filepath.Join(pwd, "out")
			err = os.MkdirAll(outDir, 0755)
			if err != nil {
				return fmt.Errorf("error creating output directory: %v", err)
			}
		}
		p := strings.ReplaceAll(patp, "~", "")
		filename := filepath.Join(outDir, fmt.Sprintf("%s-%d.key", p, lifeInt))
		err = os.WriteFile(filename, []byte(keyfile), 0600)
		if err != nil {
			return fmt.Errorf("error writing keyfile to disk: %v", err)
		}
		fmt.Printf("Keyfile written to %s\n", filename)
		return nil
	},
}

func validatePath(pathStr string) (string, error) {
	cleanPath := filepath.Clean(pathStr)
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		return "", fmt.Errorf("invalid path: %v", err)
	}
	if !filepath.IsAbs(absPath) {
		return "", fmt.Errorf("path must be absolute")
	}
	info, err := os.Stat(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("path does not exist")
		}
		return "", fmt.Errorf("error checking path: %v", err)
	}
	if !info.IsDir() {
		return "", fmt.Errorf("path must be a directory")
	}
	return filepath.Base(absPath), nil
}
