package tuiform

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input struct {
	Field     *Field
	TextInput textinput.Model
}

func NewInput(
	f *Field,
) (i *Input) {
	i = &Input{Field: f}
	return
}

func (i *Input) setTextInput() {
	ti := textinput.New()
	// ti.Placeholder = i.Field.Model.Placeholder
	ti.Prompt = ""
	ti.SetValue(i.Field.answer())
	i.TextInput = ti
}

func (i *Input) Init() (c tea.Cmd) {
	i.setTextInput()
	return
}

func (i *Input) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	if i.Field.IsFocus {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter:
				c = i.Field.enter()
				return
			default:
				if i.validity() == "" {
					i.setAnswer()
				}
			}
		}
		i.TextInput, c = i.TextInput.Update(msg)
	}
	return
}

func (i *Input) View() (v string) {
	i.TextInput.Width = i.Field.Width - 2
	v = i.TextInput.View()
	return
}

func (i *Input) setAnswer() {
	i.Field.setAnswer(i.value())
}

func (i *Input) focus(string) (c tea.Cmd) {
	i.TextInput.Focus()
	c = textinput.Blink
	return
}

func (i *Input) blur() {
	i.TextInput.Blur()
}

func (i *Input) value() (v string) {
	v = i.TextInput.Value()
	return
}

func (i *Input) validity() string { return "" }
