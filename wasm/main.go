//go:build js && wasm

package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"syscall/js"
	"time"

	"github.com/Native-Planet/perigee/libprg"
	"github.com/Native-Planet/perigee/roller"
	perigeeTypes "github.com/Native-Planet/perigee/types"
)

const requestTimeout = 45 * time.Second

var (
	currentRollerURL   = "https://roller.urbit.org/v1/roller"
	currentEthProvider string
	currentProxyURL    string
)

type baseRequest struct {
	Roller string `json:"roller,omitempty"`
}

type pointRequest struct {
	baseRequest
	Ship    string `json:"ship,omitempty"`
	Address string `json:"address,omitempty"`
	Hash    string `json:"hash,omitempty"`
}

type materialRequest struct {
	baseRequest
	Ship       string `json:"ship"`
	Ticket     string `json:"ticket"`
	Passphrase string `json:"passphrase"`
	Life       int    `json:"life"`
	Step       int    `json:"step"`
}

type operationRequest struct {
	baseRequest
	Operation      string `json:"operation"`
	Ship           string `json:"ship"`
	CredentialType string `json:"credentialType"`
	Ticket         string `json:"ticket"`
	PrivateKey     string `json:"privateKey"`
	Passphrase     string `json:"passphrase"`
	Seed           string `json:"seed"`
	Sponsor        string `json:"sponsor"`
	Adoptee        string `json:"adoptee"`
	NewOwner       string `json:"newOwner"`
	Proxy          string `json:"proxy"`
	Reset          bool   `json:"reset"`
}

type walletPrepareRequest struct {
	operationRequest
	Address string `json:"address"`
}

type walletSubmitRequest struct {
	walletPrepareRequest
	Signature string `json:"signature"`
}

type pendingSummary struct {
	Operation    string `json:"operation,omitempty"`
	Ship         string `json:"ship,omitempty"`
	Hash         string `json:"hash,omitempty"`
	Signature    string `json:"signature,omitempty"`
	Status       string `json:"status,omitempty"`
	SubmittedAt  int64  `json:"submittedAt,omitempty"`
	NextPollAt   int64  `json:"nextPollAt,omitempty"`
	PollInterval int    `json:"pollInterval,omitempty"`
}

type pointResponse struct {
	OK        bool                     `json:"ok"`
	Ship      string                   `json:"ship,omitempty"`
	Point     *perigeeTypes.Point      `json:"point,omitempty"`
	Pending   []perigeeTypes.PendingTx `json:"pending,omitempty"`
	Batch     *perigeeTypes.BatchInfo  `json:"batch,omitempty"`
	Status    string                   `json:"status,omitempty"`
	PendingTx any                      `json:"pendingTx,omitempty"`
}

type operationResponse struct {
	OK              bool            `json:"ok"`
	Ship            string          `json:"ship"`
	Operation       string          `json:"operation"`
	Transaction     any             `json:"transaction,omitempty"`
	Pending         *pendingSummary `json:"pending,omitempty"`
	ExportSuggested bool            `json:"exportSuggested,omitempty"`
	Message         string          `json:"message,omitempty"`
}

type prepareResponse struct {
	OK             bool           `json:"ok"`
	Ship           string         `json:"ship"`
	Operation      string         `json:"operation"`
	Address        string         `json:"address"`
	Seed           string         `json:"seed,omitempty"`
	SigningPayload string         `json:"signingPayload"`
	SignMethod     string         `json:"signMethod"`
	Method         string         `json:"method"`
	From           map[string]any `json:"from"`
	Data           any            `json:"data"`
	Nonce          int            `json:"nonce"`
}

type rpcRequest struct {
	Version string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
	ID      string `json:"id"`
}

type rpcResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func main() {
	api := map[string]any{
		"setRollerURL":           promiseFunc(setRollerURL),
		"setEthProvider":         promiseFunc(setEthProvider),
		"setRollerProxy":         promiseFunc(setRollerProxy),
		"point":                  promiseFunc(getPoint),
		"keyfile":                promiseFunc(generateKeyfile),
		"code":                   promiseFunc(generateCode),
		"operation":              promiseFunc(submitOperation),
		"prepareWalletOperation": promiseFunc(prepareWallet),
		"submitWalletOperation":  promiseFunc(submitWallet),
	}
	js.Global().Set("perigee", js.ValueOf(api))
	select {}
}

func promiseFunc(fn func([]js.Value) (any, error)) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		promise := js.Global().Get("Promise")
		executor := js.FuncOf(func(_ js.Value, callbacks []js.Value) any {
			resolve := callbacks[0]
			reject := callbacks[1]
			go func() {
				result, err := fn(args)
				if err != nil {
					reject.Invoke(js.Global().Get("Error").New(err.Error()))
					return
				}
				resolve.Invoke(marshalResult(result))
			}()
			return nil
		})
		return promise.New(executor)
	})
}

func marshalResult(value any) js.Value {
	if value == nil {
		return js.Null()
	}
	if str, ok := value.(string); ok {
		return js.ValueOf(str)
	}
	raw, err := json.Marshal(value)
	if err != nil {
		return js.ValueOf(fmt.Sprintf(`{"ok":false,"error":%q}`, err.Error()))
	}
	return js.ValueOf(string(raw))
}

func decodeArg(args []js.Value, dst any) error {
	if len(args) < 1 || args[0].IsUndefined() || args[0].IsNull() {
		return fmt.Errorf("payload is required")
	}
	var raw string
	if args[0].Type() == js.TypeString {
		raw = args[0].String()
	} else {
		raw = js.Global().Get("JSON").Call("stringify", args[0]).String()
	}
	if err := json.Unmarshal([]byte(raw), dst); err != nil {
		return fmt.Errorf("invalid payload: %w", err)
	}
	return nil
}

func setRollerURL(args []js.Value) (any, error) {
	if len(args) < 1 || args[0].IsUndefined() || args[0].IsNull() {
		currentRollerURL = "https://roller.urbit.org/v1/roller"
	} else {
		currentRollerURL = normalizeRollerEndpoint(args[0].String())
	}
	configureGlobals(currentRollerURL)
	return map[string]any{"ok": true, "roller": currentRollerURL}, nil
}

func setEthProvider(args []js.Value) (any, error) {
	if len(args) < 1 || args[0].IsUndefined() || args[0].IsNull() {
		currentEthProvider = ""
	} else {
		currentEthProvider = strings.TrimSpace(args[0].String())
	}
	libprg.EthProvider = currentEthProvider
	return map[string]any{"ok": true, "ethProvider": currentEthProvider}, nil
}

func setRollerProxy(args []js.Value) (any, error) {
	if len(args) < 1 || args[0].IsUndefined() || args[0].IsNull() {
		currentProxyURL = ""
	} else {
		currentProxyURL = strings.TrimSpace(args[0].String())
	}
	configureGlobals(currentRollerURL)
	return map[string]any{"ok": true, "proxy": currentProxyURL}, nil
}

func getPoint(args []js.Value) (any, error) {
	var req pointRequest
	if err := decodeArg(args, &req); err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	endpoint, headers := rollerTransport(req.Roller)
	client := roller.New(roller.Config{Endpoint: endpoint, HTTPClient: http.DefaultClient, Headers: headers})
	resp := pointResponse{OK: true}
	if strings.TrimSpace(req.Ship) != "" {
		ship, err := normalizeShip(req.Ship)
		if err != nil {
			return nil, err
		}
		point, err := client.GetPoint(ctx, ship)
		if err != nil {
			return nil, err
		}
		resp.Ship = ship
		resp.Point = point
		if strings.TrimSpace(point.Ownership.Owner.Address) != "" {
			if pending, err := client.GetPendingByAddress(ctx, point.Ownership.Owner.Address); err == nil {
				resp.Pending = pending
			}
		}
	} else if strings.TrimSpace(req.Address) != "" {
		address, err := normalizeAddress(req.Address)
		if err != nil {
			return nil, err
		}
		pending, err := client.GetPendingByAddress(ctx, address)
		if err != nil {
			return nil, err
		}
		resp.Pending = pending
	}
	if batch, err := client.WhenNextBatch(ctx); err == nil {
		resp.Batch = batch
	}
	if strings.TrimSpace(req.Hash) != "" {
		target := normalizeRollerEndpoint(req.Roller)
		if statusRaw, err := rollerRPC(ctx, target, "getTransactionStatus", map[string]any{"hash": strings.TrimSpace(req.Hash)}); err == nil {
			_ = json.Unmarshal(statusRaw, &resp.Status)
		}
		if pendingRaw, err := rollerRPC(ctx, target, "getPendingTx", map[string]any{"hash": strings.TrimSpace(req.Hash)}); err == nil {
			var raw any
			if json.Unmarshal(pendingRaw, &raw) == nil {
				resp.PendingTx = raw
			}
		}
	}
	return resp, nil
}

func generateKeyfile(args []js.Value) (any, error) {
	var req materialRequest
	if err := decodeArg(args, &req); err != nil {
		return nil, err
	}
	configureGlobals(req.Roller)
	ship, err := normalizeShip(req.Ship)
	if err != nil {
		return nil, err
	}
	keyfile, err := libprg.Keyfile(context.Background(), ship, normalizeTicket(req.Ticket), req.Passphrase, req.Life)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"ok":       true,
		"ship":     ship,
		"filename": strings.TrimPrefix(ship, "~") + ".key",
		"keyfile":  keyfile,
	}, nil
}

func generateCode(args []js.Value) (any, error) {
	var req materialRequest
	if err := decodeArg(args, &req); err != nil {
		return nil, err
	}
	configureGlobals(req.Roller)
	ship, err := normalizeShip(req.Ship)
	if err != nil {
		return nil, err
	}
	code, err := libprg.GenerateCode(context.Background(), ship, normalizeTicket(req.Ticket), req.Passphrase, req.Life, req.Step)
	if err != nil {
		return nil, err
	}
	return map[string]any{"ok": true, "ship": ship, "code": code}, nil
}

func submitOperation(args []js.Value) (any, error) {
	var req operationRequest
	if err := decodeArg(args, &req); err != nil {
		return nil, err
	}
	configureGlobals(req.Roller)
	tx, ship, err := runSoftwareOperation(req)
	if err != nil {
		return nil, err
	}
	operation := normalizeOperation(req.Operation)
	resp := operationResponse{
		OK:              true,
		Ship:            ship,
		Operation:       operation,
		Transaction:     tx,
		Pending:         summarizePending(ship, operation, tx),
		ExportSuggested: operation == "breach",
	}
	if resp.ExportSuggested {
		resp.Message = walletOperationMessage(operation)
	}
	return resp, nil
}

func prepareWallet(args []js.Value) (any, error) {
	var req walletPrepareRequest
	if err := decodeArg(args, &req); err != nil {
		return nil, err
	}
	return prepareWalletOperation(context.Background(), req)
}

func submitWallet(args []js.Value) (any, error) {
	var req walletSubmitRequest
	if err := decodeArg(args, &req); err != nil {
		return nil, err
	}
	prepared, err := prepareWalletOperation(context.Background(), req.walletPrepareRequest)
	if err != nil {
		return nil, err
	}
	signature := strings.TrimSpace(req.Signature)
	if signature == "" {
		return nil, fmt.Errorf("signature is required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	target := normalizeRollerEndpoint(req.Roller)
	params := map[string]any{
		"address": prepared.Address,
		"from":    prepared.From,
		"data":    prepared.Data,
		"sig":     signature,
		"force":   false,
	}
	result, err := rollerRPC(ctx, target, prepared.Method, params)
	if err != nil {
		return nil, err
	}
	var txHash string
	if err := json.Unmarshal(result, &txHash); err != nil {
		return nil, fmt.Errorf("unexpected roller response: %w", err)
	}
	tx := perigeeTypes.Transaction{Signature: signature, Hash: txHash, Type: prepared.Method}
	operation := normalizeOperation(req.Operation)
	return operationResponse{
		OK:              true,
		Ship:            prepared.Ship,
		Operation:       operation,
		Transaction:     tx,
		Pending:         summarizePending(prepared.Ship, operation, tx),
		ExportSuggested: operation == "breach",
		Message:         walletOperationMessage(operation),
	}, nil
}

func runSoftwareOperation(req operationRequest) (any, string, error) {
	ship, err := normalizeShip(req.Ship)
	if err != nil {
		return nil, "", err
	}
	credential, err := operationCredential(req)
	if err != nil {
		return nil, "", err
	}
	switch normalizeOperation(req.Operation) {
	case "breach":
		if req.CredentialType == "private-key" {
			tx, err := breachWithPrivateKey(ship, credential, req.Passphrase, req.Seed, req.Roller)
			return tx, ship, err
		}
		tx, err := libprg.Breach(context.Background(), ship, credential, req.Passphrase, req.Seed)
		return tx, ship, err
	case "escape":
		sponsor, err := normalizeShip(req.Sponsor)
		if err != nil {
			return nil, "", fmt.Errorf("invalid sponsor: %w", err)
		}
		tx, err := libprg.Escape(context.Background(), ship, sponsor, credential, req.Passphrase)
		return tx, ship, err
	case "cancel-escape":
		sponsor, err := normalizeShip(req.Sponsor)
		if err != nil {
			return nil, "", fmt.Errorf("invalid sponsor: %w", err)
		}
		tx, err := libprg.CancelEscape(context.Background(), ship, sponsor, credential, req.Passphrase)
		return tx, ship, err
	case "adopt":
		adoptee, err := normalizeShip(req.Adoptee)
		if err != nil {
			return nil, "", fmt.Errorf("invalid adoptee: %w", err)
		}
		tx, err := libprg.Adopt(context.Background(), ship, adoptee, credential, req.Passphrase)
		return tx, ship, err
	case "transfer":
		newOwner, err := normalizeAddress(req.NewOwner)
		if err != nil {
			return nil, "", fmt.Errorf("invalid new owner: %w", err)
		}
		tx, err := libprg.TransferOwnership(context.Background(), ship, credential, req.Passphrase, newOwner, req.Reset)
		return tx, ship, err
	case "set-management-proxy":
		proxy, err := normalizeAddress(req.Proxy)
		if err != nil {
			return nil, "", fmt.Errorf("invalid management proxy: %w", err)
		}
		tx, err := libprg.SetManagementProxy(context.Background(), ship, credential, req.Passphrase, proxy)
		return tx, ship, err
	case "set-spawn-proxy":
		proxy, err := normalizeAddress(req.Proxy)
		if err != nil {
			return nil, "", fmt.Errorf("invalid spawn proxy: %w", err)
		}
		tx, err := libprg.SetSpawnProxy(context.Background(), ship, credential, req.Passphrase, proxy)
		return tx, ship, err
	case "set-transfer-proxy":
		proxy, err := normalizeAddress(req.Proxy)
		if err != nil {
			return nil, "", fmt.Errorf("invalid transfer proxy: %w", err)
		}
		tx, err := libprg.SetTransferProxy(context.Background(), ship, credential, req.Passphrase, proxy)
		return tx, ship, err
	default:
		return nil, "", fmt.Errorf("unsupported operation: %s", req.Operation)
	}
}

func prepareWalletOperation(parent context.Context, req walletPrepareRequest) (prepareResponse, error) {
	ship, err := normalizeShip(req.Ship)
	if err != nil {
		return prepareResponse{}, err
	}
	address, err := normalizeAddress(req.Address)
	if err != nil {
		return prepareResponse{}, err
	}
	operation := normalizeOperation(req.Operation)
	ctx, cancel := context.WithTimeout(parent, requestTimeout)
	defer cancel()
	target := normalizeRollerEndpoint(req.Roller)
	endpoint, headers := rollerTransport(req.Roller)
	rClient := roller.New(roller.Config{Endpoint: endpoint, HTTPClient: http.DefaultClient, Headers: headers})
	method, data, proxyKind, seed, err := walletOperationData(operation, req.operationRequest)
	if err != nil {
		return prepareResponse{}, err
	}
	point, err := rClient.GetPoint(ctx, ship)
	if err != nil {
		return prepareResponse{}, fmt.Errorf("getting point: %w", err)
	}
	proxy, err := proxyTypeForAddress(point, address, proxyKind)
	if err != nil {
		return prepareResponse{}, err
	}
	from := map[string]any{"ship": ship, "proxy": proxy}
	client := perigeeTypes.Client{Endpoint: endpoint, HttpClient: http.DefaultClient, Headers: headers}
	nonce, err := client.GetNonce(ctx, map[string]any{"from": from})
	if err != nil {
		return prepareResponse{}, fmt.Errorf("getting nonce: %w", err)
	}
	signingPayloadRaw, err := rollerRPC(ctx, target, "prepareForSigning", map[string]any{
		"nonce": nonce,
		"from":  from,
		"tx":    method,
		"data":  data,
	})
	if err != nil {
		return prepareResponse{}, fmt.Errorf("preparing transaction for signing: %w", err)
	}
	var signingPayload string
	if err := json.Unmarshal(signingPayloadRaw, &signingPayload); err != nil {
		return prepareResponse{}, fmt.Errorf("unexpected signing payload: %w", err)
	}
	return prepareResponse{
		OK:             true,
		Ship:           ship,
		Operation:      operation,
		Address:        address,
		Seed:           seed,
		SigningPayload: signingPayload,
		SignMethod:     "personal_sign",
		Method:         method,
		From:           from,
		Data:           data,
		Nonce:          nonce,
	}, nil
}

func walletOperationData(operation string, req operationRequest) (string, any, string, string, error) {
	switch operation {
	case "breach":
		seed := strings.TrimPrefix(strings.TrimSpace(req.Seed), "0x")
		if seed == "" {
			generated, err := defaultNetworkSeed()
			if err != nil {
				return "", nil, "", "", err
			}
			seed = generated
		}
		keys, err := libprg.GenerateNetworkKeysFromSeed(seed)
		if err != nil {
			return "", nil, "", "", err
		}
		return "configureKeys", map[string]any{
			"encrypt":     "0x" + keys.Crypt.Public,
			"auth":        "0x" + keys.Auth.Public,
			"cryptoSuite": "1",
			"breach":      true,
		}, "management", seed, nil
	case "escape":
		sponsor, err := normalizeShip(req.Sponsor)
		if err != nil {
			return "", nil, "", "", fmt.Errorf("invalid sponsor: %w", err)
		}
		return "escape", map[string]any{"ship": sponsor}, "management", "", nil
	case "cancel-escape":
		sponsor, err := normalizeShip(req.Sponsor)
		if err != nil {
			return "", nil, "", "", fmt.Errorf("invalid sponsor: %w", err)
		}
		return "cancel-escape", map[string]any{"ship": sponsor}, "management", "", nil
	case "adopt":
		adoptee, err := normalizeShip(req.Adoptee)
		if err != nil {
			return "", nil, "", "", fmt.Errorf("invalid adoptee: %w", err)
		}
		return "adopt", map[string]any{"ship": adoptee}, "management", "", nil
	case "transfer":
		newOwner, err := normalizeAddress(req.NewOwner)
		if err != nil {
			return "", nil, "", "", fmt.Errorf("invalid new owner: %w", err)
		}
		return "transferPoint", map[string]any{"address": newOwner, "reset": req.Reset}, "transfer", "", nil
	case "set-management-proxy":
		proxy, err := normalizeAddress(req.Proxy)
		if err != nil {
			return "", nil, "", "", fmt.Errorf("invalid management proxy: %w", err)
		}
		return "setManagementProxy", map[string]any{"address": proxy}, "management", "", nil
	case "set-spawn-proxy":
		proxy, err := normalizeAddress(req.Proxy)
		if err != nil {
			return "", nil, "", "", fmt.Errorf("invalid spawn proxy: %w", err)
		}
		return "setSpawnProxy", map[string]any{"address": proxy}, "management", "", nil
	case "set-transfer-proxy":
		proxy, err := normalizeAddress(req.Proxy)
		if err != nil {
			return "", nil, "", "", fmt.Errorf("invalid transfer proxy: %w", err)
		}
		return "setTransferProxy", map[string]any{"address": proxy}, "management", "", nil
	default:
		return "", nil, "", "", fmt.Errorf("unsupported wallet operation: %s", operation)
	}
}

func breachWithPrivateKey(ship, privateKey, passphrase, seed, rollerEndpoint string) (*perigeeTypes.Transaction, error) {
	seed = strings.TrimPrefix(strings.TrimSpace(seed), "0x")
	if seed == "" {
		generated, err := defaultNetworkSeed()
		if err != nil {
			return nil, err
		}
		seed = generated
	}
	privKey, derivedPubkey, pointInfo, networkKeys, _, err := libprg.ValidateKey(context.Background(), ship, privateKey, passphrase, seed, true)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", libprg.ErrKeyMaterial, err)
	}
	if pointInfo.Dominion != "l2" {
		return nil, fmt.Errorf("private-key breach through Roller is only available for L2 ships")
	}
	endpoint, headers := rollerTransport(rollerEndpoint)
	return roller.New(roller.Config{Endpoint: endpoint, HTTPClient: http.DefaultClient, Headers: headers}).ConfigureKeys(
		context.Background(),
		ship,
		"0x"+networkKeys.Crypt.Public,
		"0x"+networkKeys.Auth.Public,
		true,
		derivedPubkey,
		privKey,
	)
}

func defaultNetworkSeed() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", fmt.Errorf("generate network seed: %w", err)
	}
	return hex.EncodeToString(buf), nil
}

func operationCredential(req operationRequest) (string, error) {
	switch req.CredentialType {
	case "ticket", "master-ticket", "":
		ticket := normalizeTicket(req.Ticket)
		if strings.TrimSpace(ticket) == "" {
			return "", fmt.Errorf("master ticket is required")
		}
		return ticket, nil
	case "private-key":
		privateKey := strings.TrimPrefix(strings.TrimSpace(req.PrivateKey), "0x")
		if privateKey == "" {
			return "", fmt.Errorf("ethereum private key is required")
		}
		return privateKey, nil
	default:
		return "", fmt.Errorf("unsupported credential type: %s", req.CredentialType)
	}
}

func proxyTypeForAddress(point *perigeeTypes.Point, address, kind string) (string, error) {
	if strings.EqualFold(point.Ownership.Owner.Address, address) {
		return "own", nil
	}
	switch kind {
	case "management":
		if strings.EqualFold(point.Ownership.ManagementProxy.Address, address) {
			return "manage", nil
		}
		return "", fmt.Errorf("connected wallet is neither owner nor management proxy for this ship")
	case "transfer":
		if strings.EqualFold(point.Ownership.TransferProxy.Address, address) {
			return "transfer", nil
		}
		return "", fmt.Errorf("connected wallet is neither owner nor transfer proxy for this ship")
	case "spawn":
		if strings.EqualFold(point.Ownership.SpawnProxy.Address, address) {
			return "spawn", nil
		}
		return "", fmt.Errorf("connected wallet is neither owner nor spawn proxy for this ship")
	default:
		return "", fmt.Errorf("unknown proxy kind: %s", kind)
	}
}

func configureGlobals(rawRoller string) {
	target := normalizeRollerEndpoint(rawRoller)
	endpoint, headers := rollerTransport(target)
	roller.RollerURL = endpoint
	roller.Client = roller.New(roller.Config{Endpoint: endpoint, HTTPClient: http.DefaultClient, Headers: headers})
	libprg.EthProvider = currentEthProvider
}

func rollerTransport(rawRoller string) (string, http.Header) {
	target := normalizeRollerEndpoint(rawRoller)
	if strings.TrimSpace(currentProxyURL) == "" {
		return target, nil
	}
	headers := http.Header{}
	headers.Set("X-Groundseg-Roller-URL", target)
	return currentProxyURL, headers
}

func normalizeRollerEndpoint(raw string) string {
	endpoint := strings.TrimSpace(raw)
	if endpoint == "" {
		endpoint = currentRollerURL
	}
	if endpoint == "" {
		endpoint = "https://roller.urbit.org/v1/roller"
	}
	if !strings.Contains(endpoint, "://") {
		endpoint = "https://" + endpoint
	}
	parsed, err := url.Parse(endpoint)
	if err != nil || parsed.Host == "" {
		return endpoint
	}
	if parsed.Path == "" || parsed.Path == "/" {
		parsed.Path = "/v1/roller"
	}
	return parsed.String()
}

func normalizeShip(ship string) (string, error) {
	patp, _, err := perigeeTypes.ValidateAndNormalizePatp(strings.TrimSpace(ship))
	if err != nil {
		return "", err
	}
	return patp, nil
}

func normalizeAddress(address string) (string, error) {
	addr := strings.TrimSpace(address)
	if addr == "" {
		return "", fmt.Errorf("address is required")
	}
	return roller.ValidateAddress(addr, false)
}

func normalizeTicket(ticket string) string {
	ticket = strings.TrimSpace(ticket)
	if ticket != "" && !strings.HasPrefix(ticket, "~") {
		return "~" + ticket
	}
	return ticket
}

func normalizeOperation(operation string) string {
	return strings.TrimSpace(strings.ToLower(operation))
}

func summarizePending(ship, operation string, tx any) *pendingSummary {
	raw, err := json.Marshal(tx)
	if err != nil {
		return nil
	}
	var values map[string]any
	if err := json.Unmarshal(raw, &values); err != nil {
		return nil
	}
	sig := stringFromAny(values["sig"])
	if sig == "" {
		sig = stringFromAny(values["Signature"])
	}
	hash := stringFromAny(values["hash"])
	if hash == "" {
		hash = stringFromAny(values["Hash"])
	}
	txType := stringFromAny(values["type"])
	if txType == "" {
		txType = stringFromAny(values["Type"])
	}
	if operation == "" {
		operation = txType
	}
	if sig == "" && hash == "" {
		return nil
	}
	now := time.Now()
	return &pendingSummary{
		Operation:    operation,
		Ship:         ship,
		Hash:         hash,
		Signature:    sig,
		Status:       "pending",
		SubmittedAt:  now.UnixMilli(),
		NextPollAt:   now.Add(time.Minute).UnixMilli(),
		PollInterval: 60,
	}
}

func stringFromAny(value any) string {
	if str, ok := value.(string); ok {
		return str
	}
	return ""
}

func walletOperationMessage(operation string) string {
	if operation == "breach" {
		return "Export this ship before booting from the new keys. A breach changes continuity."
	}
	return ""
}

func rollerRPC(ctx context.Context, target, method string, params any) (json.RawMessage, error) {
	endpoint, headers := rollerTransport(target)
	reqPayload := rpcRequest{
		Version: "2.0",
		Method:  method,
		Params:  params,
		ID:      fmt.Sprintf("%d", time.Now().UnixNano()),
	}
	body, err := json.Marshal(reqPayload)
	if err != nil {
		return nil, fmt.Errorf("marshal roller request: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create roller request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("roller request failed: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read roller response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("roller status %d: %s", resp.StatusCode, string(respBody))
	}
	var rpcResp rpcResponse
	if err := json.Unmarshal(respBody, &rpcResp); err != nil {
		return nil, fmt.Errorf("decode roller response: %w", err)
	}
	if rpcResp.Error != nil {
		return nil, fmt.Errorf("roller rpc error: %s", rpcResp.Error.Message)
	}
	return rpcResp.Result, nil
}
