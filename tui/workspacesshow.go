package tui

import (
	"fmt"
	"sf/controllers"
	"sf/models"
	"sf/tui/tuisupport"
	"sf/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type WorkspacesShow struct {
	Body        *Body
	ID          string
	Reader      *models.WorkspaceReader
	FrameItems  []*controllers.FramesIndexItemResult
	FramesTable *Table
}

func newWorkspacesShow(b *Body, id string) (ws *WorkspacesShow) {
	ws = &WorkspacesShow{Body: b, ID: id}
	return
}

func (ws *WorkspacesShow) Init() (c tea.Cmd) {
	ws.setReader()
	ws.setFrameItems()
	ws.setFramesTable()
	return
}

func (ws *WorkspacesShow) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = ws
	_, c = ws.FramesTable.Update(msg)
	return
}

func (ws *WorkspacesShow) View() (v string) {
	style := lipgloss.NewStyle().Padding(1)
	boldStyle := lipgloss.NewStyle().Bold(true)
	v = lipgloss.JoinVertical(lipgloss.Left,
		style.Render(fmt.Sprintf("%s %s", boldStyle.Render(ws.Reader.Name), ws.Reader.About)),
		ws.FramesTable.View(),
		string(utils.YamlMarshal(ws.Reader)),
	)
	return
}

func (ws *WorkspacesShow) setSize(w int, h int) {
	ws.FramesTable.setSize(w, h)
}

func (ws *WorkspacesShow) setReader() {
	result := ws.Body.call(
		controllers.WorkspacesRead,
		&controllers.WorkspacesReadParams{
			Workspace: ws.ID,
		},
		"..",
	)
	if result != nil {
		ws.Reader = result.Payload.(*models.WorkspaceReader)
	}
}

func (ws *WorkspacesShow) setFrameItems() {
	result := ws.Body.call(
		controllers.FramesIndex,
		&controllers.FramesIndexParams{
			Workspace: ws.ID,
		},
		"..",
	)
	if result != nil {
		ws.FrameItems = result.Payload.([]*controllers.FramesIndexItemResult)
	}
}

func (ws *WorkspacesShow) setFramesTable() {
	propeties := []string{"Frame", "About"}
	data := []map[string]string{}
	for _, f := range ws.FrameItems {
		data = append(data, map[string]string{
			"ID":    fmt.Sprintf("%s.%s", f.Workspace, f.Frame),
			"Frame": f.Frame,
			"About": f.About,
		})
	}
	navigator := func(id string) (p string) {
		p = fmt.Sprintf("/frames/@%s", id)
		return
	}
	ws.FramesTable = NewTable("workspace-frames", propeties, data, navigator)
}

func (ws *WorkspacesShow) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{ws.FramesTable}
	return
}
