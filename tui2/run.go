package tui2

import "sf/app/errors"

func Run() {
	if _, err := NewTApp(NewApp()).Run(); err != nil {
		errors.ErrorHandler(err)
	}
}
