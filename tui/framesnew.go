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
	FormComponents []*models.FormComponent
	FormAnswers    map[string]string
	Form           *tuiform.Form
	WorkspaceItems []*controllers.WorkspacesIndexItemResult
}

func newFramesNew(b *Body) (fn *FramesNew) {
	fn = &FramesNew{Body: b}
	return
}

func (fn *FramesNew) Init() (c tea.Cmd) {
	fn.setWorkspaceItems()
	fn.setFormComponents()
	fn.setFormAnswers()
	fn.setForm()
	c = fn.Form.Init()
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
	fn.Form.SetSize(w, h)
}

func (fn *FramesNew) workspaceOptions() (opts []*models.FormOption) {
	opts = []*models.FormOption{}
	for _, w := range fn.WorkspaceItems {
		opt := &models.FormOption{Value: w.Workspace}
		opts = append(opts, opt)
	}
	return
}

func (fn *FramesNew) setFormComponents() {
	fn.FormComponents = []*models.FormComponent{
		{Key: "Workspace", Type: "selects", Options: fn.workspaceOptions()},
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

func (fn *FramesNew) setWorkspaceItems() {
	result := fn.Body.call(
		controllers.WorkspacesIndex,
		nil,
		"/",
	)
	if result != nil {
		fn.WorkspaceItems = result.Payload.([]*controllers.WorkspacesIndexItemResult)
	}
}

func (fn *FramesNew) setForm() {
	fn.Form = tuiform.NewForm(
		"workspaces-new",
		"New workspace",
		fn.FormComponents,
		nil,
		fn.submit,
		fn.cancel,
	)
}

func (fn *FramesNew) submit(answers map[string]string) (c tea.Cmd) {
	result := fn.Body.call(
		controllers.FramesCreate,
		&controllers.FramesCreateParams{
			Frame: answers["Name"],
			About: answers["About"],
		},
		".",
	)
	if result != nil {
		id := result.Payload.(*controllers.FramesCreateResult).Frame
		c = Open(fmt.Sprintf("../@%s", id))
	}
	return
}

func (fn *FramesNew) cancel() (c tea.Cmd) {
	c = Open("..")
	return
}

func (fn *FramesNew) focusChain() (fc []tuisupport.Focuser) {
	fc = fn.Form.FocusChain()
	return
}
