package cliapp

import (
	"sf/app"
)

type Router struct {
	Name        string
	Summary     string
	Usage       string
	Description []string
	Flags       []string
	Commands    []func() any
}

func (r *Router) Run(args []string) {
	if err := uApp(r).Run(args); err != nil {
		app.ErrorHandler(err)
	}
}
