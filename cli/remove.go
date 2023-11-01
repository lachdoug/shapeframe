package cli

import "sf/cli/cliapp"

func remove() (commandset any) {
	commandset = &cliapp.CommandSet{
		Name:    "remove",
		Summary: "Remove a repository or directory from workspace",
		Aliases: ss("r"),
		Commands: cs(
			removeRepository,
			removeDirectory,
		),
	}
	return
}
