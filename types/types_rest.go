package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"strconv"
	"strings"

	"github.com/deelawn/urbit-gob/co"
)

type WalletResp struct {
	Wallet interface{} `json:"wallet"`
	Error  string      `json:"error,omitempty"`
}

type PointResp struct {
	Point    *Point `json:"point"`
	PatpName string `json:"patp"`
}

type BreachReq struct {
	Point      json.RawMessage `json:"point"`
	Ticket     string          `json:"ticket"`
	Passphrase string          `json:"passphrase"`
}

func UnmarshalPoint(raw json.RawMessage) (string, uint32, error) {
	var strPoint string
	if err := json.Unmarshal(raw, &strPoint); err == nil {
		return validateAndNormalizePatp(strPoint)
	}
	var numPoint int64
	if err := json.Unmarshal(raw, &numPoint); err == nil {
		if numPoint < 0 {
			return "", 0, fmt.Errorf("point cannot be negative")
		}
		patp, err := co.Point2Patp(big.NewInt(numPoint))
		if err != nil {
			return "", 0, fmt.Errorf("invalid point number: %v", err)
		}
		return ensureTildePrefix(patp), uint32(numPoint), nil
	}

	return "", 0, fmt.Errorf("point must be a string or number")
}

// get the point from a url param string
func ParsePointParam(param string) (string, uint32, error) {
	if param == "" {
		return "", 0, fmt.Errorf("point parameter is required")
	}
	decoded, err := url.QueryUnescape(param)
	if err != nil {
		return "", 0, fmt.Errorf("invalid URL encoding: %v", err)
	}
	if num, err := strconv.ParseInt(decoded, 10, 64); err == nil {
		if num < 0 {
			return "", 0, fmt.Errorf("point cannot be negative")
		}
		patp, err := co.Point2Patp(big.NewInt(num))
		if err != nil {
			return "", 0, fmt.Errorf("invalid point number: %v", err)
		}
		return ensureTildePrefix(patp), uint32(num), nil
	}
	patp, point, err := validateAndNormalizePatp(decoded)
	if err != nil {
		return "", 0, err
	}
	return ensureTildePrefix(patp), point, nil
}

// get a point and an int back
func validateAndNormalizePatp(patp string) (string, uint32, error) {
	cleanPatp := strings.TrimPrefix(patp, "~")
	point, err := co.Patp2Point("~" + cleanPatp)
	if err != nil {
		return "", 0, fmt.Errorf("invalid patp: %v", err)
	}
	validPatp, err := co.Point2Patp(point)
	if err != nil {
		return "", 0, fmt.Errorf("error normalizing patp: %v", err)
	}

	return ensureTildePrefix(validPatp), uint32(point.Uint64()), nil
}

func ensureTildePrefix(patp string) string {
	if !strings.HasPrefix(patp, "~") {
		return "~" + patp
	}
	return patp
}
