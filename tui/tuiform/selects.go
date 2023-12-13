package tuiform

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/slices"
)

type Selects struct {
	Form           *Form
	Field          *Field
	Items          []*Option
	SelectionIndex int
	Selections     []int
	Width          int
}

func NewSelects(fm *Form, f *Field) (ss *Selects) {
	ss = &Selects{Form: fm, Field: f}
	return
}

func (ss *Selects) setWidth() {
	ss.Width = ss.Field.Width
}

func (ss *Selects) setItems() {
	items := []*Option{}
	for _, opt := range ss.Field.ComponentModel.Options {
		items = append(items, NewOption(opt))
	}
	ss.Items = items
}

func (ss *Selects) setSelections() {
	ans := strings.Split(ss.Form.Answers[ss.Field.ComponentModel.Key], "\n")
	for i, item := range ss.Field.ComponentModel.Options {
		if slices.Contains(ans, item.Value) {
			ss.Selections = append(ss.Selections, i)
		}
	}
}

func (ss *Selects) Init() (c tea.Cmd) {
	ss.setWidth()
	ss.setItems()
	ss.setSelections()
	return
}

func (ss *Selects) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = ss
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			c = ss.enter()
			return
		case tea.KeySpace:
			ss.selectItem()
			ss.answer()
			return
		// case tea.KeyTab:
		// 	c = ss.next()
		// 	return
		// case tea.KeyShiftTab:
		// 	c = ss.previous()
		// 	return
		case tea.KeyDown:
			ss.down()
			return
		case tea.KeyUp:
			ss.up()
			return
		}
	case error:
		panic(msg)
	}
	return
}

func (ss *Selects) View() (v string) {
	borderStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Width(ss.Width - 2)
	if ss.Field.IsFocus {
		borderStyle.BorderForeground(lipgloss.Color("15"))
	} else {
		borderStyle.BorderForeground(lipgloss.Color("8"))
	}

	v = borderStyle.Render(ss.body())
	return
}

func (ss *Selects) body() (f string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("8"))

	if !ss.Field.IsFocus {
		labels := []string{}
		for _, it := range ss.selectedItems() {
			labels = append(labels, it.View(false, ""))
		}
		f = strings.Join(labels, ", ")
		return
	}
	lines := []string{}
	for i, it := range ss.selectedItems() {
		if i == ss.SelectionIndex {
			lines = append(lines, it.View(true, "⮞"))
		} else {
			lines = append(lines, it.View(false, " "))
		}
	}
	l := len(lines)
	lines = append(lines, style.Render(strings.Repeat("─", ss.Width-2)))
	for i, it := range ss.unselectedItems() {
		if i == ss.SelectionIndex-l {
			lines = append(lines, it.View(true, "⮞"))
		} else {
			lines = append(lines, it.View(false, " "))
		}
	}
	f = strings.Join(lines, "\n")
	return
}

func (ss *Selects) selectedItems() (sels []*Option) {
	sels = []*Option{}
	for i, it := range ss.Items {
		if slices.Contains(ss.Selections, i) {
			sels = append(sels, it)
		}
	}
	return
}

func (ss *Selects) unselectedItems() (unsels []*Option) {
	unsels = []*Option{}
	for i, it := range ss.Items {
		if !slices.Contains(ss.Selections, i) {
			unsels = append(unsels, it)
		}
	}
	return
}

func (ss *Selects) selectItem() {
	selected := ss.selectedItems()
	unselected := ss.unselectedItems()
	if ss.SelectionIndex < len(selected) {
		ss.Selections = slices.Delete(ss.Selections, ss.SelectionIndex, ss.SelectionIndex+1)
	} else {
		uit := unselected[ss.SelectionIndex-len(selected)]
		for i, it := range ss.Items {
			if uit.Value == it.Value {
				ss.Selections = append(ss.Selections, i)
			}
		}
		if ss.isUnselectedEmpty() {
			ss.SelectionIndex = 0
		} else if ss.isUnselectedFocusNotLast() {
			ss.SelectionIndex++
		}
	}
}

func (ss *Selects) isUnselectedEmpty() (is bool) {
	is = len(ss.Selections) == len(ss.Items)
	return
}

func (ss *Selects) isUnselectedFocusNotLast() (is bool) {
	is = len(ss.Items) > ss.SelectionIndex+1 && len(ss.Selections) <= ss.SelectionIndex+1
	return
}

func (ss *Selects) answer() {
	ss.Field.set(ss.Field.ComponentModel.Key, ss.value())
}

func (ss *Selects) width() (w int) {
	w = ss.Width
	return
}

func (ss *Selects) resize() {
	ss.setWidth()
}

func (ss *Selects) enter() (c tea.Cmd) {
	c = ss.Field.enter()
	return
}

// func (ss *Selects) next() (c tea.Cmd) {
// 	c = ss.Field.next()
// 	return
// }

// func (ss *Selects) previous() (c tea.Cmd) {
// 	c = ss.Field.previous()
// 	return
// }

func (ss *Selects) focus() tea.Cmd { return nil }
func (ss *Selects) blur()          {}

// func (ss *Selects) depend()            {}
func (ss *Selects) set(string, string) {}

func (ss *Selects) value() (v string) {
	values := []string{}
	for _, it := range ss.selectedItems() {
		values = append(values, it.Value)
	}
	v = strings.Join(values, "\n")
	return
}

// func (ss *Selects) validity() string { return "" }

func (ss *Selects) up() {
	ss.SelectionIndex--
	if ss.SelectionIndex == -1 {
		ss.SelectionIndex++
	}
}

func (ss *Selects) down() {
	ss.SelectionIndex++
	if ss.SelectionIndex == len(ss.Items) {
		ss.SelectionIndex--
	}
}

// func (ss *Selects) shown() []string { return nil }
