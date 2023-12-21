package tui

import (
	"fmt"
	"sf/controllers"
	"sf/models"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FramesShow struct {
	Body                  *Body
	ID                    string
	Reader                *models.FrameReader
	Label                 *tuisupport.LinkBox
	ShapeLinks            []*tuisupport.Link
	ConfigureFrameLinkBox *tuisupport.LinkBox
	Delete                *tuisupport.Delete
}

func newFramesShow(b *Body, id string) (fs *FramesShow) {
	fs = &FramesShow{Body: b, ID: id}
	return
}

func (fs *FramesShow) Init() (c tea.Cmd) {
	c = tea.Batch(
		fs.setReader(),
		fs.setShapeLinks(),
		fs.setLabel(),
		fs.setConfigureFrameLinkBox(),
		fs.setDelete(),
	)
	return
}

func (fs *FramesShow) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = fs
	cs := []tea.Cmd{}
	_, c = fs.Label.Update(msg)
	cs = append(cs, c)
	for _, lk := range fs.ShapeLinks {
		_, c = lk.Update(msg)
		cs = append(cs, c)
	}
	_, c = fs.ConfigureFrameLinkBox.Update(msg)
	cs = append(cs, c)
	_, c = fs.Delete.Update(msg)
	cs = append(cs, c)
	c = tea.Batch(cs...)
	return
}

func (fs *FramesShow) View() (v string) {
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(1)

	if fs.isFocus() {
		style.
			BorderForeground(lipgloss.Color("15"))
	} else {
		style.
			Foreground(lipgloss.Color("7")).
			BorderForeground(lipgloss.Color("8"))
	}

	v = lipgloss.JoinVertical(lipgloss.Left,
		style.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				fs.viewLabel(),
				fs.viewShapeLinks(),
				fs.viewConfiguration(),
				// string(utils.YamlMarshal(fs.Reader)),
			),
		),
		fs.Delete.View(),
	)
	return
}

func (fs *FramesShow) viewLabel() (v string) {
	style := lipgloss.NewStyle().PaddingBottom(1)
	v = style.Render(fs.Label.View())
	return
}

func (fs *FramesShow) viewShapeLinks() (v string) {
	style := lipgloss.NewStyle().PaddingBottom(1)
	indentStyle := lipgloss.NewStyle().PaddingLeft(2)
	vs := []string{}
	for _, lk := range fs.ShapeLinks {
		vs = append(vs, lk.View())
	}
	v = style.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			"Shapes",
			indentStyle.Render(
				lipgloss.JoinVertical(lipgloss.Left, vs...),
			),
		),
	)
	return
}

func (fs *FramesShow) viewConfiguration() (v string) {
	indentStyle := lipgloss.NewStyle().PaddingLeft(2)
	vs := []string{"Configuration"}
	vs = append(vs, indentStyle.Render(fs.ConfigureFrameLinkBox.View()))
	v = lipgloss.JoinVertical(lipgloss.Left, vs...)
	return
}

func (fs *FramesShow) setSize(w int, h int) {
	fs.Label.SetSize(w - 4)
	fs.ConfigureFrameLinkBox.SetSize(w - 6)
	fs.Delete.SetSize(w)
}

func (fs *FramesShow) setReader() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = fs.Body.App.call(
		controllers.FramesRead,
		&controllers.FramesReadParams{
			Frame: fs.ID,
		},
		tuisupport.Open(".."),
	)
	if result != nil {
		fs.Reader = result.Payload.(*models.FrameReader)
	}
	return
}

func (fs *FramesShow) setLabel() (c tea.Cmd) {
	style := lipgloss.NewStyle().Bold(true)
	fs.Label = tuisupport.NewLinkBox(
		fmt.Sprintf("%s %s", style.Render(fs.ID), fs.Reader.About),
		fmt.Sprintf("frame-%s-label-link", fs.ID),
		"Edit",
		"label",
	)
	c = fs.Label.Init()
	return
}

func (fs *FramesShow) setShapeLinks() (c tea.Cmd) {
	for _, s := range fs.Reader.Shapes {
		lk := tuisupport.NewLink(
			fmt.Sprintf("frame-%s-shape-%s-link", fs.ID, s),
			s,
			fmt.Sprintf("/shapes/@%s.%s", fs.ID, s),
		)
		fs.ShapeLinks = append(fs.ShapeLinks, lk)
	}
	return
}

func (fs *FramesShow) setConfigureFrameLinkBox() (c tea.Cmd) {
	t := configurationInfo("Frame", fs.Reader.Configuration)
	fs.ConfigureFrameLinkBox = tuisupport.NewLinkBox(
		t,
		fmt.Sprintf("frame-%s-configure-frame-link", fs.ID),
		"Edit",
		"configure-frame",
	)
	c = fs.ConfigureFrameLinkBox.Init()
	return
}

func (fs *FramesShow) setDelete() (c tea.Cmd) {
	fs.Delete = tuisupport.NewDelete(
		fmt.Sprintf("frame-%s", fs.ID),
	)
	c = fs.Delete.Init()
	return
}

func (fs *FramesShow) isFocus() (is bool) {
	if fs.Label.IsFocus() {
		is = true
		return
	}
	for _, lk := range fs.ShapeLinks {
		if lk.IsFocus {
			is = true
			return
		}
	}
	if fs.ConfigureFrameLinkBox.IsFocus() {
		is = true
		return
	}
	return
}

func (fs *FramesShow) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{
		fs.Label,
	}
	for _, lk := range fs.ShapeLinks {
		fc = append(fc, lk)
	}
	fc = append(fc, fs.ConfigureFrameLinkBox)
	fc = append(fc, fs.Delete)
	return
}
