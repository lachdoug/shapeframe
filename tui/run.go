package tui

import "sf/app/errors"

func Run() {
	if _, err := newTApp(newApp()).Run(); err != nil {
		errors.ErrorHandler(err)
	}
}
