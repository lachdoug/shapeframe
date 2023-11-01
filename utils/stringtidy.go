package utils

import (
	"regexp"
	"strings"
)

// Remove nonprintable and leading/trailing whitespace chars
func StringTidy(in string) (out string) {
	out = regexp.MustCompile(`[\t\n]+`).ReplaceAllString(strings.TrimSpace(in), "")
	return
}
