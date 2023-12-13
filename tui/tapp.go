package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func newTApp(app *App) (tapp *tea.Program) {
	zone.NewGlobal()
	tapp = tea.NewProgram(
		app,
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	// tea.WithoutCatchPanics(),
	)
	return
}
