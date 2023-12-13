package tui

import (
	"fmt"
	"path/filepath"
	"regexp"
	"sf/tui/tuisupport"
	"sf/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type App struct {
	Path   string
	TopBar *TopBar
	TopNav *TopNav
	Body   *Body
	Width  int
	Height int
	// IsHover    bool
	// IsFocus    bool
	FocusChain []tuisupport.Focuser
	FocusIndex int
}

func newApp() (a *App) {
	a = &App{}
	return
}

func (a *App) Init() (c tea.Cmd) {
	a.Path = "/"
	a.setComponents()
	c = tea.Batch(
		a.TopBar.Init(),
		a.TopNav.Init(),
		a.Body.Init(),
		// a.focus("first", "next"),
		a.setFocus(),
	)
	return
}

func (a *App) focusOn(target tuisupport.Focuser) (c tea.Cmd) {
	for i, f := range a.FocusChain {
		if f == target {
			a.FocusChain[a.FocusIndex].Blur()
			a.FocusIndex = i
			c = a.Focus("next")
			break
		}
	}
	return
}

func (a *App) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = a
	cs := []tea.Cmd{}
	switch msg := msg.(type) {
	case tuisupport.Navigation:
		c = a.open(msg.Path)
		return
	case tuisupport.TakeFocus:
		c = a.focusOn(msg)
		return
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			c = tea.Quit
			return
		case tea.KeyTab:
			c = a.next()
			return
		case tea.KeyShiftTab:
			c = a.previous()
			return
		case tea.KeyCtrlLeft:
			c = Open("..")
			return

			// case tea.KeyEnter, tea.KeyTab, tea.KeyShiftTab:
			// 	if a.FocusIndex == 0 {
			// 		_, c = a.TopBar.Update(msg)
			// 	} else if a.FocusIndex == 1 {
			// 		_, c = a.TopNav.Update(msg)
			// 	} else {
			// 		_, c = a.Body.Update(msg)
			// 	}
			// 	return
		}
	case tea.WindowSizeMsg:
		a.Width = msg.Width
		a.Height = msg.Height
		a.setSize()
		cs = append(cs, tea.ClearScreen)
	}
	_, c = a.TopBar.Update(msg)
	cs = append(cs, c)
	_, c = a.TopNav.Update(msg)
	cs = append(cs, c)
	_, c = a.Body.Update(msg)
	cs = append(cs, c)
	c = tea.Batch(cs...)
	return
}

func (a *App) View() (v string) {
	v = zone.Scan(
		lipgloss.JoinVertical(lipgloss.Left,
			a.TopBar.View(),
			a.TopNav.View(),
			a.Body.View(),
		),
	)
	return
}

func (a *App) setComponents() {
	a.TopBar = newTopBar(a)
	a.TopNav = newTopNav(a)
	a.Body = newBody(a)
}

func (a *App) setSize() {
	w, h := utils.TerminalSize()
	a.Width = w
	a.Height = h
	a.TopBar.setSize(w, 1)
	a.TopNav.setSize(w, 3)
	a.Body.setSize(w, h-4)
}

func (a *App) open(path string) (c tea.Cmd) {
	a.FocusChain[a.FocusIndex].Blur()
	if path[:1] == "/" {
		a.Path = path
	} else {
		a.Path = filepath.Clean(fmt.Sprintf("%s/%s", a.Path, path))
	}
	c = tea.Batch(
		a.TopNav.Init(),
		a.Body.Init(),
		a.setFocus(),
	)
	a.setSize()
	return
}

func (a *App) matchRoute(route string) (is bool, captures []string) {
	exp := regexp.MustCompile("^" + route + "$")
	match := exp.FindStringSubmatch(a.Path)
	if len(match) == 0 {
		return
	}
	is = true
	captures = match[1:]
	return
}

func (a *App) setFocus() (c tea.Cmd) {
	a.FocusChain = []tuisupport.Focuser{}
	a.FocusChain = append(a.FocusChain, a.TopBar.focusChain()...)
	a.FocusChain = append(a.FocusChain, a.TopNav.focusChain()...)
	a.FocusChain = append(a.FocusChain, a.Body.focusChain()...)
	a.FocusIndex = a.navFocusIndex()
	c = a.Focus("next")
	return
}

func (a *App) navFocusIndex() (i int) {
	i = len(a.TopBar.focusChain()) + len(a.TopNav.focusChain())
	// i equals length of a.FocusChain when body has not focusable components.
	if i == len(a.FocusChain) {
		i = 0
	}
	return
}

func (a *App) Focus(aspect string) (c tea.Cmd) {
	c = a.FocusChain[a.FocusIndex].Focus(aspect)
	return
}

func (a *App) next() (c tea.Cmd) {
	a.FocusChain[a.FocusIndex].Blur()
	a.FocusIndex++
	if a.FocusIndex == len(a.FocusChain) {
		a.FocusIndex = 0
	}
	c = a.Focus("next")
	return
}

func (a *App) previous() (c tea.Cmd) {
	a.FocusChain[a.FocusIndex].Blur()
	a.FocusIndex--
	if a.FocusIndex == -1 {
		a.FocusIndex = len(a.FocusChain) - 1
	}
	c = a.Focus("previous")
	return
}
