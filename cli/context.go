package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func context() (command any) {
	command = &cliapp.Command{
		Name:    "context",
		Summary: "Report context",
		Aliases: ss("x"),
		Usage: ss(
			"sf context [options]",
		),
		Controller: controllers.ContextsRead,
		Viewer:     cliapp.View("contexts/read", "contexts/context"),
	}
	return
}
