package tui

import (
	"sf/tui/tuisupport"
	"sf/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TopBar struct {
	App        *App
	Title      *Title
	Back       *Back
	Width      int
	IsFocus    bool
	FocusIndex int
}

func newTopBar(a *App) (tb *TopBar) {
	tb = &TopBar{App: a}
	return
}

func (tb *TopBar) Init() (c tea.Cmd) {
	tb.setComponents()
	// tb.focus()
	return
}

func (tb *TopBar) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = tb
	cs := []tea.Cmd{}
	_, c = tb.Title.Update(msg)
	cs = append(cs, c)
	_, c = tb.Back.Update(msg)
	cs = append(cs, c)
	c = tea.Batch(cs...)
	return
}

func (tb *TopBar) View() (v string) {
	style := lipgloss.NewStyle().Width(tb.Width)
	leftpadStyle := lipgloss.NewStyle().PaddingLeft(1)
	path := utils.FixedLengthString(tb.App.Path, tb.Width-14)
	v = style.Render(lipgloss.JoinHorizontal(lipgloss.Top,
		tb.Title.View(),
		leftpadStyle.Render(tb.Back.View()),
		leftpadStyle.Render(path),
	))
	return
}

func (tb *TopBar) setSize(w int, h int) {
	tb.Width = w
}

func (tb *TopBar) setComponents() {
	tb.Title = newTitle(tb)
	tb.Back = newBack(tb)
}

func (tb *TopBar) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{
		tb.Title,
		tb.Back,
	}
	return
}

// func (tb *TopBar) focus(on ...string) (c tea.Cmd) {
// 	if slices.Contains(on, "first") {
// 		tb.FocusIndex = 0
// 	} else if slices.Contains(on, "last") {
// 		tb.FocusIndex = 1
// 	}
// 	tb.IsFocus = true
// 	if tb.FocusIndex == 0 {
// 		c = tb.Title.focus(on...)
// 	} else if tb.FocusIndex == 1 {
// 		c = tb.Back.focus(on...)
// 	}
// 	return
// }

// func (tb *TopBar) blur() {
// 	if tb.FocusIndex == 0 {
// 		tb.Title.blur()
// 	} else if tb.FocusIndex == 1 {
// 		tb.Back.blur()
// 	}
// }

// func (tb *TopBar) next() (c tea.Cmd) {
// 	// tb.blur()
// 	tb.FocusIndex++
// 	if tb.FocusIndex == 2 {
// 		tb.FocusIndex = 1
// 		c = tb.App.next()
// 	} else {
// 		c = tb.focus("next")
// 	}
// 	return
// }

// func (tb *TopBar) previous() (c tea.Cmd) {
// 	// tb.blur()
// 	tb.FocusIndex--
// 	if tb.FocusIndex == -1 {
// 		tb.FocusIndex = 0
// 		c = tb.App.previous()
// 	} else {
// 		c = tb.focus("previous")
// 	}
// 	return
// }
