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

type FramesShow struct {
	Body        *Body
	WID         string
	ID          string
	Reader      *models.FrameReader
	ShapeItems  []*controllers.ShapesIndexItemResult
	ShapesTable *Table
}

func newFramesShow(b *Body, wid string, id string) (fs *FramesShow) {
	fs = &FramesShow{Body: b, WID: wid, ID: id}
	return
}

func (fs *FramesShow) Init() (c tea.Cmd) {
	fs.setReader()
	fs.setShapeItems()
	fs.setShapesTable()
	return
}

func (fs *FramesShow) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = fs
	_, c = fs.ShapesTable.Update(msg)
	return
}

func (fs *FramesShow) View() (v string) {
	style := lipgloss.NewStyle().Padding(1)
	boldStyle := lipgloss.NewStyle().Bold(true)
	v = lipgloss.JoinVertical(lipgloss.Left,
		style.Render(fmt.Sprintf("%s %s", boldStyle.Render(fs.Reader.Name), fs.Reader.About)),
		fs.ShapesTable.View(),
		string(utils.YamlMarshal(fs.Reader)),
	)
	return
}

func (fs *FramesShow) setSize(w int, h int) {
	fs.ShapesTable.setSize(w, h)
}

func (fs *FramesShow) setReader() {
	result := fs.Body.call(
		controllers.FramesRead,
		&controllers.FramesReadParams{
			Frame: fs.ID,
		},
		"..",
	)
	if result != nil {
		fs.Reader = result.Payload.(*models.FrameReader)
	}
}

func (fs *FramesShow) setShapeItems() {
	result := fs.Body.call(
		controllers.ShapesIndex,
		&controllers.ShapesIndexParams{
			Workspace: fs.WID,
			Frame:     fs.ID,
		},
		"..",
	)
	if result != nil {
		fs.ShapeItems = result.Payload.([]*controllers.ShapesIndexItemResult)
	}
}

func (fs *FramesShow) setShapesTable() {
	propeties := []string{"Shape", "About"}
	data := []map[string]string{}
	for _, s := range fs.ShapeItems {
		data = append(data, map[string]string{
			"ID":    fmt.Sprintf("%s.%s.%s", s.Workspace, s.Frame, s.Shape),
			"Shape": s.Shape,
			"About": s.About,
		})
	}
	navigator := func(id string) (p string) {
		p = fmt.Sprintf("/shapes/@%s", id)
		return
	}
	fs.ShapesTable = NewTable("workspace-frames", propeties, data, navigator)
}

func (fs *FramesShow) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{fs.ShapesTable}
	return
}
