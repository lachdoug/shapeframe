package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func Run() {
	if _, err := newTApp(newApp()).Run(); err != nil {
		panic(err)
	}
}

func newTApp(app *App) (tapp *tea.Program) {
	zone.NewGlobal()
	tapp = tea.NewProgram(
		app,
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)
	return
}
