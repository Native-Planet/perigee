//go:build js

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"syscall/js"

	"github.com/Native-Planet/perigee/libprg"
	"github.com/Native-Planet/perigee/pontifex"
	"github.com/Native-Planet/perigee/roller"
	"github.com/Native-Planet/perigee/types"
)

type session struct {
	Ship       string
	Ticket     string
	Passphrase string
	Point      types.PointResp
	AuthType   string
}

type sessionPayload struct {
	Ship     string          `json:"ship"`
	AuthType string          `json:"authType"`
	Point    types.PointResp `json:"point"`
}

var (
	current   *session
	callbacks []js.Func
)

func main() {
	registerCallbacks()
	select {}
}

func registerCallbacks() {
	perigee := js.Global().Get("Object").New()
	perigee.Set("login", makePromise(login))
	perigee.Set("session", makePromise(getSession))
	perigee.Set("wallet", makePromise(wallet))
	perigee.Set("keyfile", makePromise(keyfile))
	perigee.Set("breach", makePromise(breach))
	perigee.Set("escape", makePromise(escape))
	perigee.Set("cancelEscape", makePromise(cancelEscape))
	perigee.Set("adopt", makePromise(adopt))
	perigee.Set("transfer", makePromise(transfer))
	perigee.Set("setEthProvider", makePromise(setEthProvider))
	perigee.Set("setRollerURL", makePromise(setRollerURL))
	perigee.Set("sigil", makePromise(sigil))
	js.Global().Set("perigee", perigee)
}

func makePromise(fn func([]js.Value) (interface{}, error)) js.Func {
	promise := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Capture caller args so we don't accidentally forward Promise resolve/reject callbacks.
		callerArgs := append([]js.Value{}, args...)
		pCtor := js.Global().Get("Promise")
		handler := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]
			go func() {
				res, err := fn(callerArgs)
				if err != nil {
					reject.Invoke(js.ValueOf(err.Error()))
					return
				}
				resolve.Invoke(toJSValue(res))
			}()
			return nil
		})
		defer handler.Release()
		return pCtor.New(handler)
	})
	callbacks = append(callbacks, promise)
	return promise
}

func toJSValue(value interface{}) js.Value {
	switch v := value.(type) {
	case nil:
		return js.Undefined()
	case js.Value:
		return v
	case string:
		return js.ValueOf(v)
	case bool:
		return js.ValueOf(v)
	default:
		raw, err := json.Marshal(v)
		if err != nil {
			return js.ValueOf(fmt.Sprintf("marshal error: %v", err))
		}
		return js.ValueOf(string(raw))
	}
}

func login(args []js.Value) (interface{}, error) {
	if len(args) < 2 {
		return nil, errors.New("login requires a ship and ticket/private key")
	}
	ship := strings.TrimSpace(args[0].String())
	key := strings.TrimSpace(args[1].String())
	passphrase := ""
	if len(args) > 2 {
		passphrase = strings.TrimSpace(args[2].String())
	}
	point, err := libprg.Point(ship)
	if err != nil {
		return nil, fmt.Errorf("point lookup: %w", err)
	}
	_, _, _, _, authType, err := libprg.ValidateKey(ship, key, passphrase, "", false)
	if err != nil {
		return nil, fmt.Errorf("key validation: %w", err)
	}
	current = &session{
		Ship:       ship,
		Ticket:     key,
		Passphrase: passphrase,
		Point:      point,
		AuthType:   authType,
	}
	return sessionPayload{
		Ship:     ship,
		AuthType: authType,
		Point:    point,
	}, nil
}

func getSession(_ []js.Value) (interface{}, error) {
	if current == nil {
		return nil, errors.New("no active session")
	}
	return sessionPayload{
		Ship:     current.Ship,
		AuthType: current.AuthType,
		Point:    current.Point,
	}, nil
}

func wallet(args []js.Value) (interface{}, error) {
	s, err := requireSession()
	if err != nil {
		return nil, err
	}
	life := optionalLife(s, args)
	w, err := libprg.Wallet(s.Ship, s.Ticket, s.Passphrase, life)
	if err != nil {
		return nil, fmt.Errorf("wallet generation: %w", err)
	}
	return w, nil
}

func keyfile(args []js.Value) (interface{}, error) {
	s, err := requireSession()
	if err != nil {
		return nil, err
	}
	life := optionalLife(s, args)
	keyfile, err := libprg.Keyfile(s.Ship, s.Ticket, s.Passphrase, life)
	if err != nil {
		return nil, fmt.Errorf("keyfile generation: %w", err)
	}
	return keyfile, nil
}

func breach(args []js.Value) (interface{}, error) {
	s, err := requireSession()
	if err != nil {
		return nil, err
	}
	seed := ""
	if len(args) > 0 {
		seed = strings.TrimSpace(args[0].String())
	}
	if s.AuthType != "ticket" && seed == "" {
		return nil, errors.New("seed is required when authenticating with an Ethereum private key")
	}
	receipt, err := libprg.Breach(s.Ship, s.Ticket, s.Passphrase, seed)
	if err != nil {
		return nil, fmt.Errorf("breach failed: %w", err)
	}
	return receipt, nil
}

func escape(args []js.Value) (interface{}, error) {
	s, err := requireSession()
	if err != nil {
		return nil, err
	}
	if len(args) == 0 {
		return nil, errors.New("escape requires a sponsor (~sponsor)")
	}
	sponsor := strings.TrimSpace(args[0].String())
	receipt, err := libprg.Escape(s.Ship, s.Ticket, s.Passphrase, sponsor)
	if err != nil {
		return nil, fmt.Errorf("escape failed: %w", err)
	}
	return receipt, nil
}

func cancelEscape(args []js.Value) (interface{}, error) {
	s, err := requireSession()
	if err != nil {
		return nil, err
	}
	sponsor := ""
	if len(args) > 0 {
		sponsor = strings.TrimSpace(args[0].String())
	}
	receipt, err := libprg.CancelEscape(s.Ship, s.Ticket, s.Passphrase, sponsor)
	if err != nil {
		return nil, fmt.Errorf("cancel escape failed: %w", err)
	}
	return receipt, nil
}

func adopt(args []js.Value) (interface{}, error) {
	s, err := requireSession()
	if err != nil {
		return nil, err
	}
	if len(args) == 0 {
		return nil, errors.New("adopt requires an adoptee (~child)")
	}
	adoptee := strings.TrimSpace(args[0].String())
	receipt, err := libprg.Adopt(s.Ship, s.Ticket, s.Passphrase, adoptee)
	if err != nil {
		return nil, fmt.Errorf("adopt failed: %w", err)
	}
	return receipt, nil
}

func transfer(args []js.Value) (interface{}, error) {
	s, err := requireSession()
	if err != nil {
		return nil, err
	}
	if len(args) < 2 {
		return nil, errors.New("transfer requires a type and address")
	}
	transferType := strings.TrimSpace(args[0].String())
	address := strings.TrimSpace(args[1].String())
	reset := true
	if len(args) > 2 {
		reset = args[2].Bool()
	}

	switch transferType {
	case "owner":
		return libprg.TransferOwnership(s.Ship, s.Ticket, s.Passphrase, address, reset)
	case "management":
		return libprg.SetManagementProxy(s.Ship, s.Ticket, s.Passphrase, address)
	case "transfer-proxy":
		return libprg.SetTransferProxy(s.Ship, s.Ticket, s.Passphrase, address)
	case "spawn-proxy":
		return libprg.SetSpawnProxy(s.Ship, s.Ticket, s.Passphrase, address)
	case "voting-proxy":
		return libprg.SetVotingProxy(s.Ship, s.Ticket, s.Passphrase, address)
	default:
		return nil, fmt.Errorf("unsupported transfer type: %s", transferType)
	}
}

func setEthProvider(args []js.Value) (interface{}, error) {
	if len(args) == 0 {
		libprg.EthProvider = ""
		return "", nil
	}
	eth := strings.TrimSpace(args[0].String())
	libprg.EthProvider = eth
	return eth, nil
}

func setRollerURL(args []js.Value) (interface{}, error) {
	url := "https://roller.urbit.org"
	if len(args) > 0 {
		url = strings.TrimSpace(args[0].String())
	}
	if url == "" {
		return nil, errors.New("roller URL required")
	}
	roller.RollerURL = url
	roller.Client = roller.New(roller.Config{
		Endpoint:   url,
		HTTPClient: http.DefaultClient,
	})
	return url, nil
}

func sigil(args []js.Value) (interface{}, error) {
	if len(args) == 0 {
		return nil, errors.New("patp required")
	}
	patp := strings.TrimSpace(args[0].String())
	size := 128
	if len(args) > 1 {
		if n, err := strconv.Atoi(strings.TrimSpace(args[1].String())); err == nil && n > 0 {
			size = n
		}
	}
	svg, err := pontifex.GenerateSigil(types.PxSvgConfig{
		Point: patp,
		Size:  size,
	})
	if err != nil {
		return nil, fmt.Errorf("sigil generation: %w", err)
	}
	return svg, nil
}

func requireSession() (*session, error) {
	if current == nil {
		return nil, errors.New("no active session; call login first")
	}
	return current, nil
}

func optionalLife(s *session, args []js.Value) int {
	if len(args) > 0 {
		if lifeVal, err := strconv.Atoi(strings.TrimSpace(args[0].String())); err == nil && lifeVal > 0 {
			return lifeVal
		}
	}
	life, _ := strconv.Atoi(s.Point.Point.Network.Keys.Life)
	return life
}
