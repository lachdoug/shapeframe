package cli

import "sf/cli/cliapp"

func list() (commandset any) {
	commandset = &cliapp.CommandSet{
		Name:    "list",
		Summary: "List shapes, frames, shapers, framers or workspaces",
		Aliases: ss("ls"),
		Commands: cs(
			listFrames,
			listShapes,
			listFramers,
			listShapers,
		),
	}
	return
}
