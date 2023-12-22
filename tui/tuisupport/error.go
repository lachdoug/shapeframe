package tuisupport

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Error struct {
	Message  string
	OK       *Button
	Callback tea.Cmd
}

type ErrorMsg struct {
	Err      error
	Callback tea.Cmd
}

type ClearErrorMsg struct{}

func NewErrorMsg(err error, callback tea.Cmd) (emsg ErrorMsg) {
	emsg = ErrorMsg{Err: err, Callback: callback}
	return
}

func NewError(err error, callback tea.Cmd) (e *Error) {
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
	e.OK = NewButton("error-ok", " OK ", e.sendClearMsg, 15)
}

func (e *Error) sendClearMsg() (c tea.Cmd) {
	c = tea.Batch(
		e.Callback,
	)
	return
}

func (e *Error) FocusChain() (fc []Focuser) {
	fc = []Focuser{e.OK}
	return
}
