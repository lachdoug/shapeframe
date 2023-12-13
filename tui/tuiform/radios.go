package tuiform

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Radios struct {
	Form      *Form
	Field     *Field
	Items     []*Option
	Selection int
	Width     int
}

func NewRadios(fm *Form, f *Field) (rs *Radios) {
	rs = &Radios{Form: fm, Field: f}
	return
}

func (rs *Radios) setWidth() {
	rs.Width = rs.Field.Width
}

func (rs *Radios) setItems() {
	items := []*Option{}
	for _, opt := range rs.Field.ComponentModel.Options {
		items = append(items, NewOption(opt))
	}
	rs.Items = items
}

func (rs *Radios) setSelection() {
	an := rs.Form.Answers[rs.Field.ComponentModel.Key]
	for i, item := range rs.Field.ComponentModel.Options {
		if an == item.Value {
			rs.Selection = i
			return
		}
	}
}

func (rs *Radios) Init() (c tea.Cmd) {
	rs.setWidth()
	rs.setItems()
	rs.setSelection()
	return
}

func (rs *Radios) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = rs
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeySpace:
			c = rs.enter()
			return
		// case tea.KeyTab:
		// 	c = rs.next()
		// 	return
		// case tea.KeyShiftTab:
		// 	c = rs.previous()
		// 	return
		case tea.KeyDown:
			rs.down()
			return
		case tea.KeyUp:
			rs.up()
			return
		}
	case error:
		panic(msg)
	}
	return
}

func (rs *Radios) View() (v string) {
	style := lipgloss.NewStyle().
		PaddingLeft(1).
		Width(rs.Width)
	if rs.Field.IsFocus {
		style.BorderForeground(lipgloss.Color("15"))
	} else {
		style.BorderForeground(lipgloss.Color("8"))
	}

	v = style.Render(rs.body())
	return
}

func (rs *Radios) body() (f string) {
	lines := []string{}
	for i, it := range rs.Items {
		lines = append(lines, rs.line(it, i == rs.Selection))
	}
	f = strings.Join(lines, "\n")
	return
}

func (rs *Radios) line(it *Option, isSelected bool) (s string) {
	b := "○"
	if isSelected {
		b = "◉"
	}
	s = it.View(rs.Field.IsFocus && isSelected, b)
	return
}

func (rs *Radios) answer() {
	rs.Field.set(rs.Field.ComponentModel.Key, rs.value())
}

// func (rs *Radios) width() (w int) {
// 	w = rs.Width
// 	return
// }

func (rs *Radios) resize() {
	rs.setWidth()
}

func (rs *Radios) enter() (c tea.Cmd) {
	c = rs.Field.enter()
	return
}

// func (rs *Radios) next() (c tea.Cmd) {
// 	c = rs.Field.next()
// 	return
// }

// func (rs *Radios) previous() (c tea.Cmd) {
// 	c = rs.Field.previous()
// 	return
// }

func (rs *Radios) focus() tea.Cmd { return nil }
func (rs *Radios) blur()          {}

// func (rs *Radios) depend()                       {}
// func (rs *Radios) set(string, string)            {}

func (rs *Radios) value() (v string) {
	v = rs.Items[rs.Selection].Value
	return
}

// func (rs *Radios) validity() string { return "" }

func (rs *Radios) up() {
	rs.Selection--
	if rs.Selection == -1 {
		rs.Selection++
	}
	rs.answer()
}

func (rs *Radios) down() {
	rs.Selection++
	if rs.Selection == len(rs.Items) {
		rs.Selection--
	}
	rs.answer()
}

// func (rs *Radios) shown() []string { return nil }
