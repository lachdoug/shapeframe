package cliapp

import "sf/app"

type Command struct {
	Name        string
	Summary     string
	Usage       []string
	Description []string
	Aliases     []string
	Flags       []string
	Parametizer func(*Context) ([]byte, *app.Validation, error)
	Controller  func([]byte) ([]byte, *app.Validation, error)
	Viewer      func(map[string]any) (string, error)
}
