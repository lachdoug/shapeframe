package cli

import "sf/cli/cliapp"

func remove() (commandset any) {
	commandset = &cliapp.CommandSet{
		Name:    "remove",
		Summary: "Remove a shape, frame, workspace, repository or directory",
		Aliases: ss("rm"),
		Commands: cs(
			removeShape,
			removeFrame,
			removeWorkspace,
			removeRepository,
			removeDirectory,
		),
	}
	return
}
