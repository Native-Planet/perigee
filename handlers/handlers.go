package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"perigee/roller"
	"perigee/types"
	"perigee/wallet"

	"github.com/deelawn/urbit-gob/co"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

var (
	adminToken = os.Getenv("ADMIN_TOKEN")
)

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
		if authHeader != adminToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func GenerateWallet(log *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shipParam := r.URL.Query().Get("ship")
		ticket := r.URL.Query().Get("ticket")
		revisionParam := r.URL.Query().Get("revision")
		if shipParam == "" || ticket == "" {
			http.Error(w, "Missing 'ship' or 'ticket' query parameter", http.StatusBadRequest)
			return
		}
		bigPoint, err := co.Patp2Point(shipParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid ship: %v", err), http.StatusBadRequest)
			return
		}
		revision := 0
		if revisionParam != "" {
			revision, err = strconv.Atoi(revisionParam)
			if err != nil {
				http.Error(w, "Invalid revision parameter", http.StatusBadRequest)
				return
			}
		}
		log.Info("Generate wallet", zap.String("ship", shipParam), zap.Int("revision", revision))
		walletData, err := wallet.GenerateWallet(types.WalletConfig{
			Ticket:   ticket,
			Ship:     uint32(bigPoint.Int64()),
			Revision: uint32(revision),
			Boot:     true,
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error generating wallet: %v", err), http.StatusInternalServerError)
			return
		}
		jsonData, err := json.MarshalIndent(types.WalletResp{Wallet: walletData}, "", "  ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshaling response: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func GetPoint(log *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PointReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding request: %v", err), http.StatusBadRequest)
			return
		}
		patp, _, err := types.UnmarshalPoint(req.Point)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid point: %v", err), http.StatusBadRequest)
			return
		}
		pointInfo, err := roller.Client.GetPoint(r.Context(), patp)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting point: %v", err), http.StatusInternalServerError)
			return
		}
		log.Info("Get point", zap.String("point", patp), zap.Any("info", pointInfo))
		resp := types.PointResp{
			Point:    pointInfo,
			PatpName: patp,
		}
		jsonData, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshaling response: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func GetPending(log *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ships, err := roller.Client.GetAllPending(r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting pending ships: %v", err), http.StatusInternalServerError)
			return
		}
		jsonData, err := json.MarshalIndent(ships, "", "  ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshaling response: %v", err), http.StatusInternalServerError)
			return
		}
		log.Info("Get pending", zap.String("address", string(jsonData)))
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func ModBreach(log *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BreachReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding request: %v", err), http.StatusBadRequest)
			return
		}
		patp, p, err := types.UnmarshalPoint(req.Point)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid point: %v", err), http.StatusBadRequest)
			return
		}
		pInfo, err := roller.Client.GetPoint(r.Context(), patp)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting point: %v", err), http.StatusInternalServerError)
			return
		}
		rev, err := strconv.Atoi(pInfo.Network.Keys.Life)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid life %v", err), http.StatusInternalServerError)
			return
		}
		rev += 1
		wallet, err := wallet.GenerateWallet(types.WalletConfig{
			Ticket:     req.Ticket,
			Ship:       uint32(p),
			Passphrase: req.Passphrase,
			Revision:   uint32(rev),
			Boot:       true,
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error generating point info: %v", err), http.StatusInternalServerError)
			return
		}
		privKey, err := crypto.HexToECDSA(wallet.Ownership.Keys.Private)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid key material: %v", err), http.StatusInternalServerError)
			return
		}
		keysTx, err := roller.Client.ConfigureKeys(r.Context(), patp, pInfo.Network.Keys.Crypt, pInfo.Network.Keys.Auth, true, wallet.Ownership.Keys.Address, privKey)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error updating breach: %v", err), http.StatusInternalServerError)
			return
		}
		log.Info("Update breach", zap.String("point", patp), zap.Any("tx", keysTx))
		jsonData, err := json.MarshalIndent(keysTx, "", "  ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshaling response: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
