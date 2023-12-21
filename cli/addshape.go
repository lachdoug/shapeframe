package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func addShape() (command any) {
	command = &cliapp.Command{
		Name:    "shape",
		Summary: "Add a shape to a frame",
		Aliases: ss("s"),
		Usage: ss(
			"sf add shape [options] NAME",
			"A shaper name must be provided as an argument",
			"Provide an optional frame name using the -frame flag",
			"  Otherwise uses frame context when not provided",
			"Provide an optional shape name using the -shape flag",
			"  Uses shaper name when not provided",
			"Provide an optional about using the -about flag",
			"  Uses shaper about when not provided",
		),
		Flags: ss(
			"string", "frame", "Name of the frame",
			"string", "shape", "Name for the shape",
			"string", "about", "About the shape",
		),
		Handler:    addShapeHandler,
		Controller: controllers.ShapesCreate,
		Viewer:     cliapp.View("shapes/create"),
	}
	return
}

func addShapeHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.ShapesCreateParams{
			Shaper: context.Argument(0),
			Frame:  context.StringFlag("frame"),
			Shape:  context.StringFlag("shape"),
			About:  context.StringFlag("about"),
		},
	}
	return
}
