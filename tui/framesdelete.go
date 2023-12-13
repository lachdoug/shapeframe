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
	WID     string
	ID      string
	Cancel  *tuisupport.Button
	Confirm *tuisupport.Button
}

func newFramesDelete(b *Body, wid string, id string) (fd *FramesDelete) {
	fd = &FramesDelete{Body: b, WID: wid, ID: id}
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
	id := fmt.Sprintf("%s.%s", fd.WID, fd.ID)
	msg := style.Render(fmt.Sprintf("Delete frame %s?", idStyle.Render(id)))
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
	result := fd.Body.call(
		controllers.FramesDelete,
		&controllers.FramesDeleteParams{
			Frame: fd.ID,
		},
		"../..",
	)
	if result != nil {
		c = Open("../..")
	}
	return
}

func (fd *FramesDelete) cancel() (c tea.Cmd) {
	c = Open("..")
	return
}

func (fd *FramesDelete) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{
		fd.Cancel,
		fd.Confirm,
	}
	return
}

func (fd *FramesDelete) Focus(aspect string) (c tea.Cmd) {
	return
}

func (fd *FramesDelete) Blur() {}
