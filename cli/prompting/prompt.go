package prompting

import "github.com/peterh/liner"

// Render a question and prompt for an answer
func Prompt(question string, opts ...string) (answer string, err error) {
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
