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
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Uses frame context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
			"string", "frame", "Frame name",
		),
		Handler:    getShapeHandler,
		Controller: controllers.ShapesRead,
		Viewer: cliapp.View(
			"shapes/read",
			"configurations/configuration",
			"configurations/datum",
		),
	}
	return
}

func getShapeHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.ShapesReadParams{
			Workspace: context.StringFlag("workspace"),
			Frame:     context.StringFlag("frame"),
			Shape:     context.Argument(0),
		},
	}
	return
}
