package cli

import (
	"fmt"
	"sf/utils"
	"strings"
)

type Table struct {
	Items   []map[string]any
	Titles  []string
	Keys    []string
	Values  []func(any) string
	Accents []func(string, map[string]any) string
}

func (wt *Table) tableValues() (rows [][]string) {
	for i := range wt.Items {
		rows = append(rows, wt.rowValues(i))
	}
	return
}

func (wt *Table) rowValues(i int) (row []string) {
	for j := range wt.Keys {
		row = append(row, wt.cellValue(i, j))
	}
	return
}

func (wt *Table) cellValue(i int, j int) (value string) {
	item := wt.Items[i]
	value = wt.Values[j](item[wt.Keys[j]])
	return
}

func (wt *Table) rows(tableValues [][]string, lengths []int) (rows []string) {
	for i, item := range wt.Items {
		rows = append(rows, wt.row(item, tableValues[i], lengths))
	}
	return
}

func (wt *Table) row(item map[string]any, rowValues []string, lengths []int) (row string) {
	cells := []string{}
	for i := range wt.Keys {
		cells = append(cells, wt.cell(item, i, rowValues[i], lengths[i]))
	}
	row = strings.Join(cells, " ")
	return
}

func (wt *Table) cell(item map[string]any, i int, value string, length int) (cell string) {
	format := fmt.Sprintf("%%-%ds", length)
	cell = wt.Accents[i](fmt.Sprintf(format, value), item)
	return
}

func (wt *Table) lengths(tableValues [][]string) (lengths []int) {
	column := []string{}
	for i := range wt.Keys {
		title := wt.Titles[i]
		for j := range wt.Items {
			column = append(column, tableValues[j][i])
		}
		lengths = append(lengths, wt.longest(title, column))
	}
	return
}

func (wt *Table) longest(title string, column []string) (max int) {
	max = len(title)
	for _, value := range column {
		if l := len(value); l > max {
			max = l
		}
	}
	return
}

func (wt *Table) headingRow(lengths []int) (heading string) {
	cells := []string{}
	for i, t := range wt.Titles {
		cells = append(cells, wt.headingCell(t, lengths[i]))
	}
	heading = strings.Join(cells, " ")
	return
}

func (wt *Table) headingCell(title string, length int) (val string) {
	format := fmt.Sprintf("%%-%ds", length)
	val = fmt.Sprintf(format, title)
	return
}

func (wt *Table) generate() (table map[string]any) {
	table = map[string]any{}
	tableValues := wt.tableValues()
	lengths := wt.lengths(tableValues)
	lines := []string{wt.headingRow(lengths)}
	lines = append(lines, wt.rows(tableValues, lengths)...)
	table["Lines"] = utils.TruncateLines(lines)
	return
}

var tableCellAsteriskIfTrueValueFn func(any) string = func(in any) (out string) {
	if in.(bool) {
		out = "*"
	}
	return
}

var tableCellStringValueFn func(any) string = func(in any) (out string) {
	out = in.(string)
	return
}

var tableCellNoAccentFn func(string, map[string]any) string = func(in string, item map[string]any) (out string) {
	out = in
	return
}

var tableCellGreenIfInContextAccentFn func(string, map[string]any) string = func(in string, item map[string]any) (out string) {
	if item["IsContext"].(bool) {
		out = green(in)
	} else {
		out = in
	}
	return
}
