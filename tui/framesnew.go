package tui

import (
	"fmt"
	"sf/controllers"
	"sf/models"
	"sf/tui/tuiform"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type FramesNew struct {
	Body           *Body
	FramerItems    []*controllers.FramersIndexItemResult
	FormComponents []*models.FormComponent
	FormAnswers    map[string]string
	Form           *tuiform.Form
}

func newFramesNew(b *Body) (fn *FramesNew) {
	fn = &FramesNew{Body: b}
	return
}

func (fn *FramesNew) Init() (c tea.Cmd) {
	cs := []tea.Cmd{}
	cs = append(cs, fn.setFramers())
	fn.setFormComponents()
	fn.setFormAnswers()
	fn.setForm()
	cs = append(cs, fn.Form.Init())
	c = tea.Batch(cs...)
	return
}

func (fn *FramesNew) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = fn
	_, c = fn.Form.Update(msg)
	return
}

func (fn *FramesNew) View() (v string) {
	v = fn.Form.View()
	return
}

func (fn *FramesNew) setSize(w int, h int) {
	fn.Form.SetWidth(w)
}

func (fn *FramesNew) setFramers() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = fn.Body.App.call(
		controllers.FramersIndex,
		nil,
		tuisupport.Open(".."),
	)
	if result != nil {
		fn.FramerItems = result.Payload.([]*controllers.FramersIndexItemResult)
	}
	return
}

func (fn *FramesNew) framerOptions() (opts []*models.FormOption) {
	opts = []*models.FormOption{}
	for _, fr := range fn.FramerItems {
		opts = append(opts, &models.FormOption{
			Value: fr.Framer,
		})
	}
	return
}

func (fn *FramesNew) setFormComponents() {
	fn.FormComponents = []*models.FormComponent{
		{Key: "Framer", Type: "select", Options: fn.framerOptions()},
		{Key: "Name"},
		{Key: "About"},
	}
	for _, fmc := range fn.FormComponents {
		fmc.Load()
	}
}

func (fn *FramesNew) setFormAnswers() {
	fn.FormAnswers = map[string]string{}
}

func (fn *FramesNew) setForm() {
	fn.Form = tuiform.NewForm(
		"frames-new",
		"New frame",
		fn.FormComponents,
		fn.FormAnswers,
		fn.submit,
		fn.cancel,
	)
}

func (fn *FramesNew) submit(answers map[string]string) (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = fn.Body.App.call(
		controllers.FramesCreate,
		&controllers.FramesCreateParams{
			Framer: answers["Framer"],
			Frame:  answers["Name"],
			About:  answers["About"],
		},
		func() tea.Msg {
			fn.Form.IsActive = true
			return fn.Body.App.setFocus()
		},
	)
	if result != nil {
		id := result.Payload.(*controllers.FramesCreateResult).Frame
		c = tuisupport.Open(fmt.Sprintf("../@%s", id))
	}
	return
}

func (fn *FramesNew) cancel() (c tea.Cmd) {
	c = tuisupport.Open("..")
	return
}

func (fn *FramesNew) focusChain() (fc []tuisupport.Focuser) {
	fc = fn.Form.FocusChain()
	return
}

func (fn *FramesNew) isFocus() (is bool) {
	is = fn.Form.IsFocus()
	return
}
