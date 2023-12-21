package tuisupport

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type NavLink struct {
	App      Apper
	ID       string
	Text     string
	Path     string
	IsHover  bool
	IsFocus  bool
	IsActive bool
}

func NewNavLink(app Apper, id string, text string, path string) (nlk *NavLink) {
	nlk = &NavLink{App: app, ID: id, Text: text, Path: path}
	return
}

func (nlk *NavLink) Init() (c tea.Cmd) {
	if is, _ := nlk.App.MatchRoute(nlk.Path + ".*"); is {
		nlk.IsActive = true
	}
	return
}

func (nlk *NavLink) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = nlk

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if nlk.IsFocus {
				c = nlk.enter()
				return
			}
		}
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			if zone.Get(nlk.ID).InBounds(msg) {
				nlk.IsHover = true
			} else {
				nlk.IsHover = false
			}
		case tea.MouseLeft:
			if zone.Get(nlk.ID).InBounds(msg) {
				c = nlk.enter()
				return
			}
		}
	}
	return
}

func (nlk *NavLink) View() (v string) {
	style := lipgloss.NewStyle().PaddingRight(1)
	textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("12"))

	if nlk.IsHover {
		textStyle.Background(lipgloss.Color("0"))
	}
	if nlk.IsFocus {
		textStyle.Underline(true)
	}
	if nlk.IsActive {
		textStyle.Foreground(lipgloss.Color("12"))
	} else {
		textStyle.Foreground(lipgloss.Color("15"))
	}

	v = zone.Mark(nlk.ID, style.Render(textStyle.Render(nlk.Text)))
	return
}

func (nlk *NavLink) Focus(_ string) (c tea.Cmd) {
	nlk.IsFocus = true
	return
}

func (nlk *NavLink) Blur() {
	nlk.IsFocus = false
}

func (nlk *NavLink) enter() (c tea.Cmd) {
	c = Open(nlk.Path)
	return
}
