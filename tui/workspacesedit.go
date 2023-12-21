package tui

import (
	"sf/controllers"
	"sf/models"
	"sf/tui/tuiform"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type WorkspacesEdit struct {
	Body           *Body
	Reader         *models.WorkspaceReader
	FormComponents []*models.FormComponent
	FormAnswers    map[string]string
	Form           *tuiform.Form
}

func newWorkspacesEdit(b *Body) (we *WorkspacesEdit) {
	we = &WorkspacesEdit{Body: b}
	return
}

func (we *WorkspacesEdit) Init() (c tea.Cmd) {
	cs := []tea.Cmd{}
	cs = append(cs, we.setReader())
	we.setFormComponents()
	we.setFormAnswers()
	we.setForm()
	cs = append(cs, we.Form.Init())
	c = tea.Batch(cs...)
	return
}

func (we *WorkspacesEdit) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = we
	_, c = we.Form.Update(msg)
	return
}

func (we *WorkspacesEdit) View() (v string) {
	v = we.Form.View()
	return
}

func (we *WorkspacesEdit) setSize(w int, h int) {
	we.Form.SetWidth(w)
}

func (we *WorkspacesEdit) setReader() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = we.Body.App.call(
		controllers.WorkspacesRead,
		nil,
		tuisupport.Open(".."),
	)
	if result != nil {
		we.Reader = result.Payload.(*models.WorkspaceReader)
	}
	return
}

func (we *WorkspacesEdit) setFormAnswers() {
	we.FormAnswers = map[string]string{
		"Name":  we.Reader.Name,
		"About": we.Reader.About,
	}
}

func (we *WorkspacesEdit) setFormComponents() {
	we.FormComponents = []*models.FormComponent{
		{Key: "Name"},
		{Key: "About"},
	}
	for _, fmc := range we.FormComponents {
		fmc.Load()
	}
}

func (we *WorkspacesEdit) setForm() {
	we.Form = tuiform.NewForm(
		"workspaces-edit",
		"Edit workspace",
		we.FormComponents,
		we.FormAnswers,
		we.submit,
		we.cancel,
	)
}

func (we *WorkspacesEdit) submit(answers map[string]string) (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = we.Body.App.call(
		controllers.WorkspacesUpdate,
		&controllers.WorkspacesUpdateParams{
			Updates: answers,
		},
		func() tea.Msg {
			we.Form.IsActive = true
			return we.Body.App.setFocus()
		},
	)
	if result != nil {
		c = tuisupport.Reload()
	}
	return
}

func (we *WorkspacesEdit) cancel() (c tea.Cmd) {
	c = tuisupport.Open("..")
	return
}

func (we *WorkspacesEdit) focusChain() (fc []tuisupport.Focuser) {
	fc = we.Form.FocusChain()
	return
}

func (we *WorkspacesEdit) isFocus() (is bool) {
	is = we.Form.IsFocus()
	return
}
