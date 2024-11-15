package main

import (
	"fmt"
	"net/http"

	"perigee/handlers"
	"perigee/logger"
)

func main() {
	if err := logger.Init(); err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	defer logger.Sync()
	log := logger.GetLogger()

	// generate a wallet from master ticket
	http.HandleFunc("/v1/generate/wallet", handlers.Auth(handlers.GenerateWallet(log)))
	// retrieve point information
	http.HandleFunc("/v1/get/point", handlers.Auth(handlers.GetPoint(log)))
	// list pending rollup queue
	http.HandleFunc("/v1/get/pending", handlers.Auth(handlers.GetPending(log)))
	// breach a point
	http.HandleFunc("/v1/modify/breach", handlers.Auth(handlers.ModBreach(log)))
	http.HandleFunc("/healthz", handlers.ReadinessProbe)
	http.HandleFunc("/readyz", handlers.LivenessProbe)

	log.Info("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
