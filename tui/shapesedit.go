package tui

import (
	"fmt"
	"sf/controllers"
	"sf/models"
	"sf/tui/tuiform"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type ShapesEdit struct {
	Body           *Body
	FID            string
	ID             string
	Reader         *models.ShapeReader
	FormComponents []*models.FormComponent
	FormAnswers    map[string]string
	Form           *tuiform.Form
}

func newShapesEdit(b *Body, fid string, id string) (se *ShapesEdit) {
	se = &ShapesEdit{Body: b, FID: fid, ID: id}
	return
}

func (se *ShapesEdit) Init() (c tea.Cmd) {
	cs := []tea.Cmd{}
	cs = append(cs, se.setReader())
	se.setFormComponents()
	se.setFormAnswers()
	se.setForm()
	cs = append(cs, se.Form.Init())
	c = tea.Batch(cs...)
	return
}

func (se *ShapesEdit) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = se
	_, c = se.Form.Update(msg)
	return
}

func (se *ShapesEdit) View() (v string) {
	v = se.Form.View()
	return
}

func (se *ShapesEdit) setSize(w int, h int) {
	se.Form.SetWidth(w)
}

func (se *ShapesEdit) setReader() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = se.Body.App.call(
		controllers.ShapesRead,
		&controllers.ShapesReadParams{
			Shape: se.ID,
		},
		tuisupport.Open(".."),
	)
	if result != nil {
		se.Reader = result.Payload.(*models.ShapeReader)
	}
	return
}

func (se *ShapesEdit) setFormAnswers() {
	se.FormAnswers = map[string]string{
		"Name":  se.Reader.Name,
		"About": se.Reader.About,
	}
}

func (se *ShapesEdit) setFormComponents() {
	se.FormComponents = []*models.FormComponent{
		{Key: "Name"},
		{Key: "About"},
	}
	for _, fmc := range se.FormComponents {
		fmc.Load()
	}
}

func (se *ShapesEdit) setForm() {
	se.Form = tuiform.NewForm(
		"workspaces-edit",
		"Edit workspace",
		se.FormComponents,
		se.FormAnswers,
		se.submit,
		se.cancel,
	)
}

func (se *ShapesEdit) submit(answers map[string]string) (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = se.Body.App.call(
		controllers.ShapesUpdate,
		&controllers.ShapesUpdateParams{
			Frame:   se.FID,
			Shape:   se.ID,
			Updates: answers,
		},
		func() tea.Msg {
			se.Form.IsActive = true
			return se.Body.App.setFocus()
		},
	)
	if result != nil {
		update := result.Payload.(*controllers.ShapesUpdateResult)
		c = tuisupport.Open(fmt.Sprintf("/shapes/@%s.%s", update.Frame, update.To.Name))
	}
	return
}

func (se *ShapesEdit) cancel() (c tea.Cmd) {
	c = tuisupport.Open("..")
	return
}

func (se *ShapesEdit) focusChain() (fc []tuisupport.Focuser) {
	fc = se.Form.FocusChain()
	return
}

func (se *ShapesEdit) isFocus() (is bool) {
	is = se.Form.IsFocus()
	return
}
