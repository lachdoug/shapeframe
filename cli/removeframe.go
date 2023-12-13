package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func removeFrame() (command any) {
	command = &cliapp.Command{
		Name:    "frame",
		Summary: "Remove a frame from a workspace",
		Aliases: ss("f"),
		Usage: ss(
			"sf remove frame [options] [frame]",
			"Provide an optional frame name as an argument",
			"  Uses frame context when not provided",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
		),
		Handler:    removeFrameHandler,
		Controller: controllers.FramesDelete,
		Viewer:     cliapp.View("frames/delete"),
	}
	return
}

func removeFrameHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.FramesDeleteParams{
			Workspace: context.StringFlag("workspace"),
			Frame:     context.Argument(0),
		},
	}
	return
}
