package tui

import (
	"fmt"
	"sf/controllers"
	"sf/tui/tuisupport"

	tea "github.com/charmbracelet/bubbletea"
)

type FramesIndex struct {
	Body  *Body
	Items []*controllers.FramesIndexItemResult
	Table *Table
}

func newFramesIndex(b *Body) (wi *FramesIndex) {
	wi = &FramesIndex{Body: b}
	return
}

func (wi *FramesIndex) Init() (c tea.Cmd) {
	wi.setItems()
	wi.setTable()
	// c = wi.Table.Init()
	return
}

func (wi *FramesIndex) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = wi
	_, c = wi.Table.Update(msg)
	return
}

func (wi *FramesIndex) View() (v string) {
	v = wi.Table.View()
	return
}

func (wi *FramesIndex) setSize(w int, h int) {
	wi.Table.setSize(w, h)
}

func (wi *FramesIndex) setItems() {
	result := wi.Body.call(
		controllers.FramesIndex,
		nil,
		"/",
	)
	if result != nil {
		wi.Items = result.Payload.([]*controllers.FramesIndexItemResult)
	}
}

func (wi *FramesIndex) setTable() {
	propeties := []string{"Workspace", "Frame", "About"}
	data := []map[string]string{}
	for _, f := range wi.Items {
		data = append(data, map[string]string{
			"ID":        fmt.Sprintf("%s.%s", f.Workspace, f.Frame),
			"Workspace": f.Workspace,
			"Frame":     f.Frame,
			"About":     f.About,
		})
	}
	navigator := func(id string) (p string) {
		p = fmt.Sprintf("/frames/@%s", id)
		return
	}
	wi.Table = NewTable("frames", propeties, data, navigator)
}

func (wi *FramesIndex) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{wi.Table}
	return
}
