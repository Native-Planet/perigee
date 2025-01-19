package roller

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/http"
	"os"

	"github.com/Native-Planet/perigee/aura"
	"github.com/Native-Planet/perigee/types"

	"github.com/stevelacy/go-urbit/noun"
)

var (
	RollerURL string
	Client    *Roller
)

func init() {
	if os.Getenv("ROLLER_URL") == "" {
		RollerURL = "https://roller.urbit.org/v1/roller"
	} else {
		RollerURL = os.Getenv("ROLLER_URL")
	}
	conf := Config{Endpoint: RollerURL, HTTPClient: http.DefaultClient}
	Client = New(conf)

}

type Roller struct {
	client *types.Client
}

type Config struct {
	Endpoint   string
	HTTPClient *http.Client
}

func New(cfg Config) *Roller {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}
	client := &types.Client{
		Endpoint:   cfg.Endpoint,
		HttpClient: cfg.HTTPClient,
	}
	return &Roller{
		client: client,
	}
}

// read operations

func (r *Roller) GetPoint(ctx context.Context, point interface{}) (*types.Point, error) {
	return r.client.GetPoint(ctx, point)
}

func (r *Roller) GetShips(ctx context.Context, address string) ([]types.ShipInfo, error) {
	return r.client.GetShips(ctx, address)
}

func (r *Roller) GetSpawned(ctx context.Context, point interface{}) ([]types.ShipInfo, error) {
	return r.client.GetSpawned(ctx, point)
}

func (r *Roller) GetUnspawned(ctx context.Context, point interface{}) ([]types.ShipInfo, error) {
	return r.client.GetUnspawned(ctx, point.(string))
}

// update operations

func (r *Roller) Spawn(ctx context.Context, parentPoint, spawnPoint interface{}, newOwnerAddress, signingAddress string, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	return r.client.Spawn(ctx, parentPoint, spawnPoint, newOwnerAddress, signingAddress, privateKey)
}

func (r *Roller) TransferPoint(ctx context.Context, point string, reset bool, newOwnerAddress, signingAddress string, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	return r.client.TransferPoint(ctx, point, reset, newOwnerAddress, signingAddress, privateKey)
}

func (r *Roller) ConfigureKeys(ctx context.Context, point, encryptPublic, authPublic string, breach bool, signingAddress string, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	return r.client.ConfigureKeys(ctx, point, encryptPublic, authPublic, breach, signingAddress, privateKey)
}

func (r *Roller) Escape(ctx context.Context, point, sponsor, signingAddress string, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	return r.client.Escape(ctx, point, sponsor, signingAddress, privateKey)
}

func (r *Roller) CancelEscape(ctx context.Context, point, sponsor, signingAddress string, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	return r.client.CancelEscape(ctx, point, sponsor, signingAddress, privateKey)
}

func (r *Roller) Adopt(ctx context.Context, point, adoptee, signingAddress string, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	return r.client.Adopt(ctx, point, adoptee, signingAddress, privateKey)
}

func (r *Roller) SetManagementProxy(ctx context.Context, point, proxyAddress, signingAddress string, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	return r.client.SetManagementProxy(ctx, point, proxyAddress, signingAddress, privateKey)
}

// roller meta

func (r *Roller) GetRollerConfig(ctx context.Context) (*types.RollerConfig, error) {
	return r.client.GetRollerConfig(ctx)
}

func (r *Roller) GetAllPending(ctx context.Context) ([]types.PendingTx, error) {
	return r.client.GetAllPending(ctx)
}

func (r *Roller) GetPendingByAddress(ctx context.Context, address string) ([]types.PendingTx, error) {
	return r.client.GetPendingByAddress(ctx, address)
}

func (r *Roller) WhenNextBatch(ctx context.Context) (*types.BatchInfo, error) {
	return r.client.WhenNextBatch(ctx)
}

// check permission

func (r *Roller) CanConfigureKeys(ctx context.Context, point, address string) (bool, error) {
	return r.client.CanConfigureKeys(ctx, point, address)
}

func (r *Roller) CanTransfer(ctx context.Context, point, address string) (bool, error) {
	return r.client.CanTransfer(ctx, point, address)
}

func (r *Roller) CanSpawn(ctx context.Context, point, address string) (bool, error) {
	return r.client.CanSpawn(ctx, point, address)
}

// validation

func ValidateAddress(address string, strip bool) (string, error) {
	return types.ValidateAddress(address, strip)
}

func ValidatePoint(point interface{}, strip bool) (*big.Int, error) {
	return types.ValidatePoint(point, strip)
}

func ValidatePrivateKey(privateKey *ecdsa.PrivateKey) error {
	return types.ValidatePrivateKey(privateKey)
}

func keyNoun(point *big.Int, revision int, bnsec *big.Int) noun.Cell {
	return noun.MakeNoun([]interface{}{
		[]interface{}{1, 0},            // [1 0]
		point,                          // point/ship number
		[]interface{}{revision, bnsec}, // [revision bnsec]
		0,                              // final 0
	}).(noun.Cell)
}

// generate a @uw keyfile
//   - the crypt and auth private keys are generated using
//     the *previous* life value of the point
//   - crypt and auth keys represented as hex strings are
//     concatenated with a 42 at the end
//   - the resulting hex value is cast as a bigint
//   - we compose the noun (see func keyNoun) and then jam
//   - @uw encode the jammed noun
func Keyfile(crypt, auth string, ship interface{}, revision int) (string, error) {
	point, err := ValidatePoint(ship, false)
	if err != nil {
		return "", err
	}
	if ship == nil {
		return "", fmt.Errorf("invalid point (nil pointer)")
	}
	ring := crypt + auth + "42"
	bnsec := new(big.Int)
	bnsec, ok := bnsec.SetString(ring, 16)
	if !ok {
		return "", fmt.Errorf("failed to parse ring as hex: %s", ring)
	}
	sed := keyNoun(point, revision, bnsec)
	jammed := noun.Jam(sed)
	jammedNoun := deref(jammed)
	encoded := aura.Big2Uw(jammedNoun)
	return encoded, nil
}

func deref(n *big.Int) *big.Int {
	result := new(big.Int)
	return result.Set(n)
}
