package cli

import (
	"fmt"
	"sf/app/io"
)

// Add green color codes to text
func green(in string) (out string) {
	out = fmt.Sprintf("%s%s%s", io.GreenColor, in, io.ResetText)
	return
}

// Slice of strings
func ss(s ...string) []string {
	return s
}

// Slice of commands
func cs(c ...func() any) []func() any {
	return c
}

// Slice of table accent functions
func tas(c ...func(string, map[string]any) string) []func(string, map[string]any) string {
	return c
}
