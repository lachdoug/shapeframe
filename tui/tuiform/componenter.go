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
	isFocus() bool
	setWidth(int)
	depend()
	validity() string
	shown() []string
}
