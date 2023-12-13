package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func removeShape() (command any) {
	command = &cliapp.Command{
		Name:    "shape",
		Summary: "Remove a shape from a frame",
		Aliases: ss("s"),
		Usage: ss(
			"sf remove shape [options] [shape]",
			"Provide an optional shape name as an argument",
			"  Uses shape context when not provided",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Is required if -workspace flag is set",
			"  Otherwise uses frame context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
			"string", "frame", "Name of the frame",
		),
		Handler:    removeShapeHandler,
		Controller: controllers.ShapesDelete,
		Viewer:     cliapp.View("shapes/delete"),
	}
	return
}

func removeShapeHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.ShapesDeleteParams{
			Workspace: context.StringFlag("workspace"),
			Frame:     context.StringFlag("frame"),
			Shape:     context.Argument(0),
		},
	}
	return
}
