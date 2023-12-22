package tui

import (
	"fmt"
	"sf/controllers"
	"sf/models"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ShapesShow struct {
	Body                  *Body
	FID                   string
	ID                    string
	Reader                *models.ShapeReader
	Label                 *tuisupport.LinkBox
	FrameLink             *tuisupport.Link
	ConfigureShapeLinkBox *tuisupport.LinkBox
	ConfigureFrameLinkBox *tuisupport.LinkBox
	Delete                *tuisupport.Delete
}

func newShapesShow(b *Body, fid string, id string) (ss *ShapesShow) {
	ss = &ShapesShow{Body: b, FID: fid, ID: id}
	return
}

func (ss *ShapesShow) Init() (c tea.Cmd) {
	c = ss.setReader()
	if ss.Reader == nil {
		return
	}
	c = tea.Batch(
		c,
		ss.setFrameLink(),
		ss.setLabel(),
		ss.setConfigureShapeLinkBox(),
		ss.setConfigureFrameLinkBox(),
		ss.setDelete(),
	)
	return
}

func (ss *ShapesShow) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = ss
	cs := []tea.Cmd{}
	_, c = ss.Label.Update(msg)
	cs = append(cs, c)
	_, c = ss.FrameLink.Update(msg)
	cs = append(cs, c)
	_, c = ss.ConfigureShapeLinkBox.Update(msg)
	cs = append(cs, c)
	_, c = ss.ConfigureFrameLinkBox.Update(msg)
	cs = append(cs, c)
	_, c = ss.Delete.Update(msg)
	cs = append(cs, c)
	c = tea.Batch(cs...)
	return
}

func (ss *ShapesShow) View() (v string) {
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(1)

	if ss.isFocus() {
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
				ss.viewLabel(),
				ss.viewFrame(),
				ss.viewConfiguration(),
				// string(utils.YamlMarshal(ss.Reader)),
			),
		),
		ss.Delete.View(),
	)
	return
}

func (ss *ShapesShow) viewLabel() (v string) {
	style := lipgloss.NewStyle().PaddingBottom(1)
	v = style.Render(ss.Label.View())
	return
}

func (ss *ShapesShow) viewFrame() (v string) {
	style := lipgloss.NewStyle().PaddingBottom(1)
	v = style.Render(fmt.Sprintf("Frame: %s", ss.FrameLink.View()))
	return
}

func (ss *ShapesShow) viewConfiguration() (v string) {
	indentStyle := lipgloss.NewStyle().PaddingLeft(2)
	vs := []string{"Configuration"}
	vs = append(vs, indentStyle.Render(ss.ConfigureShapeLinkBox.View()))
	vs = append(vs, indentStyle.Render(ss.ConfigureFrameLinkBox.View()))
	v = lipgloss.JoinVertical(lipgloss.Left, vs...)
	return
}

func (ss *ShapesShow) setSize(w int, h int) {
	ss.Label.SetSize(w - 4)
	ss.ConfigureShapeLinkBox.SetSize(w - 6)
	ss.ConfigureFrameLinkBox.SetSize(w - 6)
	ss.Delete.SetSize(w)
}

func (ss *ShapesShow) setReader() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = ss.Body.App.call(
		controllers.ShapesRead,
		&controllers.ShapesReadParams{
			Frame: ss.FID,
			Shape: ss.ID,
		},
		tuisupport.Open(".."),
	)
	if result != nil {
		ss.Reader = result.Payload.(*models.ShapeReader)
	}
	return
}

func (ss *ShapesShow) setLabel() (c tea.Cmd) {
	style := lipgloss.NewStyle().Bold(true)
	ss.Label = tuisupport.NewLinkBox(
		fmt.Sprintf("%s %s", style.Render(ss.ID), ss.Reader.About),
		fmt.Sprintf("frame-%s-shape-%s-label-link", ss.FID, ss.ID),
		"Edit",
		"label",
	)
	c = ss.Label.Init()
	return
}

func (ss *ShapesShow) setFrameLink() (c tea.Cmd) {
	ss.FrameLink = tuisupport.NewLink(
		fmt.Sprintf("frame-%s-shape-%s-frame-link", ss.FID, ss.ID),
		ss.FID,
		fmt.Sprintf("/frames/@%s", ss.FID),
	)
	return
}

func (ss *ShapesShow) setConfigureShapeLinkBox() (c tea.Cmd) {
	t := configurationInfo("Shape", ss.Reader.Configuration.Shape)
	ss.ConfigureShapeLinkBox = tuisupport.NewLinkBox(
		t,
		fmt.Sprintf("frame-%s-shape-%s-configure-shape-link", ss.FID, ss.ID),
		"Edit",
		"configure-shape",
	)
	c = ss.ConfigureShapeLinkBox.Init()
	return
}

func (ss *ShapesShow) setConfigureFrameLinkBox() (c tea.Cmd) {
	t := configurationInfo("Frame", ss.Reader.Configuration.Frame)
	ss.ConfigureFrameLinkBox = tuisupport.NewLinkBox(
		t,
		fmt.Sprintf("frame-%s-shape-%s-configure-frame-link", ss.FID, ss.ID),
		"Edit",
		"configure-frame",
	)
	c = ss.ConfigureFrameLinkBox.Init()
	return
}

func (ss *ShapesShow) setDelete() (c tea.Cmd) {
	ss.Delete = tuisupport.NewDelete(
		fmt.Sprintf("frame-%s-shape-%s", ss.FID, ss.ID),
	)
	c = ss.Delete.Init()
	return
}

func (ss *ShapesShow) isFocus() (is bool) {
	if ss.Label.IsFocus() {
		is = true
		return
	}
	if ss.FrameLink.IsFocus {
		is = true
		return
	}
	if ss.ConfigureShapeLinkBox.IsFocus() {
		is = true
		return
	}
	if ss.ConfigureFrameLinkBox.IsFocus() {
		is = true
		return
	}
	return
}

func (ss *ShapesShow) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{
		ss.Label,
		ss.FrameLink,
		ss.ConfigureShapeLinkBox,
		ss.ConfigureFrameLinkBox,
		ss.Delete,
	}
	return
}
