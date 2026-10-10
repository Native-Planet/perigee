package types

// PlanetCode claims an L2 invite in Bridge. Bridge creates a fresh master ticket
// in the recipient's browser and transfers the planet to that wallet.
type PlanetCode struct {
	Planet string `json:"planet"`
	Code   string `json:"code"`
	Owner  string `json:"owner"`
	URL    string `json:"url"`
}
