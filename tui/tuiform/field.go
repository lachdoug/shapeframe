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
	Form           *Form
	Parent         Parenter
	ID             string
	ComponentModel *models.FormComponent
	Input          Inputter
	Width          int
	Validity       string
	IsHover        bool
	IsFocus        bool
	IsVisible      bool
}

func NewField(fm *Form, pt Parenter, fmcm *models.FormComponent) (f *Field) {
	f = &Field{Form: fm, Parent: pt, ComponentModel: fmcm}
	return
}

func (f *Field) setWidth() {
	f.Width = f.Parent.width() * f.ComponentModel.Width / 12
}

func (f *Field) setInput() {
	switch f.ComponentModel.As {
	case "select":
		f.Input = NewSelect(f.Form, f)
	case "radios":
		f.Input = NewRadios(f.Form, f)
	case "selects":
		f.Input = NewSelects(f.Form, f)
	case "checks":
		f.Input = NewChecks(f.Form, f)
	default:
		f.Input = NewInput(f.Form, f)
	}
}

func (f *Field) Init() (c tea.Cmd) {
	f.setID()
	f.setWidth()
	f.setInput()
	f.Input.Init()
	return
}

func (f *Field) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
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
			if f.IsHover {
				c = f.takeFocus()
			}
		}
	}
	if f.IsFocus {
		m, c = f.Input.Update(msg)
	}
	return
}

func (f *Field) View() (v string) {
	style := lipgloss.NewStyle()
	validityStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("13"))

	if !f.IsVisible {
		return
	}

	if f.IsHover {
		style = style.Background(lipgloss.Color("0"))
	}

	validity := utils.FixedLengthString(f.Validity, f.width())

	lines := []string{
		f.ComponentModel.Label,
		f.Input.View(),
		validityStyle.Render(validity),
	}
	v = zone.Mark(
		f.ID,
		style.Render(lipgloss.JoinVertical(lipgloss.Top, lines...)),
	)

	return
}

func (f *Field) setID() {
	f.ID = fmt.Sprintf("%s-%s-input", f.Form.ID, f.ComponentModel.Key)
}

func (f *Field) width() (w int) {
	w = f.Width
	return
}

func (f *Field) resize() {
	f.setWidth()
	f.Input.resize()
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

func (f *Field) takeFocus() (c tea.Cmd) {
	c = tuisupport.TakeFocusCommand(f)
	return
}

// func (f *Field) next() (c tea.Cmd) {
// 	c = f.Parent.next()
// 	return
// }

// func (f *Field) previous() (c tea.Cmd) {
// 	c = f.Parent.previous()
// 	return
// }

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
	c = f.Input.focus()
	return
}

func (f *Field) Blur() {
	f.IsFocus = false
	f.Input.blur()
}

func (f *Field) depend() {
	f.IsVisible = f.Form.isDependMatch(f.ComponentModel.Depend)
}

func (f *Field) set(key string, value string) {
	f.Parent.set(key, value)
}

// func (f *Field) value() string { return "" }

func (f *Field) validity() (vy string) {
	v := f.Input.value()
	vn := validations.NewValidation()
	vy = ""
	f.ComponentModel.ValueValidation(v, vn)
	if vn.IsInvalid() {
		vy = vn.Failures[0].Message
	}
	f.Validity = vy
	return
}

func (f *Field) shown() (ks []string) {
	if f.IsVisible {
		ks = []string{f.ComponentModel.Key}
	}
	return
}
