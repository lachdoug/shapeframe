package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func getShape() (command any) {
	command = &cliapp.Command{
		Name:    "inspect",
		Summary: "Inspect shape",
		Aliases: ss("s"),
		Usage: ss(
			"sf inspect shape [options] [name]",
			"Provide an optional shape name as an argument",
			"  Uses shape context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Uses frame context when not provided",
		),
		Flags: ss(
			"string", "frame", "Frame name",
		),
		Handler:    getShapeHandler,
		Controller: controllers.ShapesRead,
		Viewer: cliapp.View(
			"shapes/read",
			"configurations/configuration",
			"configurations/settings",
			"configurations/setting",
		),
	}
	return
}

func getShapeHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.ShapesReadParams{
			Frame: context.StringFlag("frame"),
			Shape: context.Argument(0),
		},
	}
	return
}
