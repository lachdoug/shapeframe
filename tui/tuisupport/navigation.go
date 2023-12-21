package tuisupport

import tea "github.com/charmbracelet/bubbletea"

type NavigationMsg struct {
	Path string
}

type ReloadMsg struct{}

func Open(path string) (c tea.Cmd) {
	c = func() tea.Msg { return NavigationMsg{Path: path} }
	return
}

func Reload() (c tea.Cmd) {
	c = func() tea.Msg { return ReloadMsg{} }
	return
}
