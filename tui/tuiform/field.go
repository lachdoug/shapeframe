package tuiform

import (
	"fmt"
	"sf/app/validations"
	"sf/models"
	"sf/tui/tuisupport"
	"sf/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Field struct {
	Form      *Form
	ID        string
	Model     *models.FormComponent
	Input     Inputter
	Width     int
	Validity  string
	IsHover   bool
	IsFocus   bool
	IsVisible bool
}

func NewField(fm *Form, fmcm *models.FormComponent) (f *Field) {
	f = &Field{Form: fm, Model: fmcm}
	return
}

func (f *Field) answer() (an string) {
	an = f.Form.Answers[f.Model.Key]
	if an == "" {
		an = f.Model.Default
	}
	return
}

func (f *Field) setWidth(w int) {
	f.Width = (w * f.Model.Width / 12)
}

func (f *Field) setInput() {
	switch f.Model.As {
	case "select":
		f.Input = NewSelect(f)
	case "radios":
		f.Input = NewRadios(f)
	case "selects":
		f.Input = NewSelects(f)
	case "checks":
		f.Input = NewChecks(f)
	default:
		f.Input = NewInput(f)
	}
}

func (f *Field) Init() (c tea.Cmd) {
	f.setID()
	f.setInput()
	f.Input.Init()
	f.Input.setAnswer()
	return
}

func (f *Field) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = f
	switch msg := msg.(type) {
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			if zone.Get(f.ID).InBounds(msg) {
				f.IsHover = true
			} else {
				f.IsHover = false
			}
		case tea.MouseLeft:
			if f.IsHover && !f.IsFocus {
				c = f.takeFocus()
				return
			}
		}
	}
	_, c = f.Input.Update(msg)
	return
}

func (f *Field) View() (v string) {
	if !f.IsVisible {
		return
	}
	style := lipgloss.NewStyle()
	inputStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Width(f.Width - 2)
	validityStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("13"))

	if f.IsFocus {
		inputStyle = inputStyle.BorderForeground(lipgloss.Color("15"))
	} else {
		inputStyle = inputStyle.BorderForeground(lipgloss.Color("8"))
	}

	if f.IsHover {
		style = style.Background(lipgloss.Color("0"))
		inputStyle = inputStyle.BorderBackground(lipgloss.Color("0"))
	}

	label := utils.FixedLengthString(f.Model.Label, f.Width)
	validity := utils.FixedLengthString(f.Validity, f.Width)

	lines := []string{
		label,
		inputStyle.Render(f.Input.View()),
		validityStyle.Render(validity),
	}
	v = zone.Mark(
		f.ID,
		style.Render(lipgloss.JoinVertical(lipgloss.Left, lines...)),
	)

	return
}

func (f *Field) setID() {
	f.ID = fmt.Sprintf("%s-%s-field", f.Form.ID, f.Model.Key)
}

func (f *Field) enter() (c tea.Cmd) {
	if f.Validity == "" {
		c = func() tea.Msg { return tea.KeyMsg{Type: tea.KeyTab} }
	}
	return
}

func (f *Field) FocusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{f}
	return
}

func (f *Field) isFocus() (is bool) {
	is = f.IsFocus
	return
}

func (f *Field) takeFocus() (c tea.Cmd) {
	c = tuisupport.TakeFocusCommand(f)
	return
}

func (f *Field) Focus(aspect string) (c tea.Cmd) {
	if !f.IsVisible {
		if aspect == "next" {
			c = func() tea.Msg { return tea.KeyMsg{Type: tea.KeyTab} }
		} else {
			c = func() tea.Msg { return tea.KeyMsg{Type: tea.KeyShiftTab} }
		}
		return
	}
	f.IsFocus = true
	c = f.Input.focus(aspect)
	return
}

func (f *Field) Blur() {
	f.IsFocus = false
	f.Input.blur()
}

func (f *Field) depend() {
	f.IsVisible = f.Form.isDependMatch(f.Model.Depend)
}

func (f *Field) setAnswer(value string) {
	f.Form.setAnswer(f.Model.Key, value)
}

func (f *Field) validity() (vy string) {
	v := f.Input.value()
	vn := validations.NewValidation()
	vy = ""
	f.Model.ValueValidation(v, vn)
	if vn.IsInvalid() {
		vy = vn.Failures[0].Message
	}
	f.Validity = vy
	return
}

func (f *Field) shown() (ks []string) {
	if f.IsVisible {
		ks = []string{f.Model.Key}
	}
	return
}
