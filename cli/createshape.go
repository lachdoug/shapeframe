package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
)

func createShape() (command any) {
	command = &cliapp.Command{
		Name:    "shape",
		Summary: "Create a shape",
		Usage:   ss("sf create shape [command options] [name]"),
		Description: ss(
			"A shaper name must be provided as the first argument",
			"Provide an optional shape name using the --name flag",
			"Provide an optional about using the --about flag",
		),
		Aliases: ss("s"),
		Flags: ss(
			"string", "name", "Name for the shape",
			"string", "about", "About the shape",
		),
		Parametizer: createShapeParams,
		Controller:  controllers.ShapesCreate,
		Viewer:      cliapp.View("shapes/create"),
	}
	return
}

func createShapeParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	jparams = jsonParams(map[string]any{
		"Shaper": context.Argument(0),
		"Name":   context.StringFlag("name"),
		"About":  context.StringFlag("about"),
	})
	return
}
