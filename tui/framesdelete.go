package tui

import (
	"fmt"
	"sf/controllers"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FramesDelete struct {
	Body    *Body
	ID      string
	Cancel  *tuisupport.Button
	Confirm *tuisupport.Button
}

func newFramesDelete(b *Body, id string) (fd *FramesDelete) {
	fd = &FramesDelete{Body: b, ID: id}
	return
}

func (fd *FramesDelete) Init() (c tea.Cmd) {
	fd.setCancel()
	fd.setConfirm()
	return
}

func (fd *FramesDelete) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = fd
	cs := []tea.Cmd{}
	_, c = fd.Confirm.Update(msg)
	cs = append(cs, c)
	_, c = fd.Cancel.Update(msg)
	cs = append(cs, c)
	c = tea.Batch(cs...)
	return
}

func (fd *FramesDelete) View() (v string) {
	style := lipgloss.NewStyle().Padding(1)
	idStyle := lipgloss.NewStyle().Bold(true)
	msg := style.Render(fmt.Sprintf("Delete frame %s?", idStyle.Render(fd.ID)))
	v = lipgloss.JoinVertical(lipgloss.Left,
		msg,
		lipgloss.JoinHorizontal(lipgloss.Top,
			fd.Cancel.View(),
			fd.Confirm.View(),
		),
	)
	return
}

func (fd *FramesDelete) setSize(w int, h int) {}

func (fd *FramesDelete) setConfirm() {
	fd.Confirm = tuisupport.NewButton(
		fmt.Sprintf("frames-%s-delete-confirm", fd.ID),
		"Confirm",
		fd.confirm,
		15,
	)
}

func (fd *FramesDelete) setCancel() {
	fd.Cancel = tuisupport.NewButton(
		fmt.Sprintf("frames-%s-delete-cancel", fd.ID),
		"Cancel",
		fd.cancel,
		9,
	)
}

func (fd *FramesDelete) confirm() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = fd.Body.App.call(
		controllers.FramesDelete,
		&controllers.FramesDeleteParams{
			Frame: fd.ID,
		},
		tuisupport.Open(".."),
	)
	if result != nil {
		c = tuisupport.Open("../..")
	}
	return
}

func (fd *FramesDelete) cancel() (c tea.Cmd) {
	c = tuisupport.Open("..")
	return
}

func (fd *FramesDelete) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{
		fd.Cancel,
		fd.Confirm,
	}
	return
}

func (fd *FramesDelete) isFocus() (is bool) {
	is = fd.Cancel.IsFocus || fd.Confirm.IsFocus
	return
}

func (fd *FramesDelete) Focus(aspect string) (c tea.Cmd) {
	return
}

func (fd *FramesDelete) Blur() {}
