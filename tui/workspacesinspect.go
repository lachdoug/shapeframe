package tui

import (
	"sf/controllers"
	"sf/models"
	"sf/tui/tuisupport"
	"sf/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

type WorkspacesInspect struct {
	Body      *Body
	ID        string
	Inspector *models.WorkspaceInspector
	Table     table.Model
}

func newWorkspacesInspect(b *Body, id string) (wi *WorkspacesInspect) {
	wi = &WorkspacesInspect{Body: b, ID: id}
	return
}

func (wi *WorkspacesInspect) Init() (c tea.Cmd) {
	wi.setWorkspace()
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

func (wi *WorkspacesInspect) setWorkspace() {
	result := wi.Body.call(
		controllers.WorkspaceInspectsRead,
		&controllers.WorkspaceInspectsReadParams{
			Workspace: wi.ID,
		},
		"..",
	)
	if result != nil {
		wi.Inspector = result.Payload.(*models.WorkspaceInspector)
	}
}

func (wi *WorkspacesInspect) focusChain() (fc []tuisupport.Focuser) {
	return
}

func (wi *WorkspacesInspect) Focus(aspect string) (c tea.Cmd) {
	return
}

func (wi *WorkspacesInspect) Blur() {}
