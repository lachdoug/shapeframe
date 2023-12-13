package tui

import (
	"fmt"
	"sf/controllers"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type ShapesIndex struct {
	Body  *Body
	Items []*controllers.ShapesIndexItemResult
	Table *Table
}

func newShapesIndex(b *Body) (si *ShapesIndex) {
	si = &ShapesIndex{Body: b}
	return
}

func (si *ShapesIndex) Init() (c tea.Cmd) {
	si.setItems()
	si.setTable()
	return
}

func (si *ShapesIndex) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = si
	_, c = si.Table.Update(msg)
	return
}

func (si *ShapesIndex) View() (v string) {
	v = si.Table.View()
	return
}

func (si *ShapesIndex) setSize(w int, h int) {
	si.Table.setSize(w, h)
}

func (si *ShapesIndex) setItems() {
	result := si.Body.call(
		controllers.ShapesIndex,
		nil,
		"/",
	)
	if result != nil {
		si.Items = result.Payload.([]*controllers.ShapesIndexItemResult)
	}
}

func (si *ShapesIndex) setTable() {
	propeties := []string{"Workspace", "Frame", "Shape", "About"}
	data := []map[string]string{}
	for _, s := range si.Items {
		data = append(data, map[string]string{
			"ID":        fmt.Sprintf("%s.%s.%s", s.Workspace, s.Frame, s.Shape),
			"Workspace": s.Workspace,
			"Frame":     s.Frame,
			"Shape":     s.Shape,
			"About":     s.About,
		})
	}
	navigator := func(id string) (p string) {
		p = fmt.Sprintf("/shapes/@%s", id)
		return
	}
	si.Table = NewTable("shapes", propeties, data, navigator)
}

func (si *ShapesIndex) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{si.Table}
	return
}
