package tui

import (
	"fmt"
	"sf/controllers"
	"sf/models"
	"sf/tui/tuiform"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type FramesEdit struct {
	Body           *Body
	ID             string
	Reader         *models.FrameReader
	FormComponents []*models.FormComponent
	FormAnswers    map[string]string
	Form           *tuiform.Form
}

func newFramesEdit(b *Body, id string) (fe *FramesEdit) {
	fe = &FramesEdit{Body: b, ID: id}
	return
}

func (fe *FramesEdit) Init() (c tea.Cmd) {
	cs := []tea.Cmd{}
	cs = append(cs, fe.setReader())
	fe.setFormComponents()
	fe.setFormAnswers()
	fe.setForm()
	cs = append(cs, fe.Form.Init())
	c = tea.Batch(cs...)
	return
}

func (fe *FramesEdit) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = fe
	_, c = fe.Form.Update(msg)
	return
}

func (fe *FramesEdit) View() (v string) {
	v = fe.Form.View()
	return
}

func (fe *FramesEdit) setSize(w int, h int) {
	fe.Form.SetWidth(w)
}

func (fe *FramesEdit) setReader() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = fe.Body.App.call(
		controllers.FramesRead,
		&controllers.FramesReadParams{
			Frame: fe.ID,
		},
		tuisupport.Open(".."),
	)
	if result != nil {
		fe.Reader = result.Payload.(*models.FrameReader)
	}
	return
}

func (fe *FramesEdit) setFormAnswers() {
	fe.FormAnswers = map[string]string{
		"Name":  fe.Reader.Name,
		"About": fe.Reader.About,
	}
}

func (fe *FramesEdit) setFormComponents() {
	fe.FormComponents = []*models.FormComponent{
		{Key: "Name"},
		{Key: "About"},
	}
	for _, fmc := range fe.FormComponents {
		fmc.Load()
	}
}

func (fe *FramesEdit) setForm() {
	fe.Form = tuiform.NewForm(
		"workspaces-edit",
		"Edit workspace",
		fe.FormComponents,
		fe.FormAnswers,
		fe.submit,
		fe.cancel,
	)
}

func (fe *FramesEdit) submit(answers map[string]string) (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = fe.Body.App.call(
		controllers.FramesUpdate,
		&controllers.FramesUpdateParams{
			Frame:   fe.ID,
			Updates: answers,
		},
		func() tea.Msg {
			fe.Form.IsActive = true
			return fe.Body.App.setFocus()
		},
	)
	if result != nil {
		update := result.Payload.(*controllers.FramesUpdateResult)
		c = tuisupport.Open(fmt.Sprintf("/frames/@%s", update.To.Name))
	}
	return
}

func (fe *FramesEdit) cancel() (c tea.Cmd) {
	c = tuisupport.Open("..")
	return
}

func (fe *FramesEdit) focusChain() (fc []tuisupport.Focuser) {
	fc = fe.Form.FocusChain()
	return
}

func (fe *FramesEdit) isFocus() (is bool) {
	is = fe.Form.IsFocus()
	return
}
