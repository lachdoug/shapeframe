package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Link struct {
	App      *App
	ID       string
	Text     string
	Path     string
	IsHover  bool
	IsFocus  bool
	IsActive bool
}

func newLink(a *App, id string, t string, p string) (lk *Link) {
	lk = &Link{App: a, ID: id, Text: t, Path: p}
	return
}

func (lk *Link) Init() (c tea.Cmd) {
	if is, _ := lk.App.matchRoute(lk.Path + ".*"); is {
		lk.IsActive = true
	}
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
	textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("12"))

	if lk.IsHover {
		textStyle.Background(lipgloss.Color("0"))
	}
	if lk.IsFocus {
		textStyle.Underline(true)
	}
	if lk.IsActive {
		textStyle.Foreground(lipgloss.Color("12"))
	} else {
		textStyle.Foreground(lipgloss.Color("15"))
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
