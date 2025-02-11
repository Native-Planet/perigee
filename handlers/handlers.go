package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Native-Planet/perigee/libprg"
	"github.com/Native-Planet/perigee/types"
	"go.uber.org/zap"
)

var adminToken = os.Getenv("ADMIN_TOKEN")

func parseLife(lifeStr string) (int, error) {
	if lifeStr == "" {
		return 0, nil
	}
	life, err := strconv.Atoi(lifeStr)
	if err != nil {
		return 0, fmt.Errorf("invalid life value: %v", err)
	}
	return life, nil
}

func LivenessProbe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ReadinessProbe(w http.ResponseWriter, r *http.Request) {
	if adminToken != "" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Admin-Token")
		if authHeader != adminToken && adminToken != "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func writeJSON(w http.ResponseWriter, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling response: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	return nil
}

func GetPoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		point := r.URL.Query().Get("point")
		resp, err := libprg.Point(point)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting point: %v", err), http.StatusInternalServerError)
			return
		}
		zap.L().Info("Get point", zap.String("point", point), zap.Any("info", resp))
		writeJSON(w, resp)
	}
}

func GetWallet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var life int
		point := r.URL.Query().Get("point")
		ticket := r.URL.Query().Get("ticket")
		life, err := parseLife(r.URL.Query().Get("life"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		passphrase := r.URL.Query().Get("passphrase")
		if point == "" || ticket == "" {
			http.Error(w, "Missing point or ticket parameter", http.StatusBadRequest)
			return
		}
		zap.L().Info("Generate wallet", zap.String("point", point), zap.Int("life", life))

		wallet, err := libprg.Wallet(point, ticket, passphrase, life)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error generating wallet: %v", err), http.StatusInternalServerError)
			return
		}
		writeJSON(w, types.WalletResp{Wallet: wallet})
	}
}

func GetPending() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		address := r.URL.Query().Get("address")
		if address == "" {
			address = r.URL.Query().Get("point")
		}
		resp, err := libprg.Pending(address)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting pending: %v", err), http.StatusInternalServerError)
			return
		}
		zap.L().Info("Get pending", zap.String("address", address))
		writeJSON(w, resp)
	}
}

func GetKeyfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		point := r.URL.Query().Get("point")
		ticket := r.URL.Query().Get("ticket")
		life, err := parseLife(r.URL.Query().Get("life"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		keyfile, err := libprg.Keyfile(point, ticket, "", life)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error generating keyfile: %v", err), http.StatusInternalServerError)
			return
		}

		writeJSON(w, types.ShipKeyfile{
			Point:   point,
			Keyfile: keyfile,
		})
	}
}

func handleTransaction(w http.ResponseWriter, r *http.Request, f func(string, string, string, string) (interface{}, error)) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	point := r.URL.Query().Get("point")
	targetPatp := r.URL.Query().Get("sponsor")
	if targetPatp == "" {
		targetPatp = r.URL.Query().Get("adoptee")
	}

	var ticket string
	if ethKey := r.URL.Query().Get("privkey"); ethKey != "" {
		ticket = strings.TrimPrefix(ethKey, "0x")
	} else {
		ticket = r.URL.Query().Get("ticket")
	}

	if point == "" || targetPatp == "" || ticket == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	passphrase := r.URL.Query().Get("passphrase")
	tx, err := f(point, targetPatp, ticket, passphrase)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing transaction: %v", err), http.StatusInternalServerError)
		return
	}

	zap.L().Info("Transaction processed", zap.String("point", point), zap.Any("tx", tx))
	writeJSON(w, tx)
}

func ModEscape() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleTransaction(w, r, libprg.Escape)
	}
}

func ModAdopt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleTransaction(w, r, libprg.Adopt)
	}
}

func ModCancelEscape() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleTransaction(w, r, libprg.CancelEscape)
	}
}

func ModBreach() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		point := r.URL.Query().Get("point")
		ticket := r.URL.Query().Get("ticket")
		passphrase := r.URL.Query().Get("passphrase")
		seed := r.URL.Query().Get("seed")

		tx, err := libprg.Breach(point, ticket, passphrase, seed)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error processing breach: %v", err), http.StatusInternalServerError)
			return
		}

		zap.L().Info("Breach processed", zap.String("point", point), zap.Any("tx", tx))
		writeJSON(w, tx)
	}
}
