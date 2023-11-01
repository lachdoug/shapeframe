package cli

import "sf/cli/cliapp"

func destroy() (commandset any) {
	commandset = &cliapp.CommandSet{
		Name:    "destroy",
		Summary: "Destroy a shape, frame or workspace",
		Aliases: ss("de"),
		Commands: cs(
			destroyShape,
			destroyFrame,
			destroyWorkspace,
		),
	}
	return
}
