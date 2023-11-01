package utils

import (
	"os"

	"golang.org/x/term"
)

func TerminalWidth() (width int) {
	width, _, err := term.GetSize(int(os.Stdin.Fd()))
	checkErr(err)
	return
}
