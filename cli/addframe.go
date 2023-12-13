package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func addFrame() (command any) {
	command = &cliapp.Command{
		Name:    "frame",
		Summary: "Add a frame to a workspace",
		Aliases: ss("f"),
		Usage: ss(
			"sf add frame [options] [framer]",
			"A framer name must be provided as an argument",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Uses framer name when not provided",
			"Provide an optional about using the -about flag",
			"  Uses framer about when not provided",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
			"string", "frame", "Name for the frame",
			"string", "about", "About the frame",
		),
		Handler:    addFrameHandler,
		Controller: controllers.FramesCreate,
		Viewer:     cliapp.View("frames/create"),
	}
	return
}

func addFrameHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.FramesCreateParams{
			Framer:    context.Argument(0),
			Workspace: context.StringFlag("workspace"),
			Frame:     context.StringFlag("frame"),
			About:     context.StringFlag("about"),
		},
	}
	return
}
