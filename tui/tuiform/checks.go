package tuiform

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/slices"
)

type Checks struct {
	Form           *Form
	Field          *Field
	Items          []*Option
	SelectionIndex int
	Selections     []int
	Width          int
}

func NewChecks(fm *Form, f *Field) (cs *Checks) {
	cs = &Checks{Form: fm, Field: f}
	return
}

func (cs *Checks) setWidth() {
	cs.Width = cs.Field.Width
}

func (cs *Checks) setItems() {
	items := []*Option{}
	for _, opt := range cs.Field.ComponentModel.Options {
		items = append(items, NewOption(opt))
	}
	cs.Items = items
}

func (cs *Checks) setSelections() {
	ans := strings.Split(cs.Form.Answers[cs.Field.ComponentModel.Key], "\n")
	for i, item := range cs.Field.ComponentModel.Options {
		if slices.Contains(ans, item.Value) {
			cs.Selections = append(cs.Selections, i)
		}
	}
}

func (cs *Checks) Init() (c tea.Cmd) {
	cs.setWidth()
	cs.setItems()
	cs.setSelections()
	return
}

func (cs *Checks) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = cs
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			c = cs.enter()
			return
		case tea.KeySpace:
			cs.selectItem()
			cs.answer()
			return
		// case tea.KeyTab:
		// 	c = cs.next()
		// 	return
		// case tea.KeyShiftTab:
		// 	c = cs.previous()
		// 	return
		case tea.KeyDown:
			cs.down()
			return
		case tea.KeyUp:
			cs.up()
			return
		}
	case error:
		panic(msg)
	}
	return
}

func (cs *Checks) View() (v string) {
	style := lipgloss.NewStyle().
		PaddingLeft(1).
		Width(cs.Width)
	if cs.Field.IsFocus {
		style.BorderForeground(lipgloss.Color("15"))
	} else {
		style.BorderForeground(lipgloss.Color("8"))
	}

	v = style.Render(cs.body())
	return
}

func (cs *Checks) body() (f string) {
	lines := []string{}
	for i, it := range cs.Items {
		lines = append(lines,
			cs.line(it,
				slices.Contains(cs.Selections, i),
				cs.Field.IsFocus && i == cs.SelectionIndex,
			),
		)
	}
	f = strings.Join(lines, "\n")
	return
}

func (cs *Checks) line(it *Option, isSelected bool, isFocused bool) (s string) {
	b := "☐"
	if isSelected {
		b = "🗹"
	}
	s = it.View(isFocused, b)
	return
}

func (cs *Checks) selectItem() {
	if slices.Contains(cs.Selections, cs.SelectionIndex) {
		cs.Selections = slices.DeleteFunc(cs.Selections, func(i int) bool { return cs.SelectionIndex == i })
	} else {
		cs.Selections = append(cs.Selections, cs.SelectionIndex)
		slices.Sort(cs.Selections)
	}
}

func (cs *Checks) answer() {
	cs.Field.set(cs.Field.ComponentModel.Key, cs.value())
}

func (cs *Checks) width() (w int) {
	w = cs.Width
	return
}

func (cs *Checks) resize() {
	cs.setWidth()
}

func (cs *Checks) enter() (c tea.Cmd) {
	c = cs.Field.enter()
	return
}

// func (cs *Checks) next() (c tea.Cmd) {
// 	c = cs.Field.next()
// 	return
// }

// func (cs *Checks) previous() (c tea.Cmd) {
// 	c = cs.Field.previous()
// 	return
// }

func (cs *Checks) focus() tea.Cmd { return nil }
func (cs *Checks) blur()          {}

// func (cs *Checks) depend()            {}
func (cs *Checks) set(string, string) {}

func (cs *Checks) value() (v string) {
	values := []string{}
	for _, i := range cs.Selections {
		it := cs.Items[i]
		values = append(values, it.Value)
	}
	v = strings.Join(values, "\n")
	return
}

// func (cs *Checks) validity() string { return "" }

func (cs *Checks) up() {
	cs.SelectionIndex--
	if cs.SelectionIndex == -1 {
		cs.SelectionIndex++
	}
}

func (cs *Checks) down() {
	cs.SelectionIndex++
	if cs.SelectionIndex == len(cs.Items) {
		cs.SelectionIndex--
	}
}

// func (cs *Checks) shown() []string { return nil }
