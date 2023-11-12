package cli

import (
	"sf/cli/cliapp"
)

func get() (commandset any) {
	commandset = &cliapp.CommandSet{
		Name:    "get",
		Aliases: ss("g"),
		Summary: "Get a shape, frame, workspace, shaper or framer",
		Commands: cs(
			getShape,
			getFrame,
			getWorkspace,
			getRepository,
		),
	}
	return
}
