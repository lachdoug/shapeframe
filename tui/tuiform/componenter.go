package tuiform

import (
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type Componenter interface {
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string
	FocusChain() []tuisupport.Focuser
	resize()
	depend()
	validity() string
	shown() []string
	// Focus(...string) tea.Cmd
	// Blur()
	// width() int
	// enter() tea.Cmd
	// next() tea.Cmd
	// previous() tea.Cmd
	// set(string, string)
	// value() string
}
