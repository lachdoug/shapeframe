package tui

import (
	"sf/controllers"
	"sf/models"
	"sf/tui/tuisupport"
	"sf/utils"

	tea "github.com/charmbracelet/bubbletea"
)

type WorkspacesInspect struct {
	Body      *Body
	Inspector *models.WorkspaceInspector
}

func newWorkspacesInspect(b *Body) (wi *WorkspacesInspect) {
	wi = &WorkspacesInspect{Body: b}
	return
}

func (wi *WorkspacesInspect) Init() (c tea.Cmd) {
	c = wi.setInspector()
	return
}

func (wi *WorkspacesInspect) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = wi
	return
}

func (wi *WorkspacesInspect) View() (v string) {
	v = string(utils.YamlMarshal(wi.Inspector))
	return
}

func (wi *WorkspacesInspect) setSize(w int, h int) {
}

func (wi *WorkspacesInspect) setInspector() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = wi.Body.App.call(
		controllers.WorkspaceInspectsRead,
		nil,
		tuisupport.Open(".."),
	)
	if result != nil {
		wi.Inspector = result.Payload.(*models.WorkspaceInspector)
	}
	return
}

func (wi *WorkspacesInspect) focusChain() (fc []tuisupport.Focuser) {
	return
}

func (wi *WorkspacesInspect) isFocus() (is bool) {
	return
}
