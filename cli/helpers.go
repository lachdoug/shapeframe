package cli

import (
	"encoding/json"
	"sf/app"
	"sf/utils"

	"github.com/fatih/color"
	"github.com/peterh/liner"
)

// Marshal a params map as JSON
func jsonParams(m map[string]any) (j []byte) {
	j, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return
}

// Take a result item and typecast it as a map
func resultItem(body map[string]any) (item map[string]any) {
	item = body["Result"].(map[string]any)
	return
}

// Take a result collection and typecast items as maps
func resultItems(body map[string]any) (items []map[string]any) {
	ris := body["Result"].([]any)
	for _, ri := range ris {
		items = append(items, ri.(map[string]any))
	}
	return
}

func stream(body map[string]any) (err error) {
	s := utils.StreamLoad(body["Stream"].(string))
	hideCursor()
	defer showCursor()
	if err = s.ReadOut(app.Out, app.Err); err != nil {
		err = app.ErrorWrap(err)
		return
	}
	return
}

// Render a question and prompt for an answer
func prompt(question string, opts ...string) (answer string, err error) {
	line := liner.NewLiner()
	defer func() {
		if err = line.Close(); err != nil {
			panic(err)
		}
	}()
	line.SetCtrlCAborts(true)

	suggest := ""
	if len(opts) > 0 {
		suggest = opts[0]
	}

	answer, err = line.PromptWithSuggestion(question+" ", suggest, -1)
	return
}

// Add green color codes to text
func green(text string) (green string) {
	green = color.New(color.FgGreen).Sprint(text)
	return
}

// Hide terminal cursor
func hideCursor() {
	app.Print("\033[?25l")
}

// Show terminal cursor
func showCursor() {
	app.Print("\033[?25h")
}

// Slice of strings
func ss(s ...string) []string {
	return s
}

// Slice of commands
func cs(c ...func() any) []func() any {
	return c
}

// Slice of table value functions
func tvs(c ...func(any) string) []func(any) string {
	return c
}

// Slice of table accent functions
func tas(c ...func(string, map[string]any) string) []func(string, map[string]any) string {
	return c
}
