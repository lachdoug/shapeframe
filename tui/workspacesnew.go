package tui

import (
	"fmt"
	"sf/controllers"
	"sf/models"
	"sf/tui/tuiform"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type WorkspacesNew struct {
	Body           *Body
	FormComponents []*models.FormComponent
	FormAnswers    map[string]string
	Form           *tuiform.Form
}

func newWorkspacesNew(b *Body) (wn *WorkspacesNew) {
	wn = &WorkspacesNew{Body: b}
	return
}

func (wn *WorkspacesNew) Init() (c tea.Cmd) {
	wn.setFormComponents()
	wn.setFormAnswers()
	wn.setForm()
	c = wn.Form.Init()
	return
}

func (wn *WorkspacesNew) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = wn
	_, c = wn.Form.Update(msg)
	return
}

func (wn *WorkspacesNew) View() (v string) {
	v = wn.Form.View()
	return
}

func (wn *WorkspacesNew) setSize(w int, h int) {
	wn.Form.SetSize(w, h)
}

func (wn *WorkspacesNew) setFormComponents() {
	wn.FormComponents = []*models.FormComponent{
		{Key: "Name", Required: true},
		{Key: "About"},
	}
	for _, fmc := range wn.FormComponents {
		fmc.Load()
	}
}

func (wn *WorkspacesNew) setFormAnswers() {
	wn.FormAnswers = map[string]string{}
}

func (wn *WorkspacesNew) setForm() {
	wn.Form = tuiform.NewForm(
		"workspaces-new",
		"New workspace",
		wn.FormComponents,
		nil,
		wn.submit,
		wn.cancel,
	)
}

func (wn *WorkspacesNew) submit(answers map[string]string) (c tea.Cmd) {
	result := wn.Body.call(
		controllers.WorkspacesCreate,
		&controllers.WorkspacesCreateParams{
			Workspace: answers["Name"],
			About:     answers["About"],
		},
		".",
	)
	if result != nil {
		id := result.Payload.(*controllers.WorkspacesCreateResult).Workspace
		c = Open(fmt.Sprintf("../@%s", id))
	}
	return
}

func (wn *WorkspacesNew) cancel() (c tea.Cmd) {
	c = Open("..")
	return
}

func (wn *WorkspacesNew) focusChain() (fc []tuisupport.Focuser) {
	fc = wn.Form.FocusChain()
	return
}
