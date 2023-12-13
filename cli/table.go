package cli

import (
	"fmt"
	"sf/app/logs"
	"sf/utils"
	"strings"
)

type Table struct {
	Items   []map[string]any
	Titles  []string
	Keys    []string
	Accents []func(string, map[string]any) string
}

func (t *Table) tableValues() (rows [][]string) {
	for i := range t.Items {
		rows = append(rows, t.rowValues(i))
	}
	return
}

func (t *Table) rowValues(i int) (row []string) {
	for j := range t.Keys {
		row = append(row, t.cellValue(i, j))
	}
	return
}

func (t *Table) cellValue(i int, j int) (value string) {
	item := t.Items[i]
	value = item[t.Keys[j]].(string)
	return
}

func (t *Table) rows(tableValues [][]string, lengths []int) (rows []string) {
	for i, item := range t.Items {
		rows = append(rows, t.row(item, tableValues[i], lengths))
	}
	return
}

func (t *Table) row(item map[string]any, rowValues []string, lengths []int) (row string) {
	cells := []string{}
	for i := range t.Keys {
		cells = append(cells, t.cell(item, i, rowValues[i], lengths[i]))
	}
	row = strings.Join(cells, " ")
	return
}

func (t *Table) cell(item map[string]any, i int, value string, length int) (cell string) {
	format := fmt.Sprintf("%%-%ds", length)
	// if len(value) > length {
	// 	value = utils.TruncateString(value, length)
	// }
	cell = t.Accents[i](fmt.Sprintf(format, value), item)
	return
}

func (t *Table) lengths(tableValues [][]string) (lengths []int) {
	// terminalWidth, _ := utils.TerminalSize()
	columnValues := []string{}
	// occupiedChars := 0
	for i := range t.Keys {
		title := t.Titles[i]
		for j := range t.Items {
			columnValues = append(columnValues, tableValues[j][i])
		}
		length := t.longest(title, columnValues)
		// if i < len(t.Keys)-1 && length > 25 {
		// 	length = 25
		// }
		// if length > terminalWidth-occupiedChars-1 || i == len(t.Keys)-1 {
		// 	length = terminalWidth - occupiedChars - 1
		// }
		// if length < 0 {
		// 	length = 0
		// }
		// occupiedChars = occupiedChars + length + 1
		lengths = append(lengths, length)
	}
	logs.Log("TABLE LENGTHS", lengths)
	return
}

func (t *Table) longest(title string, column []string) (max int) {
	max = utils.LongestString(append(column, title))
	return
}

func (t *Table) headingRow(lengths []int) (heading string) {
	cells := []string{}
	for i, title := range t.Titles {
		cells = append(cells, t.headingCell(title, lengths[i]))
	}
	heading = strings.Join(cells, " ")
	return
}

func (t *Table) headingCell(title string, length int) (val string) {
	format := fmt.Sprintf("%%-%ds", length)
	val = fmt.Sprintf(format, title)
	return
}

func (t *Table) generate() (table map[string]any) {
	table = map[string]any{}
	tableValues := t.tableValues()
	lengths := t.lengths(tableValues)
	lines := []string{t.headingRow(lengths)}
	table["Lines"] = append(lines, t.rows(tableValues, lengths)...)
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
