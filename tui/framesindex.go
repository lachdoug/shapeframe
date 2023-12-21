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

func newFramesIndex(b *Body) (fi *FramesIndex) {
	fi = &FramesIndex{Body: b}
	return
}

func (fi *FramesIndex) Init() (c tea.Cmd) {
	c = fi.setItems()
	fi.setTable()
	return
}

func (fi *FramesIndex) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = fi
	_, c = fi.Table.Update(msg)
	return
}

func (fi *FramesIndex) View() (v string) {
	v = fi.Table.View()
	return
}

func (fi *FramesIndex) setSize(w int, h int) {
	fi.Table.setSize(w, h)
}

func (fi *FramesIndex) setItems() (c tea.Cmd) {
	result := &controllers.Result{}
	result, c = fi.Body.App.call(
		controllers.FramesIndex,
		nil,
		tuisupport.Open(".."),
	)
	if result != nil {
		fi.Items = result.Payload.([]*controllers.FramesIndexItemResult)
	}
	return
}

func (fi *FramesIndex) setTable() {
	propeties := []string{"Frame", "About"}
	data := []map[string]string{}
	for _, f := range fi.Items {
		data = append(data, map[string]string{
			"ID":    f.Frame,
			"Frame": f.Frame,
			"About": f.About,
		})
	}
	navigator := func(id string) (p string) {
		p = fmt.Sprintf("/frames/@%s", id)
		return
	}
	fi.Table = NewTable("frames", propeties, data, navigator)
}

func (fi *FramesIndex) focusChain() (fc []tuisupport.Focuser) {
	fc = []tuisupport.Focuser{fi.Table}
	return
}

func (fi *FramesIndex) isFocus() (is bool) {
	is = fi.Table.IsFocus
	return
}
