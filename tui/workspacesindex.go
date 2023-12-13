package tui

import (
	"fmt"
	"sf/controllers"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type WorkspacesIndex struct {
	Body  *Body
	Items []*controllers.WorkspacesIndexItemResult
	Table *Table
}

func newWorkspacesIndex(b *Body) (wi *WorkspacesIndex) {
	wi = &WorkspacesIndex{Body: b}
	return
}

func (wi *WorkspacesIndex) Init() (c tea.Cmd) {
	wi.setItems()
	wi.setTable()
	// c = wi.Table.Init()
	return
}

func (wi *WorkspacesIndex) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = wi
	_, c = wi.Table.Update(msg)
	return
}

func (wi *WorkspacesIndex) View() (v string) {
	v = wi.Table.View()
	return
}

func (wi *WorkspacesIndex) setSize(w int, h int) {
	wi.Table.setSize(w, h)
}

func (wi *WorkspacesIndex) setItems() {
	result := wi.Body.call(
		controllers.WorkspacesIndex,
		nil,
		"/",
	)
	if result != nil {
		wi.Items = result.Payload.([]*controllers.WorkspacesIndexItemResult)
	}
}

func (wi *WorkspacesIndex) setTable() {
	propeties := []string{"Workspace", "About"}
	data := []map[string]string{}
	for _, w := range wi.Items {
		data = append(data, map[string]string{
			"ID":        w.Workspace,
			"Workspace": w.Workspace,
			"About":     w.About,
		})
	}
	navigator := func(id string) (p string) {
		p = fmt.Sprintf("/workspaces/@%s", id)
		return
	}
	wi.Table = NewTable("workspaces", propeties, data, navigator)
}

func (wi *WorkspacesIndex) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{wi.Table}
	return
}
