package cli

import "sf/cli/cliapp"

func label() (commandset any) {
	commandset = &cliapp.CommandSet{
		Name:    "label",
		Summary: "Change name and/or about for shape, frame or workspace",
		Aliases: ss("ll"),
		Commands: cs(
			labelShape,
			labelFrame,
			labelWorkspace,
		),
	}
	return
}
