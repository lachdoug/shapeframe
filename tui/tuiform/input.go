package tuiform

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Input struct {
	Form      *Form
	Field     *Field
	TextInput textinput.Model
	Width     int
}

func NewInput(
	fm *Form,
	f *Field,
) (i *Input) {
	i = &Input{Form: fm, Field: f}
	return
}

func (i *Input) setWidth() {
	i.Width = i.Field.width()
}

func (i *Input) setTextInput() {
	ti := textinput.New()
	ti.Width = i.Width
	// ti.Placeholder = i.Field.ComponentModel.Placeholder
	ti.Prompt = ""
	ti.SetValue(i.Form.Answers[i.Field.ComponentModel.Key])
	i.TextInput = ti
}

func (i *Input) Init() (c tea.Cmd) {
	i.setWidth()
	i.setTextInput()
	return
}

func (i *Input) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			c = i.enter()
			return
			// case tea.KeyTab:
			// 	c = i.next()
			// 	return
			// case tea.KeyShiftTab:
			// 	c = i.previous()
			// 	return
		}
	case error:
		panic(msg)
	}
	i.TextInput, c = i.TextInput.Update(msg)
	if i.validity() == "" {
		i.answer()
	}
	return
}

func (i *Input) View() (v string) {
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Width(i.Width - 2)
	if i.Field.IsFocus {
		style.
			BorderForeground(lipgloss.Color("15"))
	} else {
		style.
			Foreground(lipgloss.Color("7")).
			BorderForeground(lipgloss.Color("8"))
	}

	v = style.Render(i.TextInput.View())
	return
}

func (i *Input) answer() {
	i.Field.set(i.Field.ComponentModel.Key, i.value())
}

func (i *Input) width() (w int) {
	w = i.Width
	return
}

func (i *Input) resize() {
	i.setWidth()
}

func (i *Input) enter() (c tea.Cmd) {
	c = i.Field.enter()
	return
}

// func (i *Input) next() (c tea.Cmd) {
// 	c = i.Field.next()
// 	return
// }

// func (i *Input) previous() (c tea.Cmd) {
// 	c = i.Field.previous()
// 	return
// }

func (i *Input) focus() (c tea.Cmd) {
	i.TextInput.Focus()
	c = textinput.Blink
	return
}

func (i *Input) blur() {
	i.TextInput.Blur()
}

// func (i *Input) depend()            {}
func (i *Input) set(string, string) {}

func (i *Input) value() (v string) {
	v = i.TextInput.Value()
	return
}

func (i *Input) validity() string { return "" }

// func (i *Input) shown() []string { return nil }
