package cli

import "sf/cli/cliapp"

func create() (commandset any) {
	commandset = &cliapp.CommandSet{
		Name:    "create",
		Summary: "Create a shape, frame or workspace",
		Aliases: ss("cr"),
		Commands: cs(
			createShape,
			createFrame,
			createWorkspace,
		),
	}
	return
}
