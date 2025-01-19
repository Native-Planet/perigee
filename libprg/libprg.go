package libprg

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Native-Planet/perigee/roller"
	"github.com/Native-Planet/perigee/types"
	"github.com/nathanlever/keygen"
)

var (
	ctx = context.Background()
)

func GetKeyfile(point, masterTicket, life string) (string, error) {
	patp, pointInt, err := types.ValidateAndNormalizePatp(point)
	if err != nil {
		return "", fmt.Errorf("invalid point")
	}
	pInfo, err := roller.Client.GetPoint(ctx, patp)
	if err != nil {
		return "", fmt.Errorf("error getting point: %v", err)
	}
	var rev string
	if life != "" {
		rev = life
	} else {
		rev = fmt.Sprintf("%v", pInfo.Network.Keys.Life)
	}
	var lifeInt int
	lifeInt, err = strconv.Atoi(rev)
	if err != nil {
		return "", fmt.Errorf("invalid life value: %v", err)
	}
	lifeInt -= 1
	wallet := keygen.GenerateWallet(masterTicket, uint32(pointInt), "", uint(lifeInt), true)
	pointKey := strings.TrimPrefix(pInfo.Network.Keys.Crypt, "0x")
	if wallet.Network.Keys.Crypt.Public != pointKey {
		return "", fmt.Errorf("could not generate public key matching PKI; 0x%s / %s", wallet.Network.Keys.Crypt.Public, pInfo.Network.Keys.Crypt)
	}
	lifeInt += 1
	keyfile, err := roller.Keyfile(wallet.Network.Keys.Crypt.Private, wallet.Network.Keys.Auth.Private, patp, lifeInt)
	if err != nil {
		return "", fmt.Errorf("error generating keyfile: %v", err)
	}
	return keyfile, nil
}
