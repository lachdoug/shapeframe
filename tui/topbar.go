package tui

import (
	"sf/controllers"
	"sf/models"
	"sf/tui/tuisupport"
	"sf/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TopBar struct {
	App    *App
	Title  *Title
	Back   *Back
	Reader *models.WorkspaceReader
	Width  int
}

func newTopBar(a *App) (tb *TopBar) {
	tb = &TopBar{App: a}
	return
}

func (tb *TopBar) Init() (c tea.Cmd) {
	tb.setReader()
	tb.setComponents()
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
	style := lipgloss.NewStyle().Width(tb.Width).Height(2)
	leftpadStyle := lipgloss.NewStyle().PaddingLeft(1)
	nameStyle := leftpadStyle.Copy().Bold(true)
	about := utils.FixedLengthString(tb.Reader.About, tb.Width-len(tb.Reader.Name)-12)
	path := utils.FixedLengthString(tb.App.Path, tb.Width-3)
	v = style.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			lipgloss.JoinHorizontal(lipgloss.Top,
				tb.Title.View(),
				nameStyle.Render(tb.Reader.Name),
				leftpadStyle.Render(about),
			),
			lipgloss.JoinHorizontal(lipgloss.Top,
				tb.Back.View(),
				leftpadStyle.Render(path),
			),
		),
	)
	return
}

func (tb *TopBar) setSize(w int) {
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

func (tb *TopBar) setReader() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = tb.App.call(
		controllers.WorkspacesRead,
		nil,
		func() tea.Msg { return tuisupport.Open(".") },
	)
	if result != nil {
		tb.Reader = result.Payload.(*models.WorkspaceReader)
	}
	return
}
