package utils

import (
	"fmt"
	"strings"
)

func TruncateLines(lines []string) (truncs []string) {
	width := TerminalWidth()
	format := fmt.Sprintf("%%.%ds...", width-3)
	for _, line := range lines {
		line = strings.TrimRight(line, " ")
		if len(line) > width {
			truncs = append(truncs, fmt.Sprintf(format, line))
		} else {
			truncs = append(truncs, line)
		}
	}
	return
}
