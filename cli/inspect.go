package cli

import (
	"sf/cli/cliapp"
)

func inspect() (commandset any) {
	commandset = &cliapp.CommandSet{
		Name:    "inspect",
		Aliases: ss("i"),
		Summary: "Inspect a shape, frame, workspace, shaper or framer",
		Commands: cs(
			// inspectShape,
			// inspectFrame,
			inspectWorkspace,
		),
	}
	return
}
