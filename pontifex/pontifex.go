package pontifex

import (
	"crypto/rand"
	"embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	html "html/template"
	"io/fs"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/Native-Planet/perigee/libprg"
	"github.com/Native-Planet/perigee/types"
	"github.com/golang-jwt/jwt/v5"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/securecookie"
)

//go:embed templates/*
var templateFS embed.FS

//go:embed templates/static/*
var staticFS embed.FS
var (
	funcMap     template.FuncMap
	funcMapOnce sync.Once
	sessionKey  = []byte{}
	ethProvider = os.Getenv("ETH_PROVIDER")
)

type Handler struct {
	tmpl *template.Template
}

func init() {
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		panic(fmt.Sprintf("error while generating random string: %s", err))
	}
	sessionKey = buf
}

func GetTemplateFuncs() template.FuncMap {
	funcMapOnce.Do(func() {
		funcMap = template.FuncMap{
			"genSigil": func(patp string, size int) string {
				svg, err := GenerateSigil(types.PxSvgConfig{
					Point: patp,
					Size:  size,
				})
				if err != nil {
					return ""
				}
				return svg
			},
			"smallSigil": func(patp string) string {
				svg, err := GenerateSigil(types.PxSvgConfig{
					Point: patp,
					Size:  32,
				})
				if err != nil {
					return ""
				}
				return svg
			},
			"largeSigil": func(patp string) string {
				svg, err := GenerateSigil(types.PxSvgConfig{
					Point: patp,
					Size:  256,
				})
				if err != nil {
					return ""
				}
				return svg
			},
		}
	})
	return funcMap
}

func setupTemplates() (*template.Template, error) {
	tmpl := template.New("")
	tmpl = tmpl.Funcs(template.FuncMap{
		"genSigil": func(patp string, size int) string {
			svg, err := GenerateSigil(types.PxSvgConfig{
				Point: patp,
				Size:  size,
			})
			if err != nil {
				return ""
			}
			return svg
		},
		"smallSigil": func(patp string) string {
			svg, err := GenerateSigil(types.PxSvgConfig{
				Point: patp,
				Size:  32,
			})
			if err != nil {
				return ""
			}
			return svg
		},
		"largeSigil": func(patp string) string {
			svg, err := GenerateSigil(types.PxSvgConfig{
				Point: patp,
				Size:  256,
			})
			if err != nil {
				return ""
			}
			return svg
		},
		"safeHTML": func(s string) html.HTML {
			return html.HTML(s)
		},
	})
	parsed, err := tmpl.ParseFS(templateFS, "templates/*.html")
	if err != nil {
		return nil, fmt.Errorf("parsing templates: %w", err)
	}
	return parsed, nil
}

func CreateSession(w http.ResponseWriter, ship, ticket string, point types.PointResp, authType, address string) (string, error) {
	clientHalf := sessionKey[:16]
	serverHalf := sessionKey[16:]
	claims := jwt.MapClaims{
		"ship":     ship,
		"ticket":   ticket,
		"point":    point,
		"authType": authType,
		"wallet":   address,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(serverHalf)
	if err != nil {
		return "", fmt.Errorf("failed to create token: %v", err)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session_key",
		Value:    base64.StdEncoding.EncodeToString(clientHalf),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600,
	})
	w.Header().Set("X-Session-Token", tokenString)
	return tokenString, nil
}

func ValidateSession(r *http.Request, jwtString string) (*types.PxSession, error) {
	cookie, err := r.Cookie("session_key")
	if err != nil {
		return nil, fmt.Errorf("no session key cookie: %v", err)
	}
	clientHalf, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil, fmt.Errorf("invalid session key: %v", err)
	}
	fullKey := make([]byte, 32)
	copy(fullKey[:16], clientHalf)
	copy(fullKey[16:], sessionKey)
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return fullKey[16:], nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims format")
	}
	pointJSON, err := json.Marshal(claims["point"])
	if err != nil {
		return nil, err
	}
	var point types.PointResp
	if err := json.Unmarshal(pointJSON, &point); err != nil {
		return nil, err
	}
	return &types.PxSession{
		Ship:        claims["ship"].(string),
		Ticket:      claims["ticket"].(string),
		Point:       point,
		AuthType:    claims["authType"].(string),
		Wallet:      claims["wallet"].(string),
		EthProvider: ethProvider,
	}, nil
}

func NewHandler() (*Handler, error) {
	tmpl, err := setupTemplates()
	if err != nil {
		return nil, fmt.Errorf("setup templates: %w", err)
	}
	staticFiles, err := fs.Sub(staticFS, "templates/static")
	if err != nil {
		return nil, fmt.Errorf("setup static files: %w", err)
	}
	h := &Handler{
		tmpl: tmpl,
	}
	fileServer := http.FileServer(http.FS(staticFiles))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	return h, nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		h.handleRoot(w, r)
	case "/auth/key-form":
		h.tmpl.ExecuteTemplate(w, "key-form", nil)
	case "/auth/hardware-form":
		h.tmpl.ExecuteTemplate(w, "hardware-form", nil)
	case "/auth/challenge":
		h.handleAuthChallenge(w, r)
	case "/auth":
		h.handleAuth(w, r)
	case "/change-sponsor":
		h.handleChangeSponsor(w, r)
	case "/download/wallet":
		h.handleDownloadWallet(w, r)
	case "/download/keyfile":
		h.handleDownloadKeyfile(w, r)
	case "/breach/form":
		h.handleBreachForm(w, r)
	case "/breach":
		h.handleBreach(w, r)
	case "/transfer/owner/form", "/transfer/management/form", "/transfer/transfer-proxy/form",
		"/transfer/spawn-proxy/form", "/transfer/voting-proxy/form":
		h.handleTransferForm(w, r)
	case "/transfer/owner", "/transfer/management", "/transfer/transfer-proxy",
		"/transfer/spawn-proxy", "/transfer/voting-proxy":
		h.handleTransferSubmit(w, r)
	case "/transfer/cancel":
		h.handleTransferCancel(w, r)
	case "/escape/form":
		h.handleEscapeForm(w, r)
	case "/escape/cancel":
		h.handleCancelEscape(w, r)
	case "/escape":
		h.handleEscape(w, r)
	case "/adopt":
		h.handleAdopt(w, r)
	case "/wallet/connect":
		h.handleWalletConnect(w, r)
	case "/wallet/sign":
		h.handleWalletSign(w, r)
	}
}

func (h *Handler) handleRoot(w http.ResponseWriter, r *http.Request) {
	h.tmpl.ExecuteTemplate(w, "login", nil)
}

func (h *Handler) handleAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		h.tmpl.ExecuteTemplate(w, "login-error", "Method not allowed")
		return
	}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.tmpl.ExecuteTemplate(w, "login-error", fmt.Sprintf("Invalid form data: %v", err))
		return
	}
	ship := r.FormValue("ship")
	authType := r.FormValue("auth_type")
	point, err := libprg.Point(ship)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		h.tmpl.ExecuteTemplate(w, "login-error", fmt.Sprintf("Error retrieving point info: %v", err))
		return
	}

	var ticket string
	var address string
	var validAuthType string
	var wallet string

	if authType == "hardware" {
		address = r.FormValue("address")
		signature := r.FormValue("signature")
		challenge := r.FormValue("challenge")
		if !verifySignature(address, signature, challenge) {
			w.WriteHeader(http.StatusUnauthorized)
			h.tmpl.ExecuteTemplate(w, "login-error", "Invalid signature")
			return
		}
		validAuthType = "hardware"
		wallet = address
	} else {
		ticket = r.FormValue("ticket")
		passphrase := r.FormValue("passphrase")
		_, address, _, _, authType, err = libprg.ValidateKey(ship, ticket, passphrase, "", false)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			h.tmpl.ExecuteTemplate(w, "login-error", err)
			return
		}
		validAuthType = authType
		wallet = address
	}
	var tokenString string
	if tokenString, err = CreateSession(w, ship, ticket, point, validAuthType, address); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.tmpl.ExecuteTemplate(w, "login-error", fmt.Sprintf("Session creation failed: %v", err))
		return
	}
	pxSession := types.PxSession{
		Ship:        ship,
		Point:       point,
		Ticket:      ticket,
		AuthType:    validAuthType,
		Wallet:      wallet,
		EthProvider: ethProvider,
	}
	data := struct {
		types.PxSession
		SessionToken string
	}{
		PxSession:    pxSession,
		SessionToken: w.Header().Get("X-Session-Token"),
	}
	w.Header().Set("X-Session-Token", tokenString)
	w.WriteHeader(http.StatusOK)
	if err := h.tmpl.ExecuteTemplate(w, "dashboard-content", data); err != nil {
		fmt.Printf("Template error: %v\n", err)
	}
	if err := h.tmpl.ExecuteTemplate(w, "dashboard-content", pxSession); err != nil {
		fmt.Printf("Template error: %v\n", err)
	}
}

func (h *Handler) handleAuthChallenge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	ship := r.FormValue("ship")
	address := r.FormValue("address")
	challenge := fmt.Sprintf("Sign this message to authenticate %s: %x", ship, securecookie.GenerateRandomKey(16))
	data := map[string]string{
		"challenge": challenge,
		"address":   address,
		"ship":      ship,
	}
	h.tmpl.ExecuteTemplate(w, "sign-challenge", data)
}

func (h *Handler) handleChangeSponsor(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
        <form hx-post="/update-sponsor" hx-target="closest .info-row">
            <input type="text" name="sponsor" placeholder="~new-sponsor">
            <button type="submit">Update</button>
        </form>
    `))
}

func (h *Handler) handleDownloadWallet(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}
	var session types.PxSession
	if err := securecookie.DecodeMulti("session", cookie.Value, &session,
		securecookie.New(sessionKey, nil)); err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	life, err := strconv.Atoi(session.Point.Point.Network.Keys.Life)
	if err != nil {
		http.Error(w, "Invalid life value", http.StatusUnauthorized)
		return
	}
	wallet, err := libprg.Wallet(session.Ship, session.Ticket, "", life)
	if err != nil {
		fmt.Printf("Wallet generation error: %v\n", err)
		http.Error(w, "Failed to generate wallet", http.StatusInternalServerError)
		return
	}
	walletJSON, err := json.Marshal(wallet)
	if err != nil {
		http.Error(w, "Failed to marshal wallet", http.StatusInternalServerError)
		return
	}
	var walletMap map[string]interface{}
	if err := json.Unmarshal(walletJSON, &walletMap); err != nil {
		http.Error(w, "Failed to process wallet data", http.StatusInternalServerError)
		return
	}
	nfoContent := formatToNFO(walletMap, fmt.Sprintf("%s WALLET", strings.ToUpper(session.Ship)))
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s-wallet.nfo", strings.TrimPrefix(session.Ship, "~")))
	w.Write([]byte(nfoContent))
}

func (h *Handler) handleDownloadKeyfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}
	var session types.PxSession
	if err := securecookie.DecodeMulti("session", cookie.Value, &session,
		securecookie.New(sessionKey, nil)); err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	keyfile, err := libprg.Keyfile(session.Ship, session.Ticket, "", 0)
	if err != nil {
		http.Error(w, "Failed to generate keyfile", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s-%s.key", strings.TrimPrefix(session.Ship, "~"), session.Point.Point.Network.Keys.Life))
	w.Write([]byte(keyfile))
}

func (h *Handler) handleBreachForm(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}
	var session types.PxSession
	if err := securecookie.DecodeMulti("session", cookie.Value, &session,
		securecookie.New(sessionKey, nil)); err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	data := struct {
		Point    types.PointResp
		AuthType string
	}{
		Point:    session.Point,
		AuthType: session.AuthType,
	}
	if err := h.tmpl.ExecuteTemplate(w, "breach-form", data); err != nil {
		http.Error(w, fmt.Sprintf("Template error: %v", err), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) handleBreach(w http.ResponseWriter, r *http.Request) {
	var seed string
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}
	var session types.PxSession
	if err := securecookie.DecodeMulti("session", cookie.Value, &session,
		securecookie.New(sessionKey, nil)); err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	if session.AuthType != "ticket" {
		seed = r.FormValue("seed")
		if seed == "" {
			http.Error(w, "Must provide a seed if not authenticating with master ticket", http.StatusBadRequest)
			return
		}
		isHex, err := regexp.MatchString("^[0-9a-fA-F]{64}$", seed)
		if err != nil || !isHex {
			http.Error(w, "Invalid hex seed", http.StatusBadRequest)
			return
		}
	}
	receipt, err := libprg.Breach(session.Ship, session.Ticket, session.Passphrase, seed)
	if err != nil {
		http.Error(w, "Failed to breach network keys", http.StatusInternalServerError)
		return
	}
	nfoContent := formatToNFO(receipt, "NETWORK BREACH RECEIPT")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s-breach-receipt.nfo", session.Ship))
	w.Write([]byte(nfoContent))
}

func (h *Handler) handleTransferForm(w http.ResponseWriter, r *http.Request) {
	transferType := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/transfer/"), "/form")
	data := types.PxTransferFormData{
		Type: transferType,
	}

	if err := h.tmpl.ExecuteTemplate(w, "transfer-form", data); err != nil {
		http.Error(w, fmt.Sprintf("Template error: %v", err), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) handleTransferSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}
	var session types.PxSession
	if err := securecookie.DecodeMulti("session", cookie.Value, &session,
		securecookie.New(sessionKey, nil)); err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	transferType := strings.TrimPrefix(r.URL.Path, "/transfer/")
	address := r.FormValue("address")
	var receipt interface{}
	switch transferType {
	case "owner":
		receipt, err = libprg.TransferOwnership(session.Ship, session.Ticket, session.Passphrase, address, true)
	case "management":
		receipt, err = libprg.SetManagementProxy(session.Ship, session.Ticket, session.Passphrase, address)
	case "transfer-proxy":
		receipt, err = libprg.SetTransferProxy(session.Ship, session.Ticket, session.Passphrase, address)
	case "spawn-proxy":
		receipt, err = libprg.SetSpawnProxy(session.Ship, session.Ticket, session.Passphrase, address)
	case "voting-proxy":
		receipt, err = libprg.SetVotingProxy(session.Ship, session.Ticket, session.Passphrase, address)
	default:
		http.Error(w, "Invalid transfer type", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("Transfer failed: %v", err), http.StatusInternalServerError)
		return
	}
	nfoContent := formatToNFO(receipt, fmt.Sprintf("%s TRANSFER RECEIPT", strings.ToUpper(transferType)))
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s-%s-transfer-receipt.nfo",
		strings.TrimPrefix(session.Ship, "~"),
		strings.ReplaceAll(transferType, "-", "_")))
	w.Write([]byte(nfoContent))
}

func (h *Handler) handleTransferCancel(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func (h *Handler) handleEscapeForm(w http.ResponseWriter, r *http.Request) {
	if err := h.tmpl.ExecuteTemplate(w, "escape-form", nil); err != nil {
		http.Error(w, fmt.Sprintf("Template error: %v", err), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) handleEscape(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	sponsor := r.FormValue("sponsor")
	if sponsor == "" {
		http.Error(w, "No sponsor specified", http.StatusBadRequest)
		return
	}
	token := r.Header.Get("X-Session-Token")
	session, err := ValidateSession(r, token)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println(session)
	var receipt interface{}
	if session.AuthType == "hardware" {
		if r.FormValue("signed_tx") == "" {
			sponsor := r.FormValue("sponsor")
			if sponsor == "" {
				http.Error(w, "No sponsor specified", http.StatusBadRequest)
				return
			}
			rawTx, err := libprg.EscapeRawTx(session.Wallet, session.Ship, sponsor)
			if err != nil {
				http.Error(w, fmt.Sprintf("Failed to encode transaction: %v", err), http.StatusInternalServerError)
				return
			}
			data := struct {
				Point     types.PointResp
				TxData    string
				Address   string
				Operation string
			}{
				Point:     session.Point,
				TxData:    hexutil.Encode(rawTx),
				Address:   session.Wallet,
				Operation: "escape",
			}
			h.tmpl.ExecuteTemplate(w, "transaction-sign", data)
			return
		}
		signature := r.FormValue("signed_tx")
		if session.Point.Point.Dominion == "l1" {
			receipt, err = libprg.SubmitSignedL1Tx(signature)
			if err != nil {
				http.Error(w, fmt.Sprintf("Failed to broadcast transaction: %v", err), http.StatusInternalServerError)
				return
			}
		} else {
			receipt = "{\"receipt\":\"l2 signed txo\"}"
		}
		nfoContent := formatToNFO(receipt, "ESCAPE REQUEST RECEIPT")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Disposition",
			fmt.Sprintf("attachment; filename=%s-escape-receipt.nfo",
				strings.TrimPrefix(session.Ship, "~")))
		if r.Header.Get("HX-Request") == "true" {
			w.Header().Set("HX-Trigger", "download")
		}
		w.Write([]byte(nfoContent))
		return
	} else {
		receipt, err = libprg.Escape(session.Ship, session.Ticket, session.Passphrase, sponsor)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to initiate escape: %v", err), http.StatusInternalServerError)
			return
		}
	}
	nfoContent := formatToNFO(receipt, "ESCAPE REQUEST RECEIPT")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition",
		fmt.Sprintf("attachment; filename=%s-escape-receipt.nfo",
			strings.TrimPrefix(session.Ship, "~")))
	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Trigger", "download")
	}
	w.Write([]byte(nfoContent))
}

func (h *Handler) handleCancelEscape(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}
	var session types.PxSession
	if err := securecookie.DecodeMulti("session", cookie.Value, &session,
		securecookie.New(sessionKey, nil)); err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	receipt, err := libprg.CancelEscape(session.Ship, session.Ticket, session.Passphrase, "")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to cancel escape: %v", err), http.StatusInternalServerError)
		return
	}
	nfoContent := formatToNFO(receipt, "ESCAPE CANCELLATION RECEIPT")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition",
		fmt.Sprintf("attachment; filename=%s-escape-cancel-receipt.nfo",
			strings.TrimPrefix(session.Ship, "~")))
	w.Write([]byte(nfoContent))
}

func (h *Handler) handleAdopt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	adoptee := r.FormValue("adoptee")
	if adoptee == "" {
		http.Error(w, "No adoptee specified", http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}
	var session types.PxSession
	if err := securecookie.DecodeMulti("session", cookie.Value, &session,
		securecookie.New(sessionKey, nil)); err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	receipt, err := libprg.Adopt(session.Ship, session.Ticket, session.Passphrase, adoptee)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to adopt point: %v", err), http.StatusInternalServerError)
		return
	}
	nfoContent := formatToNFO(receipt, "ADOPTION RECEIPT")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition",
		fmt.Sprintf("attachment; filename=%s-adopt-%s-receipt.nfo",
			strings.TrimPrefix(session.Ship, "~"),
			strings.TrimPrefix(adoptee, "~")))
	w.Write([]byte(nfoContent))
}

func (h *Handler) handleWalletConnect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}
	var session types.PxSession
	if err := securecookie.DecodeMulti("session", cookie.Value, &session,
		securecookie.New(sessionKey, nil)); err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	session.AuthType = r.FormValue("type")
	session.Wallet = r.FormValue("address")
	CreateSession(w, session.Ship, session.Ticket, session.Point, session.AuthType, r.FormValue("address"))
	if err := h.tmpl.ExecuteTemplate(w, "wallet-connected", session); err != nil {
		http.Error(w, fmt.Sprintf("Template error: %v", err), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) handleWalletSign(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}
	var session types.PxSession
	if err := securecookie.DecodeMulti("session", cookie.Value, &session,
		securecookie.New(sessionKey, nil)); err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	if session.AuthType == "hardware" {
		if err := h.tmpl.ExecuteTemplate(w, "wallet-sign-prompt", map[string]string{
			"message": r.FormValue("message"),
			"address": session.Wallet,
		}); err != nil {
			http.Error(w, fmt.Sprintf("Template error: %v", err), http.StatusInternalServerError)
		}
		return
	}
}

func verifySignature(address, signature, challenge string) bool {
	if address[:2] == "0x" {
		address = address[2:]
	}
	if signature[:2] == "0x" {
		signature = signature[2:]
	}
	sigBytes, err := hexutil.Decode("0x" + signature)
	if err != nil {
		return false
	}
	msg := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(challenge)) + challenge)
	hash := crypto.Keccak256Hash(msg)
	if len(sigBytes) != 65 {
		return false
	}
	if sigBytes[64] >= 27 {
		sigBytes[64] -= 27
	}
	pubKey, err := crypto.SigToPub(hash.Bytes(), sigBytes)
	if err != nil {
		return false
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return strings.EqualFold(recoveredAddr.Hex(), "0x"+address)
}
