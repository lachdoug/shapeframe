package tui

import (
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

func Open(path string) (c tea.Cmd) {
	c = func() tea.Msg { return tuisupport.Navigation{Path: path} }
	return
}
