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

type ShapesShow struct {
	Body   *Body
	WID    string
	FID    string
	ID     string
	Reader *models.ShapeReader
}

func newShapesShow(b *Body, wid string, fid string, id string) (ss *ShapesShow) {
	ss = &ShapesShow{Body: b, WID: wid, FID: fid, ID: id}
	return
}

func (ss *ShapesShow) Init() (c tea.Cmd) {
	ss.setReader()
	return
}

func (ss *ShapesShow) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = ss
	return
}

func (ss *ShapesShow) View() (v string) {
	style := lipgloss.NewStyle().Padding(1)
	boldStyle := lipgloss.NewStyle().Bold(true)
	v = lipgloss.JoinVertical(lipgloss.Left,
		style.Render(fmt.Sprintf("%s %s", boldStyle.Render(ss.Reader.Name), ss.Reader.About)),
		string(utils.YamlMarshal(ss.Reader)),
	)
	return
}

func (ss *ShapesShow) setSize(w int, h int) {
}

func (ss *ShapesShow) setReader() {
	result := ss.Body.call(
		controllers.ShapesRead,
		&controllers.ShapesReadParams{
			Shape: ss.ID,
		},
		"..",
	)
	if result != nil {
		ss.Reader = result.Payload.(*models.ShapeReader)
	}
}

func (ss *ShapesShow) focusChain() (fc []tuisupport.Focuser) {
	return
}
