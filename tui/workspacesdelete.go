package tui

import (
	"fmt"
	"sf/controllers"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type WorkspacesDelete struct {
	Body    *Body
	ID      string
	Cancel  *tuisupport.Button
	Confirm *tuisupport.Button
}

func newWorkspacesDelete(b *Body, id string) (wd *WorkspacesDelete) {
	wd = &WorkspacesDelete{Body: b, ID: id}
	return
}

func (wd *WorkspacesDelete) Init() (c tea.Cmd) {
	wd.setCancel()
	wd.setConfirm()
	return
}

func (wd *WorkspacesDelete) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = wd
	cs := []tea.Cmd{}
	_, c = wd.Confirm.Update(msg)
	cs = append(cs, c)
	_, c = wd.Cancel.Update(msg)
	cs = append(cs, c)
	c = tea.Batch(cs...)
	return
}

func (wd *WorkspacesDelete) View() (v string) {
	style := lipgloss.NewStyle().Padding(1)
	idStyle := lipgloss.NewStyle().Bold(true)
	msg := style.Render(fmt.Sprintf("Delete workspace %s?", idStyle.Render(wd.ID)))
	v = lipgloss.JoinVertical(lipgloss.Left,
		msg,
		lipgloss.JoinHorizontal(lipgloss.Top,
			wd.Cancel.View(),
			wd.Confirm.View(),
		),
	)
	return
}

func (wd *WorkspacesDelete) setSize(w int, h int) {}

func (wd *WorkspacesDelete) setConfirm() {
	wd.Confirm = tuisupport.NewButton(
		fmt.Sprintf("wokspaces-%s-delete-confirm", wd.ID),
		"Confirm",
		wd.confirm,
		15,
	)
}

func (wd *WorkspacesDelete) setCancel() {
	wd.Cancel = tuisupport.NewButton(
		fmt.Sprintf("wokspaces-%s-delete-cancel", wd.ID),
		"Cancel",
		wd.cancel,
		9,
	)
}

func (wd *WorkspacesDelete) confirm() (c tea.Cmd) {
	result := wd.Body.call(
		controllers.WorkspacesDelete,
		&controllers.WorkspacesDeleteParams{
			Workspace: wd.ID,
		},
		"../..",
	)
	if result != nil {
		c = Open("../..")
	}
	return
}

func (wd *WorkspacesDelete) cancel() (c tea.Cmd) {
	c = Open("..")
	return
}

func (wd *WorkspacesDelete) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{
		wd.Cancel,
		wd.Confirm,
	}
	return
}

func (wd *WorkspacesDelete) Focus(aspect string) (c tea.Cmd) {
	return
}

func (wd *WorkspacesDelete) Blur() {
}
