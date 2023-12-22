package tui

import (
	"fmt"
	"sf/controllers"
	"sf/models"
	"sf/tui/tuiform"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type ShapesNew struct {
	Body           *Body
	ShaperItems    []*controllers.ShapersIndexItemResult
	FrameItems     []*controllers.FramesIndexItemResult
	FormComponents []*models.FormComponent
	FormAnswers    map[string]string
	Form           *tuiform.Form
}

func newShapesNew(b *Body) (sn *ShapesNew) {
	sn = &ShapesNew{Body: b}
	return
}

func (sn *ShapesNew) Init() (c tea.Cmd) {
	cs := []tea.Cmd{}
	cs = append(cs, sn.setFrames())
	cs = append(cs, sn.setShapers())
	sn.setFormComponents()
	sn.setFormAnswers()
	sn.setForm()
	cs = append(cs, sn.Form.Init())
	c = tea.Batch(cs...)
	return
}

func (sn *ShapesNew) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = sn
	_, c = sn.Form.Update(msg)
	return
}

func (sn *ShapesNew) View() (v string) {
	v = sn.Form.View()
	return
}

func (sn *ShapesNew) setSize(w int, h int) {
	sn.Form.SetWidth(w)
}

func (sn *ShapesNew) setShapers() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = sn.Body.App.call(
		controllers.ShapersIndex,
		nil,
		tuisupport.Open(".."),
	)
	if result != nil {
		sn.ShaperItems = result.Payload.([]*controllers.ShapersIndexItemResult)
	}
	return
}

func (sn *ShapesNew) setFrames() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = sn.Body.App.call(
		controllers.FramesIndex,
		nil,
		tuisupport.Open(".."),
	)
	if result != nil {
		sn.FrameItems = result.Payload.([]*controllers.FramesIndexItemResult)
	}
	return
}

func (sn *ShapesNew) frameOptions() (opts []*models.FormOption) {
	opts = []*models.FormOption{}
	for _, f := range sn.FrameItems {
		opts = append(opts, &models.FormOption{
			Value: f.Frame,
		})
	}
	return
}

func (sn *ShapesNew) shaperOptions() (opts []*models.FormOption) {
	opts = []*models.FormOption{}
	for _, srspecitem := range sn.ShaperItems {
		opts = append(opts, &models.FormOption{
			Value: srspecitem.Shaper,
		})
	}
	return
}

func (sn *ShapesNew) setFormComponents() {
	sn.FormComponents = []*models.FormComponent{
		{Key: "Frame", Type: "select", Options: sn.frameOptions()},
		{Key: "Shaper", Type: "select", Options: sn.shaperOptions()},
		{Key: "Name"},
		{Key: "About"},
	}
	for _, fmc := range sn.FormComponents {
		fmc.Load()
	}
}

func (sn *ShapesNew) setFormAnswers() {
	sn.FormAnswers = map[string]string{}
}

func (sn *ShapesNew) setForm() {
	sn.Form = tuiform.NewForm(
		"shapes-new",
		"New shape",
		sn.FormComponents,
		nil,
		sn.submit,
		sn.cancel,
	)
}

func (sn *ShapesNew) submit(answers map[string]string) (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = sn.Body.App.call(
		controllers.ShapesCreate,
		&controllers.ShapesCreateParams{
			Frame:  answers["Frame"],
			Shaper: answers["Shaper"],
			Shape:  answers["Name"],
			About:  answers["About"],
		},
		func() tea.Msg {
			sn.Form.IsActive = true
			return sn.Body.App.setFocus()
		},
	)
	if result != nil {
		create := result.Payload.(*controllers.ShapesCreateResult)
		c = tuisupport.Open(fmt.Sprintf("../@%s.%s", create.Frame, create.Shape))
	}
	return
}

func (sn *ShapesNew) cancel() (c tea.Cmd) {
	c = tuisupport.Open("..")
	return
}

func (sn *ShapesNew) focusChain() (fc []tuisupport.Focuser) {
	fc = sn.Form.FocusChain()
	return
}

func (sn *ShapesNew) isFocus() (is bool) {
	is = sn.Form.IsFocus()
	return
}
