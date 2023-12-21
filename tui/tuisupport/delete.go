package tuisupport

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Delete struct {
	ID    string
	Link  *Link
	Width int
}

func NewDelete(id string) (d *Delete) {
	d = &Delete{ID: id}
	return
}

func (d *Delete) Init() (c tea.Cmd) {
	c = d.setLink()
	return
}

func (d *Delete) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	_, c = d.Link.Update(msg)
	return
}

func (d *Delete) View() (v string) {
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Width(d.Width).
		Align(lipgloss.Right)

	if d.isFocus() || d.isHover() {
		style.
			BorderForeground(lipgloss.Color("9"))
	} else {
		style.
			Foreground(lipgloss.Color("7")).
			BorderForeground(lipgloss.Color("8"))
	}

	v = style.Render(d.Link.View())
	return
}

func (d *Delete) setLink() (c tea.Cmd) {
	d.Link = NewLink(
		fmt.Sprintf("%s-delete-link", d.ID),
		"Delete",
		"delete",
	)
	return
}

func (d *Delete) SetSize(w int) {
	d.Width = w - 2
}

func (d *Delete) isFocus() (is bool) {
	is = d.Link.IsFocus
	return
}

func (d *Delete) isHover() (is bool) {
	is = d.Link.IsHover
	return
}

func (d *Delete) Focus(_ string) (c tea.Cmd) {
	d.Link.IsFocus = true
	return
}

func (d *Delete) Blur() {
	d.Link.IsFocus = false
}
