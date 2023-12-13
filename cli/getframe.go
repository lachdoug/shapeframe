package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func getFrame() (command any) {
	command = &cliapp.Command{
		Name:    "frame",
		Summary: "Get a frame",
		Aliases: ss("f"),
		Usage: ss(
			"sf get frame [options] [name]",
			"Provide an optional frame name as an argument",
			"  Uses frame context when not provided",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
		),
		Handler:    getFrameHandler,
		Controller: controllers.FramesRead,
		Viewer: cliapp.View(
			"frames/read",
			"frames/shapes",
			"configurations/configuration",
			"configurations/datum",
		),
	}
	return
}

func getFrameHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.FramesReadParams{
			Workspace: context.StringFlag("workspace"),
			Frame:     context.Argument(0),
		},
	}
	return
}
