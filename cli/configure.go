package cli

import "sf/cli/cliapp"

func configure() (commandset any) {
	commandset = &cliapp.CommandSet{
		Name:    "configure",
		Aliases: ss("c"),
		Summary: "Configure a shape, frame or workspace",
		Commands: cs(
			configureShape,
			configureFrame,
			// configureWorkspace,
		),
	}
	return
}
