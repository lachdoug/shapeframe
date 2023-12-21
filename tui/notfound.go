package tui

import (
	"fmt"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type NotFound struct {
	Body *Body
}

func newNotFound(b *Body) (nf *NotFound) {
	nf = &NotFound{Body: b}
	return
}

func (nf *NotFound) Init() (c tea.Cmd) {
	return
}

func (nf *NotFound) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	return
}

func (nf *NotFound) View() (v string) {
	v = fmt.Sprintf("Not found: %s", nf.Body.App.Path)
	return
}

func (nf *NotFound) focusChain() (fc []tuisupport.Focuser) {
	return
}

func (nf *NotFound) isFocus() (is bool) {
	return
}

func (nf *NotFound) setSize(int, int) {}
