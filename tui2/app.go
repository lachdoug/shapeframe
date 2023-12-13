package tui2

import (
	"fmt"
	"sf/controllers"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type App struct {
	Workspace  string
	Frame      string
	Shape      string
	Action     string
	Errors     []string
	Views      []tea.Model
	Clickables []string
}

func NewApp() (a *App) {
	a = &App{}
	return
}

func (a *App) Init() (c tea.Cmd) {
	a.setContext()
	a.setAction()
	c = a.setViews()
	return
}

func (a *App) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = a
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			c = tea.Quit
			return
		}
		// case tea.WindowSizeMsg:
		// 	fm.resize()
		// 	return
		// case error:
		// 	a.Error = msg.Error()

	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseLeft:
			if zone.Get("back").InBounds(msg) {
				a.back()
			}
			if zone.Get("edit").InBounds(msg) {
				a.Action = "edit"
			}
			a.setViews()
		}
	}
	// c = fm.updateControls(msg)
	// fm.depend()
	// return

	return
}

func (a *App) View() (v string) {
	v = zone.Scan(
		lipgloss.JoinVertical(
			lipgloss.Left,
			a.viewTopbar(),
			time.Now().String(),
			a.viewBody(),
		),
	)
	return
}

func (a *App) viewTopbar() (v string) {
	cx := []string{a.Workspace}
	if a.Frame != "" {
		cx = append(cx, a.Frame)
	}
	if a.Shape != "" {
		cx = append(cx, a.Shape)
	}
	if a.Action != "show" {
		cx = append(cx, a.Action)
	}
	titleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	errorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	title := titleStyle.Render("Shapeframe")
	lines := []string{fmt.Sprintf("%s /%s", title, strings.Join(cx, "/"))}
	if len(a.Errors) > 0 {
		for _, errmsg := range a.Errors {
			line := errorStyle.Render(fmt.Sprintf("Error: %s", errmsg))
			lines = append(lines, line)
		}
	}
	v = lipgloss.JoinVertical(lipgloss.Left, lines...)
	return
}

func (a *App) viewBody() (v string) {
	lines := []string{}
	for _, vm := range a.Views {
		lines = append(lines, vm.View())
	}
	v = lipgloss.JoinVertical(lipgloss.Left, lines...)
	return
}

func (a *App) setAction() {
	a.Action = "show"
}

func (a *App) setContext() {
	var err error
	var result *controllers.Result
	if result, err = controllers.ContextsRead(nil); err != nil {
		a.Errors = append(a.Errors, err.Error())
		return
	}
	r := result.Payload.(*controllers.ContextsReadResult)
	a.Workspace = r.Workspace
	a.Frame = r.Frame
	a.Shape = r.Shape
}

func (a *App) setViews() (c tea.Cmd) {
	a.Views = []tea.Model{}
	if a.Shape != "" {
		a.Views = append(a.Views, NewShape(a))
	}
	cs := []tea.Cmd{}
	for _, vw := range a.Views {
		cs = append(cs, vw.Init())
	}
	c = tea.Batch(cs...)
	return
}

func (a *App) back() {
	if a.Action != "" {
		a.Action = "show"
	} else if a.Shape != "" {
		a.Shape = ""
	} else if a.Frame != "" {
		a.Frame = ""
	} else if a.Workspace != "" {
		a.Workspace = ""
	}
}

func (a *App) registerClickable(id string) {
	a.Clickables = append(a.Clickables, id)
}
