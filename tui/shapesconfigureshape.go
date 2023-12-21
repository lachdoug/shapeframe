package tui

import (
	"sf/controllers"
	"sf/models"
	"sf/tui/tuiform"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type ShapesConfigureShape struct {
	Body        *Body
	FID         string
	ID          string
	Reader      *models.ShapeReader
	FormAnswers map[string]string
	Form        *tuiform.Form
	Width       int
	Height      int
}

func newShapesConfigureShape(b *Body, fid string, id string) (scgs *ShapesConfigureShape) {
	scgs = &ShapesConfigureShape{Body: b, FID: fid, ID: id}
	return
}

func (scgs *ShapesConfigureShape) Init() (c tea.Cmd) {
	cs := []tea.Cmd{}
	cs = append(cs, scgs.setReader())
	scgs.setFormAnswers()
	scgs.setForm()
	cs = append(cs, scgs.Form.Init())
	c = tea.Batch(cs...)
	return
}

func (scgs *ShapesConfigureShape) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = scgs
	_, c = scgs.Form.Update(msg)
	return
}

func (scgs *ShapesConfigureShape) View() (v string) {
	v = scgs.Form.View()
	return
}

func (scgs *ShapesConfigureShape) setSize(w int, h int) {
	scgs.Width = w
	scgs.Height = h
	scgs.Form.SetWidth(scgs.Width)
}

func (scgs *ShapesConfigureShape) setReader() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = scgs.Body.App.call(
		controllers.ShapesRead,
		&controllers.ShapesReadParams{
			Frame: scgs.FID,
			Shape: scgs.ID,
		},
		tuisupport.Open(".."),
	)
	if result != nil {
		scgs.Reader = result.Payload.(*models.ShapeReader)
	}
	return
}

func (scgs *ShapesConfigureShape) setFormAnswers() {
	scgs.FormAnswers = scgs.Reader.Configuration.Shape.Settings
}

func (scgs *ShapesConfigureShape) setForm() {
	scgs.Form = tuiform.NewForm(
		"shapes-configure-shape",
		"Configure shape",
		scgs.Reader.Configuration.Shape.Form,
		scgs.FormAnswers,
		scgs.submit,
		scgs.cancel,
	)
}

func (scgs *ShapesConfigureShape) submit(answers map[string]string) (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = scgs.Body.App.call(
		controllers.ShapeConfigurationsUpdate,
		&controllers.ShapeConfigurationsUpdateParams{
			Frame:   scgs.FID,
			Shape:   scgs.ID,
			Updates: answers,
		},
		func() tea.Msg {
			scgs.Form.IsActive = true
			return scgs.Body.App.setFocus()
		},
	)
	if result != nil {
		c = tuisupport.Open("..")
	}
	return
}

func (scgs *ShapesConfigureShape) cancel() (c tea.Cmd) {
	c = tuisupport.Open("..")
	return
}

func (scgs *ShapesConfigureShape) focusChain() (fc []tuisupport.Focuser) {
	fc = scgs.Form.FocusChain()
	return
}

func (scgs *ShapesConfigureShape) isFocus() (is bool) {
	is = scgs.Form.IsFocus()
	return
}
