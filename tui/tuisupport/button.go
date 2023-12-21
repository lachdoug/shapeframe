package tuisupport

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Button struct {
	ID        string
	Text      string
	Callback  func() tea.Cmd
	IsHover   bool
	IsFocus   bool
	IsEnabled bool
	Color     string
}

func NewButton(
	id string,
	text string,
	callback func() tea.Cmd,
	color int,
) (btn *Button) {
	btn = &Button{
		ID:        id,
		Text:      text,
		Callback:  callback,
		Color:     fmt.Sprint(color),
		IsEnabled: true,
	}
	return
}

func (btn *Button) Init() tea.Cmd { return nil }

func (btn *Button) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = btn
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if btn.IsFocus {
				c = btn.Callback()
			}
		}
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			if zone.Get(btn.ID).InBounds(msg) {
				btn.IsHover = true
			} else {
				btn.IsHover = false
			}
		case tea.MouseLeft:
			if btn.IsEnabled && btn.IsHover {
				c = btn.Callback()
			}
		}
	}
	return
}

func (btn *Button) View() (v string) {
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		PaddingLeft(1).PaddingRight(1)

	textStyle := lipgloss.NewStyle()

	if btn.IsEnabled && btn.IsHover {
		style = style.
			Background(lipgloss.Color("0")).
			BorderBackground(lipgloss.Color("0"))
	}

	if btn.IsEnabled && btn.IsFocus {
		style = style.BorderForeground(lipgloss.Color("15"))
		textStyle = textStyle.Foreground(lipgloss.Color(btn.Color))
	} else {
		style = style.BorderForeground(lipgloss.Color("8"))
		textStyle = textStyle.Foreground(lipgloss.Color("8"))
	}

	if btn.IsEnabled && (btn.IsHover || btn.IsFocus) {
		textStyle = textStyle.Foreground(lipgloss.Color(btn.Color))
	} else {
		textStyle = textStyle.Foreground(lipgloss.Color("8"))
	}

	v = zone.Mark(btn.ID, style.Render(textStyle.Render(btn.Text)))
	return
}

func (btn *Button) Focus(aspect string) (c tea.Cmd) {
	if btn.IsEnabled {
		btn.IsFocus = true
	} else {
		if aspect == "next" {
			c = func() tea.Msg { return tea.KeyMsg{Type: tea.KeyTab} }
		} else {
			c = func() tea.Msg { return tea.KeyMsg{Type: tea.KeyShiftTab} }
		}
	}
	return
}

func (btn *Button) Blur() {
	btn.IsFocus = false
}

func (btn *Button) Enabled(is bool) {
	btn.IsEnabled = is
}
