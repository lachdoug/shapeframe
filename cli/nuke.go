package cli

import (
	"os"
	"sf/cli/cliapp"
	"sf/cli/prompting"
	"sf/controllers"
)

func nuke() (command any) {
	command = &cliapp.Command{
		Name:    "nuke",
		Summary: "Reset shapeframe environment (data will be lost)",
		Usage: ss(
			"sf nuke [options]",
			"Confirm nuke by setting the -confirm flag",
			"  Otherwise propmt for confirmation",
		),

		Flags: ss(
			"bool", "confirm", "Confirm nuke",
		),
		Handler:    nukeHandler,
		Controller: controllers.NukesCreate,
		Viewer:     cliapp.View("nukes/create"),
	}
	return
}

func nukeHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	var confirmed bool
	if !context.BoolFlag("confirm") {
		confirmed, err = prompting.Confirmation(
			"Are you sure that you want to destroy existing data?",
		)
		if !confirmed {
			os.Exit(0)
		}
	}
	return
}
