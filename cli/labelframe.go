package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func labelFrame() (command any) {
	command = &cliapp.Command{
		Name:    "frame",
		Summary: "Change name and/or about for frame",
		Aliases: ss("f"),
		Usage: ss(
			"sf label frame [options] [name]",
			"Provide an optional frame name as an argument",
			"  Uses frame context when not provided",
			"Provide an optional frame update name using the -name flag",
			"Provide an optional frame update about using the -about flag",
		),
		Flags: ss(
			"string", "name", "New name for the frame",
			"string", "about", "New about for the frame",
		),
		Handler:    labelFrameHandler,
		Controller: controllers.FramesUpdate,
		Viewer:     cliapp.View("frames/update", "labels/label"),
	}
	return
}

func labelFrameHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	updates := map[string]string{}
	if context.IsSet("name") {
		updates["Name"] = context.StringFlag("name")
	}
	if context.IsSet("about") {
		updates["About"] = context.StringFlag("about")
	}

	params = &controllers.Params{
		Payload: &controllers.FramesUpdateParams{
			Frame:   context.Argument(0),
			Updates: updates,
		},
	}
	return
}
