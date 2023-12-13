package tui2

import (
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func NewTApp(app *App) (tapp *tea.Program) {
	zone.NewGlobal()
	tapp = tea.NewProgram(
		app,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	// tea.WithoutCatchPanics(),
	)
	return
}
