package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Native-Planet/perigee/libprg"
	"github.com/Native-Planet/perigee/types"

	"github.com/spf13/cobra"
)

var GetPointCmd = &cobra.Command{
	Use:   "get-point",
	Short: "Retrieve point information",
	RunE: func(cmd *cobra.Command, args []string) error {
		point, _ := cmd.Flags().GetString("point")
		if point == "" {
			return fmt.Errorf("point is required")
		}
		resp, err := libprg.Point(point)
		if err != nil {
			return fmt.Errorf("error getting point: %v", err)
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
		var addr string
		if point != "" {
			addr = point
		} else if address != "" {
			addr = address
		}
		resp, err := libprg.Pending(addr)
		if err != nil {
			return fmt.Errorf("error retrieving pending: %v", err)
		}
		jsonData, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshaling response: %v", err)
		}
		fmt.Println(string(jsonData))
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
			if !cmd.Flags().Changed("private-key") {
				return fmt.Errorf("master-ticket is required")
			} else {
				ethKey, err := cmd.Flags().GetString("private-key")
				if err != nil {
					return fmt.Errorf("error getting private-key flag: %v", err)
				}
				masterTicket = ethKey
			}
		}
		passphrase, err := cmd.Flags().GetString("passphrase")
		if err != nil {
			return fmt.Errorf("error getting passphrase flag: %v", err)
		}
		keysTx, err := libprg.Breach(point, masterTicket, passphrase)
		if err != nil {
			return fmt.Errorf("error processing breach: %v", err)
		}
		jsonData, err := json.MarshalIndent(keysTx, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshaling response: %v", err)
		}
		fmt.Println(string(jsonData))
		if cmd.Flags().Changed("wait") {
			duration, err := cmd.Flags().GetDuration("wait")
			if err != nil {
				return fmt.Errorf("error getting wait flag: %v", err)
			}
			if err = libprg.Wait(duration, keysTx.Signature); err != nil {
				return fmt.Errorf("error waiting for keys transaction: %v", err)
			}
		}
		return nil
	},
}

var ModEscapeCmd = &cobra.Command{
	Use:   "escape",
	Short: "Escape from a sponsor",
	RunE: func(cmd *cobra.Command, args []string) error {
		point, err := cmd.Flags().GetString("point")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		sponsor, err := cmd.Flags().GetString("sponsor")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		if point == "" {
			return fmt.Errorf("point is required")
		}
		if sponsor == "" {
			return fmt.Errorf("sponsor is required")
		}
		masterTicket, err := cmd.Flags().GetString("master-ticket")
		if err != nil {
			return fmt.Errorf("error getting master-ticket flag: %v", err)
		}
		if masterTicket == "" {
			if !cmd.Flags().Changed("private-key") {
				return fmt.Errorf("master-ticket is required")
			} else {
				ethKey, err := cmd.Flags().GetString("private-key")
				if err != nil {
					return fmt.Errorf("error getting private-key flag: %v", err)
				}
				masterTicket = ethKey
			}
		}
		passphrase, err := cmd.Flags().GetString("passphrase")
		if err != nil {
			return fmt.Errorf("error getting passphrase flag: %v", err)
		}
		keysTx, err := libprg.Escape(point, sponsor, masterTicket, passphrase)
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
	Short: "Cancel an escape request",
	RunE: func(cmd *cobra.Command, args []string) error {
		point, err := cmd.Flags().GetString("point")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		sponsor, err := cmd.Flags().GetString("sponsor")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		if point == "" {
			return fmt.Errorf("point is required")
		}
		if sponsor == "" {
			return fmt.Errorf("sponsor is required")
		}
		masterTicket, err := cmd.Flags().GetString("master-ticket")
		if err != nil {
			return fmt.Errorf("error getting master-ticket flag: %v", err)
		}
		if masterTicket == "" {
			if !cmd.Flags().Changed("private-key") {
				return fmt.Errorf("master-ticket is required")
			} else {
				ethKey, err := cmd.Flags().GetString("private-key")
				if err != nil {
					return fmt.Errorf("error getting private-key flag: %v", err)
				}
				masterTicket = ethKey
			}
		}
		passphrase, err := cmd.Flags().GetString("passphrase")
		if err != nil {
			return fmt.Errorf("error getting passphrase flag: %v", err)
		}
		keysTx, err := libprg.CancelEscape(point, sponsor, masterTicket, passphrase)
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
	Short: "Adopt a point as sponsor",
	RunE: func(cmd *cobra.Command, args []string) error {
		point, err := cmd.Flags().GetString("point")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		adoptee, err := cmd.Flags().GetString("sponsor")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		if point == "" {
			return fmt.Errorf("point is required")
		}
		if adoptee == "" {
			return fmt.Errorf("adoptee is required")
		}
		masterTicket, err := cmd.Flags().GetString("master-ticket")
		if err != nil {
			return fmt.Errorf("error getting master-ticket flag: %v", err)
		}
		if masterTicket == "" {
			if !cmd.Flags().Changed("private-key") {
				return fmt.Errorf("master-ticket is required")
			} else {
				ethKey, err := cmd.Flags().GetString("private-key")
				if err != nil {
					return fmt.Errorf("error getting private-key flag: %v", err)
				}
				masterTicket = ethKey
			}
		}
		passphrase, err := cmd.Flags().GetString("passphrase")
		if err != nil {
			return fmt.Errorf("error getting passphrase flag: %v", err)
		}
		keysTx, err := libprg.Adopt(point, adoptee, masterTicket, passphrase)
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
		output, err := cmd.Flags().GetString("output-dir")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		point, err := cmd.Flags().GetString("point")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		if point == "" {
			return fmt.Errorf("point is required")
		}
		life, err := cmd.Flags().GetInt("life")
		if err != nil {
			return fmt.Errorf("error getting life flag: %v", err)
		}
		masterTicket, err := cmd.Flags().GetString("master-ticket")
		if err != nil {
			return fmt.Errorf("error getting master-ticket flag: %v", err)
		}
		if masterTicket == "" {
			if !cmd.Flags().Changed("private-key") {
				return fmt.Errorf("master-ticket is required")
			} else {
				ethKey, err := cmd.Flags().GetString("private-key")
				if err != nil {
					return fmt.Errorf("error getting private-key flag: %v", err)
				}
				masterTicket = ethKey
			}
		}
		passphrase, err := cmd.Flags().GetString("passphrase")
		if err != nil {
			return fmt.Errorf("error getting passphrase flag: %v", err)
		}
		walletData, err := libprg.Wallet(point, masterTicket, passphrase, life)
		if err != nil {
			return fmt.Errorf("error generating wallet: %v", err)
		}
		jsonData, err := json.MarshalIndent(types.WalletResp{Wallet: walletData}, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshaling response: %v", err)
		}
		fmt.Println(string(jsonData))
		return writeToFile([]byte(jsonData), output, point, life, "", "-wallet.json")
	},
}

var GetKeyfileCmd = &cobra.Command{
	Use:   "get-keyfile",
	Short: "Generate a keyfile from master ticket",
	RunE: func(cmd *cobra.Command, args []string) error {
		output, err := cmd.Flags().GetString("output-dir")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		point, err := cmd.Flags().GetString("point")
		if err != nil {
			return fmt.Errorf("error getting point flag: %v", err)
		}
		if point == "" {
			return fmt.Errorf("point is required")
		}
		life, err := cmd.Flags().GetInt("life")
		if err != nil {
			return fmt.Errorf("error getting life flag: %v", err)
		}
		if life == 0 {
			pInfo, err := libprg.Point(point)
			if err != nil {
				return fmt.Errorf("error retrieving point: %v", err)
			}
			life, err = strconv.Atoi(pInfo.Point.Network.Keys.Life)
			if err != nil {
				return fmt.Errorf("invalid life: %v", err)
			}
		}
		masterTicket, err := cmd.Flags().GetString("master-ticket")
		if err != nil {
			return fmt.Errorf("error getting master-ticket flag: %v", err)
		}
		if masterTicket == "" {
			if !cmd.Flags().Changed("private-key") {
				return fmt.Errorf("master-ticket is required")
			} else {
				ethKey, err := cmd.Flags().GetString("private-key")
				if err != nil {
					return fmt.Errorf("error getting private-key flag: %v", err)
				}
				masterTicket = ethKey
			}
		}
		passphrase, err := cmd.Flags().GetString("passphrase")
		if err != nil {
			return fmt.Errorf("error getting passphrase flag: %v", err)
		}
		keyfile, err := libprg.Keyfile(point, masterTicket, passphrase, life)
		if err != nil {
			return fmt.Errorf("error generating wallet: %v", err)
		}
		fmt.Println(keyfile)
		return writeToFile([]byte(keyfile), output, point, life, "", ".key")
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

func writeToFile(data []byte, outputDir, point string, life int, prefix, ext string) error {
	var outDir string
	var err error
	if outputDir != "" {
		outDir, err = validatePath(outputDir)
		if err != nil {
			return fmt.Errorf("error validating output directory: %v", err)
		}
	} else {
		pwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("error getting current directory: %v", err)
		}
		outDir = filepath.Join(pwd, "out")
		if err := os.MkdirAll(outDir, 0755); err != nil {
			return fmt.Errorf("error creating output directory: %v", err)
		}
	}
	p := strings.TrimPrefix(point, "~")
	filename := filepath.Join(outDir, fmt.Sprintf("%s-%s%v%s", p, prefix, life, ext))

	if err := os.WriteFile(filename, data, 0600); err != nil {
		return fmt.Errorf("error writing file to disk: %v", err)
	}
	fmt.Printf("File written to %s\n", filename)
	return nil
}
