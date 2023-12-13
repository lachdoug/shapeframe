package tuiform

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Select struct {
	Form      *Form
	Field     *Field
	Items     []*Option
	Selection int
	Width     int
}

func NewSelect(fm *Form, f *Field) (s *Select) {
	s = &Select{Form: fm, Field: f}
	return
}

func (s *Select) setWidth() {
	s.Width = s.Field.Width
}

func (s *Select) setItems() {
	items := []*Option{}
	for _, opt := range s.Field.ComponentModel.Options {
		items = append(items, NewOption(opt))
	}
	s.Items = items
}

func (s *Select) setSelection() {
	an := s.Form.Answers[s.Field.ComponentModel.Key]
	for i, item := range s.Field.ComponentModel.Options {
		if an == item.Value {
			s.Selection = i
			return
		}
	}
}

func (s *Select) Init() (c tea.Cmd) {
	s.setWidth()
	s.setItems()
	s.setSelection()
	return
}

func (s *Select) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = s
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeySpace:
			c = s.enter()
			return
		// case tea.KeyTab:
		// 	c = s.next()
		// 	return
		// case tea.KeyShiftTab:
		// 	c = s.previous()
		// 	return
		case tea.KeyDown:
			s.down()
			return
		case tea.KeyUp:
			s.up()
			return
		}
	case error:
		panic(msg)
	}
	return
}

func (s *Select) View() (v string) {
	borderStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Width(s.Width - 2)
	if s.Field.IsFocus {
		borderStyle.BorderForeground(lipgloss.Color("15"))
	} else {
		borderStyle.BorderForeground(lipgloss.Color("8"))
	}

	v = borderStyle.Render(s.body())
	return
}

func (s *Select) body() (f string) {
	if !s.Field.IsFocus {
		f = s.Items[s.Selection].View(false, "")
		return
	}
	lines := []string{}
	for i, it := range s.Items {
		lines = append(lines, it.View(i == s.Selection, "⮞"))
	}
	f = strings.Join(lines, "\n")
	return
}

func (s *Select) answer() {
	s.Field.set(s.Field.ComponentModel.Key, s.value())
}

func (s *Select) width() (w int) {
	w = s.Width
	return
}

func (s *Select) resize() {
	s.setWidth()
}

func (s *Select) enter() (c tea.Cmd) {
	c = s.Field.enter()
	return
}

// func (s *Select) next() (c tea.Cmd) {
// 	c = s.Field.next()
// 	return
// }

// func (s *Select) previous() (c tea.Cmd) {
// 	c = s.Field.previous()
// 	return
// }

func (s *Select) focus() tea.Cmd { return nil }
func (s *Select) blur()          {}

// func (s *Select) depend()            {}
func (s *Select) set(string, string) {}

func (s *Select) value() (v string) {
	v = s.Items[s.Selection].Value
	return
}

// func (s *Select) validity() string { return "" }

func (s *Select) up() {
	s.Selection--
	if s.Selection == -1 {
		s.Selection++
	}
	s.answer()
}

func (s *Select) down() {
	s.Selection++
	if s.Selection == len(s.Items) {
		s.Selection--
	}
	s.answer()
}

// func (s *Select) shown() []string { return nil }
