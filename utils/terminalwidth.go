package utils

import (
	"os"

	"golang.org/x/term"
)

func TerminalSize() (width int, height int) {
	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	checkErr(err)
	return
}
