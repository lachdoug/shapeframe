package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/utils"
)

func configureFrameShape() (command any) {
	command = &cliapp.Command{
		Name:    "frame-shape",
		Summary: "Configure frame settings for a shape",
		Aliases: ss("f-s"),
		Usage: ss(
			"sf configure frameshape [options] [value1] [value2] [value3]...",
			"Configuration settings must be provided as the argument",
			"  Encode settings as YAML (accepts JSON)",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Is required if -workspace flag is set",
			"  Otherwise uses frame context when not provided",
			"Provide an optional shape name using the -shape flag",
			"  Is required if -frame flag is set",
			"  Uses shape context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
			"string", "frame", "Name of the frame",
			"string", "shape", "Name of the shape",
		),
		Handler:    configureFrameShapeHandler,
		Controller: controllers.FrameShapeConfigurationsUpdate,
		Viewer: cliapp.View(
			"shapeframeconfigurations/update",
			"configurations/configuration",
			"configurations/datum",
		),
	}
	return
}

func configureFrameShapeHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	var updates map[string]string
	yaml := context.Argument(0)

	utils.YamlUnmarshal([]byte(yaml), &updates)

	params = &controllers.Params{
		Payload: &controllers.FrameShapeConfigurationsUpdateParams{
			Workspace: context.StringFlag("workspace"),
			Frame:     context.StringFlag("frame"),
			Shape:     context.StringFlag("shape"),
			Updates:   updates,
		},
	}
	return
}
