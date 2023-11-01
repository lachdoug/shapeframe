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
		Flags: ss(
			"bool", "confirm", "Do not propmt for confirmation",
		),
		Parametizer: nukeParams,
		Controller:  controllers.NukesCreate,
		Viewer:      cliapp.View("nukes/create"),
	}
	return
}

func nukeParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	if !context.BoolFlag("confirm") {
		nukePrompt()
	}
	return
}

func nukePrompt() {
	s := prompt("Are you sure that you want to destroy existing data? (Y/n)")
	if answer := strings.TrimSpace(s); answer != "Y" {
		os.Exit(0)
	}
}
