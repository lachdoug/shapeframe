package tui

import (
	"fmt"
	"sf/controllers"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ShapesDelete struct {
	Body    *Body
	WID     string
	FID     string
	ID      string
	Cancel  *tuisupport.Button
	Confirm *tuisupport.Button
}

func newShapesDelete(b *Body, wid string, fid string, id string) (sd *ShapesDelete) {
	sd = &ShapesDelete{Body: b, WID: wid, FID: fid, ID: id}
	return
}

func (sd *ShapesDelete) Init() (c tea.Cmd) {
	sd.setCancel()
	sd.setConfirm()
	return
}

func (sd *ShapesDelete) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = sd
	cs := []tea.Cmd{}
	_, c = sd.Confirm.Update(msg)
	cs = append(cs, c)
	_, c = sd.Cancel.Update(msg)
	cs = append(cs, c)
	c = tea.Batch(cs...)
	return
}

func (sd *ShapesDelete) View() (v string) {
	style := lipgloss.NewStyle().Padding(1)
	idStyle := lipgloss.NewStyle().Bold(true)
	id := fmt.Sprintf("%s.%s.%s", sd.WID, sd.FID, sd.ID)
	msg := style.Render(fmt.Sprintf("Delete shape %s?", idStyle.Render(id)))
	v = lipgloss.JoinVertical(lipgloss.Left,
		msg,
		lipgloss.JoinHorizontal(lipgloss.Top,
			sd.Cancel.View(),
			sd.Confirm.View(),
		),
	)
	return
}

func (sd *ShapesDelete) setSize(w int, h int) {}

func (sd *ShapesDelete) setConfirm() {
	sd.Confirm = tuisupport.NewButton(
		fmt.Sprintf("shapes-%s-delete-confirm", sd.ID),
		"Confirm",
		sd.confirm,
		15,
	)
}

func (sd *ShapesDelete) setCancel() {
	sd.Cancel = tuisupport.NewButton(
		fmt.Sprintf("shapes-%s-delete-cancel", sd.ID),
		"Cancel",
		sd.cancel,
		9,
	)
}

func (sd *ShapesDelete) confirm() (c tea.Cmd) {
	result := sd.Body.call(
		controllers.ShapesDelete,
		&controllers.ShapesDeleteParams{
			Shape: sd.ID,
		},
		"../..",
	)
	if result != nil {
		c = Open("../..")
	}
	return
}

func (sd *ShapesDelete) cancel() (c tea.Cmd) {
	c = Open("..")
	return
}

func (sd *ShapesDelete) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{
		sd.Cancel,
		sd.Confirm,
	}
	return
}

func (sd *ShapesDelete) Focus(aspect string) (c tea.Cmd) {
	return
}

func (sd *ShapesDelete) Blur() {}
