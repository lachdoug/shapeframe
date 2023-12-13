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

// func (nf *NotFound) focus(on ...string) (c tea.Cmd) {
// 	if slices.Contains(on, "next") {
// 		c = nf.next()
// 	} else {
// 		c = nf.previous()
// 	}
// 	return
// }

// func (nf *NotFound) next() (c tea.Cmd) {
// 	c = nf.Body.next()
// 	return
// }

// func (nf *NotFound) previous() (c tea.Cmd) {
// 	c = nf.Body.previous()
// 	return
// }

// func (vn *NotFound) blur()            {}
func (nf *NotFound) setSize(int, int) {}
