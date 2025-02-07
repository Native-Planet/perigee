package pontifex

import (
	"fmt"
	"strings"

	"github.com/Native-Planet/perigee/types"
)

const (
	stringNone    string = "none"
	stringDefault string = "default"
	stringLarge   string = "large"
	DetailNone    string = "none"
	DetailDefault string = "default"
)

func DefaultConfig() types.SvgConfig {
	return types.SvgConfig{
		Size:       128,
		Background: "#000",
		Foreground: "#FFF",
		Space:      stringDefault,
		Detail:     DetailDefault,
	}
}

func chunkStr(str string, size int) []string {
	var chunks []string
	runes := []rune(str)

	for i := 0; i < len(runes); i += size {
		end := i + size
		if end > len(runes) {
			end = len(runes)
		}
		chunks = append(chunks, string(runes[i:end]))
	}

	if len(chunks) == 0 {
		return []string{""}
	}
	return chunks
}

func cleanPatp(point string) string {
	point = strings.ReplaceAll(point, "~", "")
	point = strings.ReplaceAll(point, "^", "")
	point = strings.ReplaceAll(point, "-", "")
	return point
}

func calculateTransform(size int, index int) string {
	scale := float64(size) / 256.0

	switch index {
	case 0:
		return fmt.Sprintf("scale(%f) translate(0,0)", scale)
	case 1:
		return fmt.Sprintf("scale(%f) translate(128,0)", scale)
	case 2:
		return fmt.Sprintf("scale(%f) translate(0,128)", scale)
	default:
		return fmt.Sprintf("scale(%f) translate(128,128)", scale)
	}
}

func calculateGroupTransform(size int, phonemeCount int, space string) string {
	switch space {
	case stringNone:
		switch phonemeCount {
		case 1:
			return "scale(2)"
		default:
			return ""
		}
	case stringLarge:
		switch phonemeCount {
		case 1:
			return fmt.Sprintf("translate(%f,%f) scale(0.50)",
				(float64(size)*0.5)-(float64(size)*0.125),
				(float64(size)*0.5)-(float64(size)*0.125))
		case 2:
			return fmt.Sprintf("translate(%f,%f) scale(0.50)",
				(float64(size)*0.5)-(float64(size)*0.25),
				(float64(size)*0.5)-(float64(size)*0.125))
		default:
			return fmt.Sprintf("translate(%f,%f) scale(0.50)",
				(float64(size)*0.5)-(float64(size)*0.25),
				(float64(size)*0.5)-(float64(size)*0.25))
		}
	default:
		switch phonemeCount {
		case 1:
			return fmt.Sprintf("translate(%f,%f) scale(0.75)",
				(float64(size)*0.5)-(float64(size)*0.1875),
				(float64(size)*0.5)-(float64(size)*0.1875))
		case 2:
			return fmt.Sprintf("translate(%f,%f) scale(0.75)",
				(float64(size)*0.5)-(float64(size)*0.3750),
				(float64(size)*0.5)-(float64(size)*0.1875))
		default:
			return fmt.Sprintf("translate(%f,%f) scale(0.75)",
				(float64(size)*0.5)-(float64(size)*0.3750),
				(float64(size)*0.5)-(float64(size)*0.3750))
		}
	}
}

func GenerateSigil(cfg types.SvgConfig) (string, error) {
	if cfg.Point == "" {
		return "", fmt.Errorf("point must be provided")
	}
	if cfg.Size == 0 {
		cfg.Size = DefaultConfig().Size
	}
	if cfg.Background == "" {
		cfg.Background = DefaultConfig().Background
	}
	if cfg.Foreground == "" {
		cfg.Foreground = DefaultConfig().Foreground
	}
	if cfg.Space == "" {
		cfg.Space = DefaultConfig().Space
	}
	if cfg.Detail == "" {
		cfg.Detail = DefaultConfig().Detail
	}
	cleanPoint := cleanPatp(cfg.Point)
	phonemes := chunkStr(cleanPoint, 3)
	if len(phonemes) != 1 && len(phonemes) != 2 && len(phonemes) != 4 {
		return "", fmt.Errorf("invalid point name length: %d", len(phonemes))
	}
	height := cfg.Size
	if cfg.Space == stringNone && len(phonemes) == 2 {
		height = cfg.Size / 2
	}
	strokeWidth := "4"
	if cfg.Size < 64 {
		strokeWidth = fmt.Sprintf("%f", 256.0/float64(cfg.Size))
	}
	var innerSVG strings.Builder
	for i, phoneme := range phonemes {
		symbol, ok := symbolIndex[phoneme]
		if !ok {
			return "", fmt.Errorf("invalid phoneme: %s", phoneme)
		}
		transform := calculateTransform(cfg.Size, i)
		symbolSVG := strings.ReplaceAll(symbol.SVG, "@FG", cfg.Foreground)
		symbolSVG = strings.ReplaceAll(symbolSVG, "@BG", cfg.Background)
		symbolSVG = strings.ReplaceAll(symbolSVG, "@TR", transform)
		symbolSVG = strings.ReplaceAll(symbolSVG, "@SW", strokeWidth)
		innerSVG.WriteString(symbolSVG)
	}

	groupTransform := calculateGroupTransform(cfg.Size, len(phonemes), cfg.Space)
	svg := fmt.Sprintf(`
    <svg
      style="display: block;"
      width="%d"
      height="%d"
      viewbox="0 0 %d %d"
      version="1.1"
      xmlns="http://www.w3.org/2000/svg"
    >
      <rect fill="%s" width="%d" height="%d" x="0" y="0" />
      <g transform="%s">
        %s
      </g>
    </svg>
  `, cfg.Size, height, cfg.Size, height,
		cfg.Background, cfg.Size, height,
		groupTransform, innerSVG.String())
	return svg, nil
}
