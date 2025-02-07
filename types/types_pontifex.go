package types

type Session struct {
	Ship   string
	Ticket string
	Point  PointResp
}

type SigilConfig struct {
	Patp       string
	Size       int
	Background string
	Foreground string
}

type SvgConfig struct {
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
