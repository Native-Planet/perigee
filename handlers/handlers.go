package handlers

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"perigee/roller"
	"perigee/types"

	"github.com/deelawn/urbit-gob/co"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nathanlever/keygen"
	"go.uber.org/zap"
)

var (
	// use this to set auth headers
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
		if authHeader != adminToken && adminToken != "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func GetWallet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		shipParam := r.URL.Query().Get("point")
		ticket := r.URL.Query().Get("ticket")
		revisionParam := r.URL.Query().Get("life")
		passphrase := r.URL.Query().Get("passphrase")
		if shipParam == "" || ticket == "" {
			http.Error(w, "Missing 'point' or 'ticket' query parameter", http.StatusBadRequest)
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
		zap.L().Info("Generate wallet", zap.String("point", shipParam), zap.Int("life", revision))
		walletData := keygen.GenerateWallet(ticket, uint32(bigPoint.Int64()), passphrase, uint(revision), true)
		jsonData, err := json.MarshalIndent(types.WalletResp{Wallet: walletData}, "", "  ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshaling response: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func GetPoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		shipParam := r.URL.Query().Get("point")
		patp, _, err := types.ParsePointParam(shipParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading @p: %v", err), http.StatusInternalServerError)
			return
		}
		pointInfo, err := roller.Client.GetPoint(r.Context(), patp)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting point: %v", err), http.StatusInternalServerError)
			return
		}
		zap.L().Info("Get point", zap.String("point", patp), zap.Any("info", pointInfo))
		resp := types.PointResp{
			Point:    pointInfo,
			PatpName: patp,
		}
		if err := resp.Point.ResolveSponsorPatp(); err != nil {
			http.Error(w, fmt.Sprintf("Error resolving sponsor: %v", err), http.StatusInternalServerError)
			return
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

func GetPending() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
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
		zap.L().Info("Get pending", zap.String("address", string(jsonData)))
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func GetKeyfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		patp := r.URL.Query().Get("point")
		ticket := r.URL.Query().Get("ticket")
		rev := r.URL.Query().Get("life")
		point, err := roller.Client.GetPoint(r.Context(), patp)
		if err != nil {
			http.Error(w, fmt.Sprintf("error retrieving point info: %v", err), http.StatusInternalServerError)
			return
		}
		bigPoint, err := co.Patp2Point(patp)
		if err != nil {
			http.Error(w, fmt.Sprintf("error converting point: %v", err), http.StatusInternalServerError)
			return
		}
		var life int
		if rev == "" {
			revInt, err := strconv.Atoi(point.Network.Keys.Life)
			if err != nil {
				http.Error(w, fmt.Sprintf("error casting azimuth life: %v", err), http.StatusInternalServerError)
				return
			}
			life = revInt - 1
		} else {
			revInt, err := strconv.Atoi(rev)
			if err != nil {
				http.Error(w, fmt.Sprintf("error casting azimuth life: %v", err), http.StatusInternalServerError)
				return
			}
			life = revInt
		}
		wallet := keygen.GenerateWallet(ticket, uint32(bigPoint.Uint64()), "", uint(life), true)
		if err != nil {
			http.Error(w, fmt.Sprintf("error generating wallet: %v", err), http.StatusInternalServerError)
			return
		}
		pointKey := strings.TrimPrefix(point.Network.Keys.Crypt, "0x")
		if wallet.Network.Keys.Crypt.Public != pointKey && rev == "" {
			http.Error(w, fmt.Sprintf("could not generate matching keyfile; 0x%s / %s", wallet.Network.Keys.Crypt.Public, point.Network.Keys.Crypt), http.StatusInternalServerError)
			return
		}
		keyfile, err := roller.Keyfile(wallet.Network.Keys.Crypt.Private, wallet.Network.Keys.Auth.Private, patp, life)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not generate keyfile: %v", err), http.StatusInternalServerError)
		}
		resp := types.ShipKeyfile{
			Point:   patp,
			Keyfile: keyfile,
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

func ModBreach() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		shipParam := r.URL.Query().Get("point")
		ticket := r.URL.Query().Get("ticket")
		life := r.URL.Query().Get("life")
		passphrase := r.URL.Query().Get("passphrase")
		patp, p, err := types.ParsePointParam(shipParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid point: %v", err), http.StatusBadRequest)
			return
		}
		pInfo, err := roller.Client.GetPoint(r.Context(), patp)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid point: %v", err), http.StatusBadRequest)
			return
		}
		var rev int
		if life == "" {
			rev, err = strconv.Atoi(pInfo.Network.Keys.Life)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid life %v", err), http.StatusInternalServerError)
				return
			}
		} else {
			rev, err = strconv.Atoi(life)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid life %v", err), http.StatusInternalServerError)
				return
			}
		}
		wallet := keygen.GenerateWallet(ticket, uint32(p), passphrase, uint(rev), true)
		privKey, err := crypto.HexToECDSA(wallet.Ownership.Keys.Private)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid key material: %v", err), http.StatusInternalServerError)
			return
		}
		keysTx, err := roller.Client.ConfigureKeys(r.Context(), patp, "0x"+wallet.Network.Keys.Crypt.Public, "0x"+wallet.Network.Keys.Auth.Public, true, wallet.Ownership.Keys.Address, privKey)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error updating breach: %v", err), http.StatusInternalServerError)
			return
		}
		zap.L().Info("Update breach", zap.String("point", patp), zap.Any("tx", keysTx))
		jsonData, err := json.MarshalIndent(keysTx, "", "  ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshaling response: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func ModEscape() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		shipParam := r.URL.Query().Get("point")
		ticket := r.URL.Query().Get("ticket")
		passphrase := r.URL.Query().Get("passphrase")
		sponsorParam := r.URL.Query().Get("sponsor")
		ethKey := r.URL.Query().Get("privkey")
		patp, p, err := types.ParsePointParam(shipParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading @p: %v", err), http.StatusInternalServerError)
			return
		}
		sponsor, _, err := types.ParsePointParam(sponsorParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading @p: %v", err), http.StatusInternalServerError)
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
		var privKey *ecdsa.PrivateKey
		var wallet keygen.Wallet
		var keysTx *types.Transaction
		if ticket != "" {
			wallet = keygen.GenerateWallet(ticket, uint32(p), passphrase, uint(rev), true)
			privKey, err = crypto.HexToECDSA(wallet.Ownership.Keys.Private)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid key material: %v", err), http.StatusInternalServerError)
				return
			}
			keysTx, err = roller.Client.Escape(r.Context(), patp, sponsor, wallet.Ownership.Keys.Address, privKey)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error updating escape: %v", err), http.StatusInternalServerError)
				return
			}
		} else if ethKey != "" {
			hexKey := strings.TrimPrefix(ethKey, "0x")
			privateKey, err := crypto.HexToECDSA(hexKey)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid private key: %v", err), http.StatusInternalServerError)
				return
			}
			address := fmt.Sprintf("%v", crypto.PubkeyToAddress(privateKey.PublicKey))
			keysTx, err = roller.Client.Escape(r.Context(), patp, sponsor, address, privKey)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error updating escape: %v", err), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, fmt.Sprintf("Must provide master ticket or eth private key: %v", err), http.StatusInternalServerError)
			return
		}
		zap.L().Info("Update escape", zap.String("point", patp), zap.Any("tx", keysTx))
		jsonData, err := json.MarshalIndent(keysTx, "", "  ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshaling response: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func ModAdopt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		shipParam := r.URL.Query().Get("point")
		ticket := r.URL.Query().Get("ticket")
		passphrase := r.URL.Query().Get("passphrase")
		adopteeParam := r.URL.Query().Get("adoptee")
		ethKey := r.URL.Query().Get("privkey")
		patp, p, err := types.ParsePointParam(shipParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading @p: %v", err), http.StatusInternalServerError)
			return
		}
		adoptee, _, err := types.ParsePointParam(adopteeParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading @p: %v", err), http.StatusInternalServerError)
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
		var privKey *ecdsa.PrivateKey
		var wallet keygen.Wallet
		var keysTx *types.Transaction
		if ticket != "" {
			wallet = keygen.GenerateWallet(ticket, uint32(p), passphrase, uint(rev), true)
			privKey, err = crypto.HexToECDSA(wallet.Ownership.Keys.Private)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid key material: %v", err), http.StatusInternalServerError)
				return
			}
			zap.L().Info("adopt info", zap.String("point", patp), zap.Any("adoptee", adoptee), zap.Any("address", wallet.Ownership.Keys.Address), zap.Any("privkey", privKey))
			keysTx, err = roller.Client.Adopt(r.Context(), patp, adoptee, wallet.Ownership.Keys.Address, privKey)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error updating adopt: %v", err), http.StatusInternalServerError)
				return
			}
		} else if ethKey != "" {
			hexKey := strings.TrimPrefix(ethKey, "0x")
			privateKey, err := crypto.HexToECDSA(hexKey)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid private key: %v", err), http.StatusInternalServerError)
				return
			}
			address := fmt.Sprintf("%v", crypto.PubkeyToAddress(privateKey.PublicKey))
			keysTx, err = roller.Client.Adopt(r.Context(), patp, adoptee, address, privateKey)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error updating adopt: %v", err), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, fmt.Sprintf("Must provide master ticket or eth private key: %v", err), http.StatusInternalServerError)
			return
		}
		zap.L().Info("Update adopt", zap.String("point", patp), zap.Any("tx", keysTx))
		jsonData, err := json.MarshalIndent(keysTx, "", "  ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshaling response: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func ModCancelEscape() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		shipParam := r.URL.Query().Get("point")
		ticket := r.URL.Query().Get("ticket")
		passphrase := r.URL.Query().Get("passphrase")
		sponsorParam := r.URL.Query().Get("sponsor")
		ethKey := r.URL.Query().Get("privkey")
		patp, p, err := types.ParsePointParam(shipParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading @p: %v", err), http.StatusInternalServerError)
			return
		}
		sponsor, _, err := types.ParsePointParam(sponsorParam)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading @p: %v", err), http.StatusInternalServerError)
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
		var privKey *ecdsa.PrivateKey
		var wallet keygen.Wallet
		var keysTx *types.Transaction
		if ticket != "" {
			wallet = keygen.GenerateWallet(ticket, uint32(p), passphrase, uint(rev), true)
			privKey, err = crypto.HexToECDSA(wallet.Ownership.Keys.Private)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid key material: %v", err), http.StatusInternalServerError)
				return
			}
			zap.L().Info("adopt info", zap.String("point", patp), zap.Any("adoptee", sponsor), zap.Any("address", wallet.Ownership.Keys.Address), zap.Any("privkey", privKey))
			keysTx, err = roller.Client.CancelEscape(r.Context(), patp, sponsor, wallet.Ownership.Keys.Address, privKey)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error updating adopt: %v", err), http.StatusInternalServerError)
				return
			}
		} else if ethKey != "" {
			hexKey := strings.TrimPrefix(ethKey, "0x")
			privateKey, err := crypto.HexToECDSA(hexKey)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid private key: %v", err), http.StatusInternalServerError)
				return
			}
			address := fmt.Sprintf("%v", crypto.PubkeyToAddress(privateKey.PublicKey))
			keysTx, err = roller.Client.CancelEscape(r.Context(), patp, sponsor, address, privateKey)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error updating adopt: %v", err), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, fmt.Sprintf("Must provide master ticket or eth private key: %v", err), http.StatusInternalServerError)
			return
		}
		zap.L().Info("Update escape", zap.String("point", patp), zap.Any("tx", keysTx))
		jsonData, err := json.MarshalIndent(keysTx, "", "  ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshaling response: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
