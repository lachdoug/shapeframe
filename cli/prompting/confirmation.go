package prompting

import "strings"

func Confirmation(question string) (confirmed bool, err error) {
	s, err := Prompt(question + " (Y/n)")
	if err != nil {
		return
	}
	if answer := strings.TrimSpace(s); answer == "Y" {
		confirmed = true
	}
	return
}
