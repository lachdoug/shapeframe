package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func labelShape() (command any) {
	command = &cliapp.Command{
		Name:    "shape",
		Summary: "Change name and/or about for shape",
		Aliases: ss("s"),
		Usage: ss(
			"sf label shape [options] [name]",
			"Provide an optional shape name as an argument",
			"  Uses shape context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Is required if -workspace flag is set",
			"  Uses workspace context when not provided",
			"Provide an optional shape update name using the -name flag",
			"Provide an optional shape update about using the -about flag",
		),
		Flags: ss(
			"string", "frame", "Name of the frame",
			"string", "name", "New name for the shape",
			"string", "about", "New about for the shape",
		),
		Handler:    labelShapeHandler,
		Controller: controllers.ShapesUpdate,
		Viewer:     cliapp.View("shapes/update", "labels/label"),
	}
	return
}

func labelShapeHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	updates := map[string]string{}
	if context.IsSet("name") {
		updates["Name"] = context.StringFlag("name")
	}
	if context.IsSet("about") {
		updates["About"] = context.StringFlag("about")
	}

	params = &controllers.Params{
		Payload: &controllers.ShapesUpdateParams{
			Frame:   context.StringFlag("frame"),
			Shape:   context.Argument(0),
			Updates: updates,
		},
	}
	return
}
