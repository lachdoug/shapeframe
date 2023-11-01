package cli

import "sf/cli/cliapp"

func configure() (commandset any) {
	commandset = &cliapp.CommandSet{
		Name:    "configure",
		Summary: "Configure a shape, frame or workspace",
		Aliases: ss("co"),
		Commands: cs(
			configureShape,
			configureFrame,
			// configureWorkspace,
		),
	}
	return
}
