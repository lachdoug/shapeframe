package cli

import "sf/cli/cliapp"

func add() (commandset any) {
	commandset = &cliapp.CommandSet{
		Name:    "add",
		Summary: "Add a shape, frame, workspace, repository or directory",
		Aliases: ss("a"),
		Commands: cs(
			addShape,
			addFrame,
			addRepository,
			addDirectory,
		),
	}
	return
}
