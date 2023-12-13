package cliapp

import "sf/app/errors"

type Router struct {
	Name        string
	Summary     string
	Usage       string
	Description []string
	Flags       []string
	Commands    []func() any
}

func (r *Router) Run(args []string) {
	if err := NewUApp(r).Run(args); err != nil {
		errors.ErrorHandler(err)
	}
}
