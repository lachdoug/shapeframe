package tui

import (
	"sf/controllers"
	"sf/models"
	"sf/tui/tuiform"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type ShapesConfigureFrame struct {
	Body        *Body
	FID         string
	ID          string
	Reader      *models.ShapeReader
	FormAnswers map[string]string
	Form        *tuiform.Form
	Width       int
	Height      int
}

func newShapesConfigureFrame(b *Body, fid string, id string) (scgf *ShapesConfigureFrame) {
	scgf = &ShapesConfigureFrame{Body: b, FID: fid, ID: id}
	return
}

func (scgf *ShapesConfigureFrame) Init() (c tea.Cmd) {
	cs := []tea.Cmd{}
	cs = append(cs, scgf.setReader())
	scgf.setFormAnswers()
	scgf.setForm()
	cs = append(cs, scgf.Form.Init())
	c = tea.Batch(cs...)
	return
}

func (scgf *ShapesConfigureFrame) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = scgf
	_, c = scgf.Form.Update(msg)
	return
}

func (scgf *ShapesConfigureFrame) View() (v string) {
	v = scgf.Form.View()
	return
}

func (scgf *ShapesConfigureFrame) setSize(w int, h int) {
	scgf.Width = w
	scgf.Height = h
	scgf.Form.SetWidth(scgf.Width)
}

func (scgf *ShapesConfigureFrame) setReader() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = scgf.Body.App.call(
		controllers.ShapesRead,
		&controllers.ShapesReadParams{
			Frame: scgf.FID,
			Shape: scgf.ID,
		},
		tuisupport.Open(".."),
	)
	if result != nil {
		scgf.Reader = result.Payload.(*models.ShapeReader)
	}
	return
}

func (scgf *ShapesConfigureFrame) setFormAnswers() {
	scgf.FormAnswers = scgf.Reader.Configuration.Frame.Settings
}

func (scgf *ShapesConfigureFrame) setForm() {
	scgf.Form = tuiform.NewForm(
		"shapes-configure-frame",
		"Configure shape frame",
		scgf.Reader.Configuration.Frame.Form,
		scgf.FormAnswers,
		scgf.submit,
		scgf.cancel,
	)
}

func (scgf *ShapesConfigureFrame) submit(answers map[string]string) (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = scgf.Body.App.call(
		controllers.ShapeFrameConfigurationsUpdate,
		&controllers.ShapeFrameConfigurationsUpdateParams{
			Frame:   scgf.FID,
			Shape:   scgf.ID,
			Updates: answers,
		},
		func() tea.Msg {
			scgf.Form.IsActive = true
			return scgf.Body.App.setFocus()
		},
	)
	if result != nil {
		c = tuisupport.Open("..")
	}
	return
}

func (scgf *ShapesConfigureFrame) cancel() (c tea.Cmd) {
	c = tuisupport.Open("..")
	return
}

func (scgf *ShapesConfigureFrame) focusChain() (fc []tuisupport.Focuser) {
	fc = scgf.Form.FocusChain()
	return
}

func (scgf *ShapesConfigureFrame) isFocus() (is bool) {
	is = scgf.Form.IsFocus()
	return
}
