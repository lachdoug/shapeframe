package tuiform

import tea "github.com/charmbracelet/bubbletea"

type Inputter interface {
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string
	focus() tea.Cmd
	blur()
	resize()
	value() string
	// validity() string
}
