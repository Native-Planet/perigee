package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Native-Planet/perigee/libprg"
	"github.com/spf13/cobra"
)

func NewSpawnCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "spawn",
		Short: "Spawn a child on the parent's L1 or L2",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			point, _ := cmd.Flags().GetString("point")
			child, _ := cmd.Flags().GetString("child")
			owner, _ := cmd.Flags().GetString("owner")
			ticket, _ := cmd.Flags().GetString("master-ticket")
			privateKey, _ := cmd.Flags().GetString("private-key")
			passphrase, _ := cmd.Flags().GetString("passphrase")
			if point == "" || child == "" || owner == "" {
				return fmt.Errorf("point, child, and owner are required")
			}
			if (ticket == "") == (privateKey == "") {
				return fmt.Errorf("provide exactly one of master-ticket or private-key")
			}
			if privateKey != "" {
				ticket = strings.TrimPrefix(privateKey, "0x")
			}
			tx, err := libprg.Spawn(cmd.Context(), point, ticket, passphrase, child, owner)
			if err != nil {
				return fmt.Errorf("spawn: %w", err)
			}
			encoder := json.NewEncoder(cmd.OutOrStdout())
			encoder.SetIndent("", "  ")
			return encoder.Encode(tx)
		},
	}
	cmd.Flags().String("point", "", "Parent galaxy or star")
	cmd.Flags().String("child", "", "Child ship to spawn")
	cmd.Flags().String("owner", "", "Recipient Ethereum address")
	cmd.Flags().String("master-ticket", "", "Parent master ticket")
	cmd.Flags().String("private-key", "", "Parent owner or spawn proxy private key")
	cmd.Flags().String("passphrase", "", "Parent wallet passphrase")
	return cmd
}
