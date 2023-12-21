package tuisupport

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Scroller struct {
	Text            string
	ID              string
	Key             tea.KeyType
	IsFocus         bool
	IsHover         bool
	IsButtonDown    bool
	WhenLastCommand *time.Time
}

type ScrollerTickMsg time.Time

func NewScroller(text string, key tea.KeyType) (s *Scroller) {
	s = &Scroller{Text: text, Key: key}
	return
}

func (s *Scroller) Init() (c tea.Cmd) {
	s.setID()
	return
}

func (s *Scroller) setID() {
	s.ID = fmt.Sprintf("scroller-%s", s.Text)
}

func (s *Scroller) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = s
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if s.IsFocus {
				c = s.command()
				return
			}
		}
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			if zone.Get(s.ID).InBounds(msg) {
				s.IsHover = true
			} else {
				s.IsHover = false
			}
		case tea.MouseRelease:
			s.IsButtonDown = false
		case tea.MouseLeft:
			if s.IsHover {
				s.IsButtonDown = true
				c = s.command()
			}
		}
	case ScrollerTickMsg:
		if s.IsButtonDown {
			c = s.command()
		}
	}
	return
}

func (s *Scroller) View() (v string) {
	style := lipgloss.NewStyle().PaddingRight(1)
	if s.IsFocus {
		style = style.Underline(true)
	} else {
		style = style.Foreground(lipgloss.Color("8"))
	}
	if s.IsHover {
		style = style.Foreground(lipgloss.Color("15"))
	}
	v = zone.Mark(s.ID, style.Render(s.Text))
	return
}

func (s *Scroller) command() (c tea.Cmd) {
	c = tea.Batch(
		func() tea.Msg { return tea.KeyMsg{Type: s.Key} },
		tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg { return ScrollerTickMsg(t) }),
	)
	return
}

func (s *Scroller) Focus(string) (c tea.Cmd) {
	s.IsFocus = true
	return
}

func (s *Scroller) Blur() {
	s.IsFocus = false
}
