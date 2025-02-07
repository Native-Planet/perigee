package pontifex

import (
	"crypto/rand"
	"embed"
	"encoding/json"
	"fmt"
	html "html/template"
	"io/fs"
	"net/http"
	"strconv"
	"sync"
	"text/template"

	"github.com/Native-Planet/perigee/libprg"
	"github.com/Native-Planet/perigee/types"

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

func CreateSession(w http.ResponseWriter, ship, ticket string, point types.PointResp) {
	session := &types.Session{Ship: ship, Ticket: ticket, Point: point}
	encoded, err := securecookie.EncodeMulti("session", session,
		securecookie.New(sessionKey, nil))
	if err != nil {
		return
	}
	cookie := &http.Cookie{
		Name:     "session",
		Value:    encoded,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)
}

func GetTemplateFuncs() template.FuncMap {
	funcMapOnce.Do(func() {
		funcMap = template.FuncMap{
			"genSigil": func(patp string, size int) string {
				svg, err := GenerateSigil(types.SvgConfig{
					Point: patp,
					Size:  size,
				})
				if err != nil {
					return ""
				}
				return svg
			},
			"smallSigil": func(patp string) string {
				svg, err := GenerateSigil(types.SvgConfig{
					Point: patp,
					Size:  32,
				})
				if err != nil {
					return ""
				}
				return svg
			},
			"largeSigil": func(patp string) string {
				svg, err := GenerateSigil(types.SvgConfig{
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
			svg, err := GenerateSigil(types.SvgConfig{
				Point: patp,
				Size:  size,
			})
			if err != nil {
				return ""
			}
			return svg
		},
		"smallSigil": func(patp string) string {
			svg, err := GenerateSigil(types.SvgConfig{
				Point: patp,
				Size:  32,
			})
			if err != nil {
				return ""
			}
			return svg
		},
		"largeSigil": func(patp string) string {
			svg, err := GenerateSigil(types.SvgConfig{
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
	case "/auth":
		h.handleAuth(w, r)
	case "/change-sponsor":
		h.handleChangeSponsor(w, r)
	case "/download/wallet":
		h.handleDownloadWallet(w, r)
	case "/download/keyfile":
		h.handleDownloadKeyfile(w, r)
	}
}

func (h *Handler) handleRoot(w http.ResponseWriter, r *http.Request) {
	h.tmpl.ExecuteTemplate(w, "login", nil)
}

func (h *Handler) handleAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Printf("Wrong method: %s\n", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Printf("Form parse error: %v\n", err)
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	ship := r.FormValue("ship")
	ticket := r.FormValue("ticket")

	fmt.Printf("Auth attempt - Ship: '%s', Ticket length: %d\n", ship, len(ticket))

	point, err := libprg.Point(ship)
	if err != nil {
		fmt.Printf("Point retrieval error: %v\n", err)
		http.Error(w, fmt.Sprintf("Error retrieving point info: %v", err), http.StatusUnauthorized)
		return
	}

	CreateSession(w, ship, ticket, point)
	data := types.Session{
		Ship:   ship,
		Point:  point,
		Ticket: ticket,
	}

	w.Header().Set("Content-Type", "text/html")
	if err := h.tmpl.ExecuteTemplate(w, "dashboard-content", data); err != nil {
		fmt.Printf("Template execution error: %v\n", err)
		http.Error(w, fmt.Sprintf("Template error: %v", err), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) handleChangeSponsor(w http.ResponseWriter, r *http.Request) {
	// Return a form for changing sponsor
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
	var session types.Session
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
		http.Error(w, "Failed to generate wallet", http.StatusInternalServerError)
		return
	}
	walletJSON, err := json.MarshalIndent(wallet, "", "  ")
	if err != nil {
		http.Error(w, "Failed to marshal wallet", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s-wallet-%s.json", session.Ship, session.Point.Point.Network.Keys.Life))
	w.Write(walletJSON)
}

func (h *Handler) handleDownloadKeyfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}
	var session types.Session
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
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s-%s.key", cleanPatp(session.Ship), session.Point.Point.Network.Keys.Life))
	w.Write([]byte(keyfile))
}
