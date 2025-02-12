package types

type PxSession struct {
	Ship       string `json:"ship"`
	Ticket     string `json:"ticket"`
	Point      PointResp
	AuthType   string `json:"authType"`
	Passphrase string `json:"passphrase"`
	Wallet     string `json:"wallet"`
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
