package types

type PxSession struct {
	Ship       string
	Ticket     string
	Point      PointResp
	AuthType   string
	Passphrase string
	Wallet     *WalletData `json:",omitempty"`
}

type WalletData struct {
	Address string
}

type PxSigilConfig struct {
	Patp       string
	Size       int
	Background string
	Foreground string
}

type PxSvgConfig struct {
	Point      string
	Background string
	Foreground string
	Size       int
	Space      string
	Detail     string
}

type SvgSymbol struct {
	SVG string
}

type PxTransferFormData struct {
	Type string
}

type PxTransferRequest struct {
	TransferType string `json:"transfer_type"`
	Address      string `json:"address"`
}
