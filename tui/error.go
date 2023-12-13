package tui

import (
	"fmt"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Error struct {
	Message  string
	OK       *tuisupport.Button
	Callback func() tea.Cmd
}

func newError(err error, callback func() tea.Cmd) (e *Error) {
	e = &Error{Message: err.Error(), Callback: callback}
	return
}

func (e *Error) Init() (c tea.Cmd) {
	e.setOK()
	return
}

func (e *Error) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m, c = e.OK.Update(msg)
	return
}

func (e *Error) View() (v string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")).
		Padding(1)
	v = lipgloss.JoinVertical(lipgloss.Left,
		style.Render(fmt.Sprintf("Error: %s", e.Message)),
		e.OK.View(),
	)
	return
}

func (e *Error) setOK() {
	e.OK = tuisupport.NewButton("error-ok", " OK ", e.Callback, 15)
}
