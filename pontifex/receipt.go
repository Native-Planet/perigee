package pontifex

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	totalWidth   = 80
	contentWidth = totalWidth - 4
	treeIndent   = 4
)

// convert json data into 80width txt
func formatToNFO(data interface{}, title string) string {
	var builder strings.Builder
	builder.WriteString("╔" + strings.Repeat("═", totalWidth-2) + "╗\n")
	builder.WriteString("║" + centerText(title, totalWidth-2) + "║\n")
	builder.WriteString("╠" + strings.Repeat("═", totalWidth-2) + "╣\n")
	builder.WriteString("║" + centerText(time.Now().Format("2006-01-02 15:04:05"), totalWidth-2) + "║\n")
	builder.WriteString("╟" + strings.Repeat("─", totalWidth-2) + "╢\n")
	var processNode func(interface{}, int) []string
	processNode = func(v interface{}, depth int) []string {
		var lines []string
		indent := strings.Repeat(" ", depth*treeIndent)
		switch val := v.(type) {
		case map[string]interface{}:
			keys := make([]string, 0, len(val))
			for k := range val {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for i, k := range keys {
				isLastKey := i == len(keys)-1
				treePrefix := "├── "
				if isLastKey {
					treePrefix = "└── "
				}
				fullLine := indent + treePrefix + k
				switch childVal := val[k].(type) {
				case map[string]interface{}, []interface{}:
					lines = append(lines, fullLine)
					childLines := processNode(childVal, depth+1)
					lines = append(lines, childLines...)
				default:
					if childVal == nil {
						lines = append(lines, fullLine+": <nil>")
					} else {
						strVal := fmt.Sprintf("%v", childVal)
						prefixForValue := fullLine + ": "
						available := contentWidth - utf8.RuneCountInString(prefixForValue)
						if available < 1 {
							available = 1
						}
						wrapped := wrapText(strVal, available)
						lines = append(lines, prefixForValue+wrapped[0])
						contPrefix := strings.Repeat(" ", utf8.RuneCountInString(prefixForValue))
						for _, extra := range wrapped[1:] {
							lines = append(lines, contPrefix+extra)
						}
					}
				}
			}
		case []interface{}:
			for i, item := range val {
				isLastItem := i == len(val)-1
				treePrefix := "├── "
				if isLastItem {
					treePrefix = "└── "
				}
				fullLine := indent + treePrefix + fmt.Sprintf("[%d]", i)
				switch childVal := item.(type) {
				case map[string]interface{}, []interface{}:
					lines = append(lines, fullLine)
					childLines := processNode(childVal, depth+1)
					lines = append(lines, childLines...)
				default:
					if childVal == nil {
						lines = append(lines, fullLine+": <nil>")
					} else {
						strVal := fmt.Sprintf("%v", childVal)
						prefixForValue := fullLine + ": "
						available := contentWidth - utf8.RuneCountInString(prefixForValue)
						if available < 1 {
							available = 1
						}
						wrapped := wrapText(strVal, available)
						lines = append(lines, prefixForValue+wrapped[0])
						contPrefix := strings.Repeat(" ", utf8.RuneCountInString(prefixForValue))
						for _, extra := range wrapped[1:] {
							lines = append(lines, contPrefix+extra)
						}
					}
				}
			}
		}
		return lines
	}
	var dataMap map[string]interface{}
	jsonBytes, _ := json.Marshal(data)
	json.Unmarshal(jsonBytes, &dataMap)
	lines := processNode(dataMap, 0)
	for _, line := range lines {
		currentWidth := utf8.RuneCountInString(line)
		if currentWidth < contentWidth {
			line += strings.Repeat(" ", contentWidth-currentWidth)
		}
		builder.WriteString("║ " + line + " ║\n")
	}
	builder.WriteString("╚" + strings.Repeat("═", totalWidth-2) + "╝\n")
	return builder.String()
}

func wrapText(text string, width int) []string {
	if utf8.RuneCountInString(text) <= width {
		return []string{text}
	}
	var lines []string
	for len(text) > 0 {
		currentWidth := 0
		breakPos := 0
		lastSpacePos := -1
		for pos, r := range text {
			currentWidth++
			if r == ' ' {
				lastSpacePos = pos
			}
			if currentWidth > width {
				break
			}
			breakPos = pos + 1
		}
		if currentWidth > width && lastSpacePos != -1 {
			breakPos = lastSpacePos
		}
		if breakPos == 0 {
			breakPos = 1
		}
		line := strings.TrimRight(text[:breakPos], " ")
		lines = append(lines, line)
		text = strings.TrimLeft(text[breakPos:], " ")
	}
	return lines
}

func centerText(text string, width int) string {
	textLen := utf8.RuneCountInString(text)
	padding := (width - textLen) / 2
	result := strings.Repeat(" ", padding) + text
	extra := width - utf8.RuneCountInString(result)
	result += strings.Repeat(" ", extra)
	return result
}
