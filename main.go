package main

import (
	"fmt"
	"net/http"
	"os"

	"perigee/cli"
	"perigee/handlers"
	"perigee/logger"
	"perigee/roller"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var port string

var rootCmd = &cobra.Command{
	Use:   "perigee",
	Short: "Perigee: Azimuth light client and macro server",
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runServer()
	},
}

func init() {
	// add flags to commands
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to run server on")
	cli.GetPointCmd.Flags().String("point", "", "Point for retrieval")

	cli.GetPendingCmd.Flags().String("point", "", "Pending transactions for a point")
	cli.GetPendingCmd.Flags().String("address", "", "Pending transactions for a address")

	cli.GetWalletCmd.Flags().String("master-ticket", "", "Master ticket for wallet generation")
	cli.GetWalletCmd.Flags().String("point", "", "Azimuth point for wallet generation")
	cli.GetWalletCmd.Flags().String("life", "", "Custom life value for wallet generation")

	cli.GetKeyfileCmd.Flags().String("master-ticket", "", "Master ticket for keyfile generation")
	cli.GetKeyfileCmd.Flags().String("point", "", "Azimuth point for keyfile generation")
	cli.GetKeyfileCmd.Flags().String("life", "", "Custom life value for keyfile generation (optional)")

	cli.ModBreachCmd.Flags().String("point", "", "Point for breach")
	cli.ModBreachCmd.Flags().String("master-ticket", "", "Master ticket for wallet generation")
	cli.ModBreachCmd.Flags().String("passphrase", "", "Passphrase for wallet")
	cli.ModBreachCmd.Flags().Duration("wait", 0, "Wait until breach processes (forever if no value, or specific time like '70m' or '1h20m')")

	cli.ModEscapeCmd.Flags().String("point", "", "Point for breach")
	cli.ModEscapeCmd.Flags().String("master-ticket", "", "Master ticket for wallet generation")
	cli.ModEscapeCmd.Flags().String("private-key", "", "Ethereum address private key")
	cli.ModEscapeCmd.Flags().String("passphrase", "", "Passphrase for wallet")
	cli.ModEscapeCmd.Flags().String("sponsor", "", "New sponsor")
	cli.ModEscapeCmd.Flags().Duration("wait", 0, "Wait until breach processes (forever if no value, or specific time like '70m' or '1h20m')")

	cli.ModCancelEscapeCmd.Flags().String("point", "", "Point for breach")
	cli.ModCancelEscapeCmd.Flags().String("master-ticket", "", "Master ticket for wallet generation")
	cli.ModCancelEscapeCmd.Flags().String("private-key", "", "Ethereum address private key")
	cli.ModCancelEscapeCmd.Flags().String("passphrase", "", "Passphrase for wallet")
	cli.ModCancelEscapeCmd.Flags().String("sponsor", "", "New sponsor")

	cli.ModAdoptCmd.Flags().String("point", "", "Point for breach")
	cli.ModAdoptCmd.Flags().String("master-ticket", "", "Master ticket for wallet generation")
	cli.ModAdoptCmd.Flags().String("private-key", "", "Ethereum address private key")
	cli.ModAdoptCmd.Flags().String("passphrase", "", "Passphrase for wallet")
	cli.ModAdoptCmd.Flags().String("adoptee", "", "Newly adopted point")
	cli.ModAdoptCmd.Flags().Duration("wait", 0, "Wait until breach processes (forever if no value, or specific time like '70m' or '1h20m')")

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(cli.GetWalletCmd)
	rootCmd.AddCommand(cli.GetKeyfileCmd)
	rootCmd.AddCommand(cli.GetPointCmd)
	rootCmd.AddCommand(cli.GetPendingCmd)
	rootCmd.AddCommand(cli.ModBreachCmd)
	rootCmd.AddCommand(cli.ModEscapeCmd)
	rootCmd.AddCommand(cli.ModCancelEscapeCmd)
	rootCmd.AddCommand(cli.ModAdoptCmd)
}

func runServer() error {

	http.HandleFunc("/v1/get/wallet", handlers.Auth(handlers.GetWallet()))
	http.HandleFunc("/v1/get/point", handlers.Auth(handlers.GetPoint()))
	http.HandleFunc("/v1/get/pending", handlers.Auth(handlers.GetPending()))
	http.HandleFunc("/v1/mod/breach", handlers.Auth(handlers.ModBreach()))
	http.HandleFunc("/v1/mod/escape", handlers.Auth(handlers.ModEscape()))
	http.HandleFunc("/v1/mod/cancel-escape", handlers.Auth(handlers.ModCancelEscape()))
	http.HandleFunc("/v1/mod/adopt", handlers.Auth(handlers.ModAdopt()))

	http.HandleFunc("/healthz", handlers.ReadinessProbe)
	http.HandleFunc("/readyz", handlers.LivenessProbe)

	zap.L().Info(fmt.Sprintf("Using roller endpoint %s", roller.RollerURL))
	addr := fmt.Sprintf(":%s", port)
	zap.L().Info(fmt.Sprintf("Starting server on %s...", addr))
	return http.ListenAndServe(addr, nil)
}

func main() {
	if err := logger.Init(); err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	defer logger.Sync()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
