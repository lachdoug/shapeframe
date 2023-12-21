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
	FID     string
	ID      string
	Cancel  *tuisupport.Button
	Confirm *tuisupport.Button
}

func newShapesDelete(b *Body, fid string, id string) (sd *ShapesDelete) {
	sd = &ShapesDelete{Body: b, FID: fid, ID: id}
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
	id := fmt.Sprintf("%s.%s", sd.FID, sd.ID)
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
	result := &controllers.Result{}
	result, c = sd.Body.App.call(
		controllers.ShapesDelete,
		&controllers.ShapesDeleteParams{
			Shape: sd.ID,
		},
		tuisupport.Open(".."),
	)
	if result != nil {
		c = tuisupport.Open("../..")
	}
	return
}

func (sd *ShapesDelete) cancel() (c tea.Cmd) {
	c = tuisupport.Open("..")
	return
}

func (sd *ShapesDelete) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{
		sd.Cancel,
		sd.Confirm,
	}
	return
}

func (sd *ShapesDelete) isFocus() (is bool) {
	is = sd.Cancel.IsFocus || sd.Confirm.IsFocus
	return
}

func (sd *ShapesDelete) Focus(aspect string) (c tea.Cmd) {
	return
}

func (sd *ShapesDelete) Blur() {}
