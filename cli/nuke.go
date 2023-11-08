package cli

import (
	"os"
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"strings"
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
		Parametizer: nukeParams,
		Controller:  controllers.NukesCreate,
		Viewer:      cliapp.View("nukes/create"),
	}
	return
}

func nukeParams(context *cliapp.Context) (jparams []byte, vn *app.Validation, err error) {
	if !context.BoolFlag("confirm") {
		err = nukePrompt()
	}
	return
}

func nukePrompt() (err error) {
	s, err := prompt("Are you sure that you want to destroy existing data? (Y/n)")
	if err != nil {
		return
	}
	if answer := strings.TrimSpace(s); answer != "Y" {
		os.Exit(0)
	}
	return
}
