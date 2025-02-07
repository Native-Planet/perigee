package pontifex

import (
	"fmt"
	"strings"
)

// Symbol represents an SVG symbol template
type Symbol struct {
	SVG string
}

// Space represents the spacing mode for sigils
type Space string

const (
	SpaceNone    Space = "none"
	SpaceDefault Space = "default"
	SpaceLarge   Space = "large"
)

// Detail represents the detail mode for sigils
type Detail string

const (
	DetailNone    Detail = "none"
	DetailDefault Detail = "default"
)

// Config holds the configuration for sigil generation
type Config struct {
	Point      string
	Background string
	Foreground string
	Size       int
	Space      Space
	Detail     Detail
}

// DefaultConfig returns a default configuration
func DefaultConfig() Config {
	return Config{
		Size:       128,
		Background: "#000",
		Foreground: "#FFF",
		Space:      SpaceDefault,
		Detail:     DetailDefault,
	}
}

// chunkStr splits a string into equal-sized chunks
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

// cleanPatp removes special characters from patp names
func cleanPatp(point string) string {
	point = strings.ReplaceAll(point, "~", "")
	point = strings.ReplaceAll(point, "^", "")
	point = strings.ReplaceAll(point, "-", "")
	return point
}

// calculateTransform returns the transform for a symbol based on its position
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

// calculateGroupTransform calculates the group transformation based on phoneme length and space setting
func calculateGroupTransform(size int, phonemeCount int, space Space) string {
	switch space {
	case SpaceNone:
		switch phonemeCount {
		case 1:
			return "scale(2)"
		default:
			return ""
		}
	case SpaceLarge:
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
	default: // SpaceDefault
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

// GenerateSigil generates an SVG sigil based on the provided configuration
func GenerateSigil(cfg Config) (string, error) {
	if cfg.Point == "" {
		return "", fmt.Errorf("point must be provided")
	}

	// Apply defaults if not set
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

	// Process point name
	cleanPoint := cleanPatp(cfg.Point)
	phonemes := chunkStr(cleanPoint, 3)

	// Validate phoneme length
	if len(phonemes) != 1 && len(phonemes) != 2 && len(phonemes) != 4 {
		return "", fmt.Errorf("invalid point name length: %d", len(phonemes))
	}

	// Calculate height based on point type
	height := cfg.Size
	if cfg.Space == SpaceNone && len(phonemes) == 2 {
		height = cfg.Size / 2
	}

	// Calculate stroke width
	strokeWidth := "4"
	if cfg.Size < 64 {
		strokeWidth = fmt.Sprintf("%f", 256.0/float64(cfg.Size))
	}

	// Build inner SVG content
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

	// Generate final SVG
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
