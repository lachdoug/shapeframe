package utils

import (
	"strings"
)

func IndentLines(in string, indents ...int) (out string) {
	lines := strings.Split(in, "\n")
	var padN, pad1 string
	if len(indents) == 0 {
		padN, pad1 = "", ""
	} else if len(indents) == 1 {
		padN = strings.Repeat(" ", indents[0])
		pad1 = padN
	} else {
		padN = strings.Repeat(" ", indents[0])
		pad1 = strings.Repeat(" ", indents[1])
	}
	for i, line := range lines {
		if line == "\n" {
			continue // Do not indent blank lines
		}
		if i == 0 {
			lines[i] = pad1 + line
		} else {
			lines[i] = padN + line
		}
	}
	out = strings.Join(lines, "\n")
	return
}
