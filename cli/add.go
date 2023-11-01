package cli

import "sf/cli/cliapp"

func add() (commandset any) {
	commandset = &cliapp.CommandSet{
		Name:    "add",
		Summary: "Add a repository or directory to workspace",
		Aliases: ss("a"),
		Commands: cs(
			addRepository,
			addDirectory,
		),
	}
	return
}
