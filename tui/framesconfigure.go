package tui

import (
	"sf/controllers"
	"sf/models"
	"sf/tui/tuiform"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type FramesConfigure struct {
	Body        *Body
	ID          string
	Reader      *models.FrameReader
	FormAnswers map[string]string
	Form        *tuiform.Form
	Width       int
	Height      int
}

func newFramesConfigure(b *Body, id string) (fcg *FramesConfigure) {
	fcg = &FramesConfigure{Body: b, ID: id}
	return
}

func (fcg *FramesConfigure) Init() (c tea.Cmd) {
	cs := []tea.Cmd{}
	cs = append(cs, fcg.setReader())
	fcg.setFormAnswers()
	fcg.setForm()
	cs = append(cs, fcg.Form.Init())
	c = tea.Batch(cs...)
	return
}

func (fcg *FramesConfigure) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = fcg
	_, c = fcg.Form.Update(msg)
	return
}

func (fcg *FramesConfigure) View() (v string) {
	v = fcg.Form.View()
	return
}

func (fcg *FramesConfigure) setSize(w int, h int) {
	fcg.Width = w
	fcg.Height = h
	fcg.Form.SetWidth(fcg.Width)
}

func (fcg *FramesConfigure) setReader() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = fcg.Body.App.call(
		controllers.FramesRead,
		&controllers.FramesReadParams{
			Frame: fcg.ID,
		},
		tuisupport.Open(".."),
	)
	if result != nil {
		fcg.Reader = result.Payload.(*models.FrameReader)
	}
	return
}

func (fcg *FramesConfigure) setFormAnswers() {
	fcg.FormAnswers = fcg.Reader.Configuration.Settings
}

func (fcg *FramesConfigure) setForm() {
	fcg.Form = tuiform.NewForm(
		"frames-new",
		"New frame",
		fcg.Reader.Configuration.Form,
		fcg.FormAnswers,
		fcg.submit,
		fcg.cancel,
	)
}

func (fcg *FramesConfigure) submit(answers map[string]string) (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = fcg.Body.App.call(
		controllers.FrameConfigurationsUpdate,
		&controllers.FrameConfigurationsUpdateParams{
			Frame:   fcg.ID,
			Updates: answers,
		},
		func() tea.Msg {
			fcg.Form.IsActive = true
			return fcg.Body.App.setFocus
		},
	)
	if result != nil {
		c = tuisupport.Open("..")
	}
	return
}

func (fcg *FramesConfigure) cancel() (c tea.Cmd) {
	c = tuisupport.Open("..")
	return
}

func (fcg *FramesConfigure) focusChain() (fc []tuisupport.Focuser) {
	fc = fcg.Form.FocusChain()
	return
}

func (fcg *FramesConfigure) isFocus() (is bool) {
	is = fcg.Form.IsFocus()
	return
}
