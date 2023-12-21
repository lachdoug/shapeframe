package tui

import (
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Back struct {
	TopBar  *TopBar
	IsHover bool
	IsFocus bool
}

func newBack(tb *TopBar) (b *Back) {
	b = &Back{TopBar: tb}
	return
}

func (b *Back) Init() (c tea.Cmd) {
	return
}

func (b *Back) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = b
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if b.IsFocus {
				c = b.enter()
				return
			}
		}
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			if zone.Get("back").InBounds(msg) {
				b.IsHover = true
			} else {
				b.IsHover = false
			}
		case tea.MouseLeft:
			if zone.Get("back").InBounds(msg) {
				c = b.enter()
			}
		}
	}
	return
}

func (b *Back) View() (v string) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("15"))

	if b.IsHover {
		style.Background(lipgloss.Color("0"))
	}
	if b.IsFocus {
		style.Underline(true)
	}

	v = zone.Mark("back", style.Render("🡠 "))
	return
}

func (b *Back) Focus(_ string) (c tea.Cmd) {
	b.IsFocus = true
	return
}

func (b *Back) Blur() {
	b.IsFocus = false
}

func (b *Back) enter() (c tea.Cmd) {
	c = tuisupport.Open("..")
	return
}
