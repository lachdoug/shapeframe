package tuiform

import (
	"fmt"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Buttons struct {
	Form   *Form
	Cancel *tuisupport.Button
	Submit *tuisupport.Button
}

func newButtons(fm *Form) (btns *Buttons) {
	btns = &Buttons{Form: fm}
	return
}

func (btns *Buttons) Init() (c tea.Cmd) {
	btns.setCancel()
	btns.setSubmit()
	return
}

func (btns *Buttons) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = btns
	cs := []tea.Cmd{}
	_, c = btns.Cancel.Update(msg)
	cs = append(cs, c)
	_, c = btns.Submit.Update(msg)
	cs = append(cs, c)
	c = tea.Batch(cs...)
	return
}

func (btns *Buttons) View() (v string) {
	v = lipgloss.JoinHorizontal(lipgloss.Top,
		btns.Cancel.View(),
		btns.Submit.View(),
	)
	return
}

func (btns *Buttons) setCancel() {
	btns.Cancel = tuisupport.NewButton(
		fmt.Sprintf("%s-cancel", btns.Form.ID),
		"Cancel",
		btns.Form.cancel,
		9,
	)
}

func (btns *Buttons) setSubmit() {
	btns.Submit = tuisupport.NewButton(
		fmt.Sprintf("%s-submit", btns.Form.ID),
		"Submit",
		btns.Form.submit,
		15,
	)
}

// func (btns *Buttons) width() int         { return 0 }
func (btns *Buttons) resize() {}

// func (btns *Buttons) enter() (c tea.Cmd) { return nil }

// func (btns *Buttons) next() (c tea.Cmd) {
// 	btns.Blur()
// 	btns.FocusIndex++
// 	if btns.FocusIndex == 2 {
// 		btns.FocusIndex = 1
// 		c = btns.Form.next()
// 	} else {
// 		c = btns.Focus("next")
// 	}
// 	return
// }

// func (btns *Buttons) previous() (c tea.Cmd) {
// 	btns.Blur()
// 	btns.FocusIndex--
// 	if btns.FocusIndex == -1 {
// 		btns.FocusIndex = 0
// 		c = btns.Form.previous()
// 	} else {
// 		c = btns.Focus("previous")
// 	}
// 	return
// }

// func (btns *Buttons) Focus(on ...string) (c tea.Cmd) {
// 	if slices.Contains(on, "first") {
// 		btns.FocusIndex = 0
// 	} else if slices.Contains(on, "last") {
// 		btns.FocusIndex = 1
// 	}
// 	if btns.FocusIndex == 0 {
// 		c = btns.Submit.Focus(on...)
// 	} else {
// 		c = btns.Cancel.Focus(on...)
// 	}
// 	return
// }

// func (btns *Buttons) Blur() {
// 	btns.Submit.Blur()
// 	btns.Cancel.Blur()
// }

func (btns *Buttons) FocusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{
		btns.Cancel,
		btns.Submit,
	}
	return
}

func (btns *Buttons) depend() {
	if btns.Form.validity() == "" {
		btns.Submit.Enabled(true)
	} else {
		btns.Submit.Enabled(false)
	}
}

// func (btns *Buttons) set(string, string) {}
// func (btns *Buttons) value() string      { return "" }
func (btns *Buttons) validity() string { return "" }

func (btns *Buttons) shown() []string { return nil }
