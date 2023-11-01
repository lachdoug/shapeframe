package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func context() (command any) {
	command = &cliapp.Command{
		Name:       "context",
		Summary:    "Report context",
		Aliases:    ss("?"),
		Controller: controllers.ContextsRead,
		Viewer:     cliapp.View("contexts/read"),
	}
	return
}
