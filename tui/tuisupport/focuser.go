package tuisupport

import tea "github.com/charmbracelet/bubbletea"

type Focuser interface {
	Focus(string) tea.Cmd
	Blur()
}

type TakeFocus Focuser

func TakeFocusCommand(target Focuser) (c tea.Cmd) {
	c = func() tea.Msg { return TakeFocus(target) }
	return
}
