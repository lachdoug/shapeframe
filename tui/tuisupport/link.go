package tuisupport

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Link struct {
	ID      string
	Text    string
	Path    string
	IsHover bool
	IsFocus bool
}

func NewLink(id string, text string, path string) (lk *Link) {
	lk = &Link{ID: id, Text: text, Path: path}
	return
}

func (lk *Link) Init() (c tea.Cmd) {
	return
}

func (lk *Link) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = lk

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if lk.IsFocus {
				c = lk.enter()
				return
			}
		}
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			if zone.Get(lk.ID).InBounds(msg) {
				lk.IsHover = true
			} else {
				lk.IsHover = false
			}
		case tea.MouseLeft:
			if zone.Get(lk.ID).InBounds(msg) {
				c = lk.enter()
				return
			}
		}
	}
	return
}

func (lk *Link) View() (v string) {
	style := lipgloss.NewStyle().PaddingRight(1)
	textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("15"))

	if lk.IsHover {
		textStyle.Background(lipgloss.Color("0"))
	}
	if lk.IsFocus {
		textStyle.Underline(true)
	}

	v = zone.Mark(lk.ID, style.Render(textStyle.Render(lk.Text)))
	return
}

func (lk *Link) Focus(_ string) (c tea.Cmd) {
	lk.IsFocus = true
	return
}

func (lk *Link) Blur() {
	lk.IsFocus = false
}

func (lk *Link) enter() (c tea.Cmd) {
	c = Open(lk.Path)
	return
}
