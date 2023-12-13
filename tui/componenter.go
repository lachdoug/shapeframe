package tui

import (
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type Componenter interface {
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string
	setSize(int, int)
	focusChain() []tuisupport.Focuser
	// focus() tea.Cmd
	// blur()
}
