package prompting

import (
	"sf/app/errors"

	"github.com/peterh/liner"
)

// Render a question and prompt for an answer
func Prompt(question string, opts ...string) (answer string, err error) {
	line := liner.NewLiner()
	defer func() {
		// Needs err scoped to defer func, otherwise prompt err is cleared
		var err error
		if err = line.Close(); err != nil {
			panic(err)
		}
	}()
	line.SetCtrlCAborts(true)

	suggest := ""
	if len(opts) > 0 {
		suggest = opts[0]
	}

	if answer, err = line.PromptWithSuggestion(question+"? ", suggest, -1); err != nil {
		err = errors.ErrorWrap(err)
		return
	}
	return
}
