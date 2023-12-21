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

func (btns *Buttons) setWidth(int) {}

func (btns *Buttons) FocusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{
		btns.Cancel,
		btns.Submit,
	}
	return
}

func (btns *Buttons) isFocus() (is bool) {
	is = btns.Cancel.IsFocus || btns.Submit.IsFocus
	return
}

func (btns *Buttons) depend() {
	if btns.Form.validity() == "" {
		btns.Submit.Enabled(true)
	} else {
		btns.Submit.Enabled(false)
	}
}

func (btns *Buttons) validity() string { return "" }

func (btns *Buttons) shown() []string { return nil }
