package tui

import (
	"fmt"
	"sf/utils"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	zone "github.com/lrstanley/bubblezone"
)

type Table struct {
	Properties    []string
	Data          []map[string]string
	Navigator     func(string) string
	Model         table.Model
	MouseLocation tea.MouseMsg
	Width         int
	// Height        int
	Lengths []int
	// PageSize      int
	Rows      []table.Row
	Cols      []table.Column
	IsFocus   bool
	Selection int
	ID        string
}

func NewTable(
	id string,
	properties []string,
	data []map[string]string,
	navigator func(string) string,
) (t *Table) {
	t = &Table{
		ID:         id,
		Properties: properties,
		Data:       data,
		Navigator:  navigator,
	}
	return
}

func (t *Table) Init() (c tea.Cmd) {
	return
}

func (t *Table) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	m = t
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if t.IsFocus {
				c = t.enter()
				return
			}
		}
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			t.MouseLocation = msg
			for i, row := range t.Model.GetVisibleRows() {
				id := row.Data["ID"].(string)
				if t.isMouseInRow(id) {
					t.Model = t.Model.WithHighlightedRow(i)
				}
			}
		case tea.MouseLeft:
			if t.isMouseInRow(t.selectedID()) {
				c = t.enter()
				return
			}
		}
	}
	t.Model, c = t.Model.Update(msg)
	t.Selection = t.Model.GetHighlightedRowIndex()
	return
}

func (t *Table) View() (v string) {
	style := lipgloss.NewStyle()
	if t.IsFocus {
		style = style.
			BorderForeground(lipgloss.Color("15"))
	} else {
		style = style.
			Foreground(lipgloss.Color("7")).
			BorderForeground(lipgloss.Color("8"))
	}
	v = t.Model.WithBaseStyle(style).View()
	return
}

func (t *Table) setLengths() {
	ls := []int{}
	width := 1
	for _, p := range t.Properties {
		ss := []string{p}
		for _, d := range t.Data {
			ss = append(ss, d[p])
		}
		l := utils.LongestString(ss)
		width = width + l + 1
		ls = append(ls, l)
	}
	if width < t.Width {
		last := len(t.Properties) - 1
		ls[last] = ls[last] + t.Width - width
	}
	t.Lengths = ls
}

func (t *Table) setRows() {
	rows := []table.Row{}
	for _, d := range t.Data {
		tr := map[string]any{
			"ID": d["ID"],
		}
		for i, p := range t.Properties {
			length := t.Lengths[i]
			value := utils.FixedLengthString(d[p], length)
			tr[p] = t.zoneMark(d["ID"], p, value)
		}
		rows = append(rows, table.NewRow(table.RowData(tr)))
	}
	t.Rows = rows
}

func (t *Table) setCols() {
	cols := []table.Column{}
	for i, p := range t.Properties {
		col := table.NewColumn(p, strings.ToUpper(p), t.Lengths[i]).
			WithStyle(lipgloss.NewStyle().Align(lipgloss.Left))
		cols = append(cols, col)
	}
	t.Cols = cols
}

// func (t *Table) setPageSize() {
// 	t.PageSize = t.Height - 6
// 	if t.PageSize < 1 {
// 		t.PageSize = 1
// 	}
// }

func (t *Table) setModel() {
	t.Model = table.New(t.Cols).
		WithRows(t.Rows).
		Border(t.border()).
		WithHighlightedRow(t.Selection).
		Focused(t.IsFocus).
		WithTargetWidth(t.Width).
		// WithPageSize(t.PageSize).
		HighlightStyle(lipgloss.NewStyle().Background(lipgloss.Color("0")))
}

func (t *Table) border() (b table.Border) {
	b = table.Border{
		Top:    "─",
		Left:   "│",
		Right:  "│",
		Bottom: "─",

		TopRight:    "╮",
		TopLeft:     "╭",
		BottomRight: "╯",
		BottomLeft:  "╰",

		TopJunction:    "┬",
		LeftJunction:   "├",
		RightJunction:  "┤",
		BottomJunction: "┴",
		InnerJunction:  "┼",

		InnerDivider: "│",
	}
	return
}

func (t *Table) zoneMark(id string, property string, value string) (z string) {
	z = zone.Mark(fmt.Sprintf("%s-%s-%s", t.ID, id, property), value)
	return
}

func (t *Table) setSize(w int, h int) {
	t.Width = w
	// t.Height = h
	// t.setPageSize()
	t.setLengths()
	t.setRows()
	t.setCols()
	t.setModel()
}

func (t *Table) enter() (c tea.Cmd) {
	c = Open(t.Navigator(t.selectedID()))
	return
}

func (t *Table) isMouseInZone(id string, property string) (is bool) {
	is = zone.Get(fmt.Sprintf("%s-%s-%s", t.ID, id, property)).InBounds(t.MouseLocation)
	return
}

func (t *Table) isMouseInRow(id string) (is bool) {
	for _, p := range t.Properties {
		if t.isMouseInZone(id, p) {
			is = true
			return
		}
	}
	return
}

func (t *Table) selectedID() (id string) {
	id = t.Model.HighlightedRow().Data["ID"].(string)
	return
}

func (t *Table) Focus(aspect string) (c tea.Cmd) {
	if len(t.Data) == 0 {
		if aspect == "next" {
			c = func() tea.Msg { return tea.KeyMsg{Type: tea.KeyTab} }
		} else {
			c = func() tea.Msg { return tea.KeyMsg{Type: tea.KeyShiftTab} }
		}
		return
	}
	t.IsFocus = true
	t.Model = t.Model.Focused(true)
	return
}

func (t *Table) Blur() {
	t.IsFocus = false
	t.Model = t.Model.Focused(false)
}
