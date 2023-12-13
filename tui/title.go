package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Title struct {
	TopBar  *TopBar
	IsHover bool
	IsFocus bool
}

func newTitle(tb *TopBar) (t *Title) {
	t = &Title{TopBar: tb}
	return
}

func (t *Title) Init() (c tea.Cmd) {
	return
}

func (t *Title) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = t
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if t.IsFocus {
				c = t.enter()
				return
			}
		}
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			if zone.Get("title").InBounds(msg) {
				t.IsHover = true
			} else {
				t.IsHover = false
			}
		case tea.MouseLeft:
			if zone.Get("title").InBounds(msg) {
				c = t.enter()
				return
			}
		}
	}
	return
}

func (t *Title) View() (v string) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("12"))

	if t.IsHover {
		style.Background(lipgloss.Color("0"))
	}
	if t.IsFocus {
		style.Underline(true)
	}

	v = zone.Mark("title", style.Render("Shapeframe"))
	return
}

func (t *Title) Focus(_ string) (c tea.Cmd) {
	t.IsFocus = true
	return
}

func (t *Title) Blur() {
	t.IsFocus = false
}

func (t *Title) enter() (c tea.Cmd) {
	c = Open("/")
	return
}
