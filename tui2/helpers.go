package tui2

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func viewErrors(errmsgs []string) (v string) {
	errorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	lines := []string{}
	lines = append(lines, errmsgs...)
	v = errorStyle.Render(lipgloss.JoinVertical(lipgloss.Top, lines...))
	return
}

var buttonStyle lipgloss.Style = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	PaddingLeft(1).PaddingRight(1)

type ViewBuilder struct {
	Buffer *bytes.Buffer
}

func NewViewBuilder() (vb *ViewBuilder) {
	vb = &ViewBuilder{
		Buffer: &bytes.Buffer{},
	}
	return
}

func (vb *ViewBuilder) Writeln(s ...any) {
	vb.Write(fmt.Sprintln(s...))
}

func (vb *ViewBuilder) Write(s string) {
	vb.Buffer.WriteString(s)
}

func (vb *ViewBuilder) String() (s string) {
	s = vb.Buffer.String()
	return
}

// Marshal a params map as JSON
func jsonParams(m map[string]any) (j []byte) {
	var err error
	if j, err = json.Marshal(m); err != nil {
		panic(err)
	}
	return
}

func jsonResult(j []byte) (m map[string]any) {
	var err error
	m = map[string]any{}
	if err = json.Unmarshal(j, &m); err != nil {
		panic(err)
	}
	return
}

// Take a result item and typecast it as a map
func resultItem(j []byte) (item map[string]any) {
	body := jsonResult(j)
	item = body["Result"].(map[string]any)
	return
}

// Take a result collection and typecast items as maps
func resultItems(j []byte) (items []map[string]any) {
	body := jsonResult(j)
	ris := body["Result"].([]any)
	for _, ri := range ris {
		items = append(items, ri.(map[string]any))
	}
	return
}

// func stream(body map[string]any) (err error) {
// 	s := utils.StreamLoad(body["Stream"].(string))
// 	hideCursor()
// 	defer showCursor()
// 	err = s.ReadOut()
// 	return
// }

// // Add green color codes to text
// func green(text string) (green string) {
// 	green = color.New(color.FgGreen).Sprint(text)
// 	return
// }

// // Hide terminal cursor
// func hideCursor() {
// 	app.Print(app.HideCursor)
// }

// // Show terminal cursor
// func showCursor() {
// 	app.Print(app.ShowCursor)
// }

// // Slice of strings
// func ss(s ...string) []string {
// 	return s
// }

// // Slice of commands
// func cs(c ...func() any) []func() any {
// 	return c
// }

// // Slice of table value functions
// func tvs(c ...func(any) string) []func(any) string {
// 	return c
// }

// // Slice of table accent functions
// func tas(c ...func(string, map[string]any) string) []func(string, map[string]any) string {
// 	return c
// }
