package cliapp

import "sf/controllers"

type Command struct {
	Name        string
	Summary     string
	Usage       []string
	Description []string
	Aliases     []string
	Flags       []string
	Handler     func(*Context) (*controllers.Params, error)
	Controller  func(*controllers.Params) (*controllers.Result, error)
	Viewer      func(*controllers.Result) (string, error)
}
