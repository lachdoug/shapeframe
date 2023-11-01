package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func inspect() (command any) {
	command = &cliapp.Command{
		Name:       "inspect",
		Summary:    "Inspect workspace",
		Aliases:    ss("i"),
		Controller: controllers.WorkspacesRead,
		Viewer:     cliapp.View("workspaces/read"),
	}
	return
}
