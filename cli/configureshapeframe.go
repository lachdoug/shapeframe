package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/utils"
)

func configureShapeFrame() (command any) {
	command = &cliapp.Command{
		Name:    "shape-frame",
		Summary: "Configure shape-frame settings for a shape",
		Aliases: ss("s-f"),
		Usage: ss(
			"sf configure shape-frame [options] SETTINGS",
			"Configuration settings must be provided as the argument",
			"  Encode settings as YAML (accepts JSON)",
			"Provide an optional frame name using the -frame flag",
			"  Otherwise uses frame context when not provided",
			"Provide an optional shape name using the -shape flag",
			"  Is required if -frame flag is set",
			"  Uses shape context when not provided",
		),
		Flags: ss(
			"string", "frame", "Name of the frame",
			"string", "shape", "Name of the shape",
		),
		Handler:    configureShapeFrameHandler,
		Controller: controllers.ShapeFrameConfigurationsUpdate,
		Viewer: cliapp.View(
			"shapeframeconfigurations/update",
			"configurations/configuration",
			"configurations/settings",
			"configurations/setting",
		),
	}
	return
}

func configureShapeFrameHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	var updates map[string]string
	yaml := context.Argument(0)

	utils.YamlUnmarshal([]byte(yaml), &updates)

	params = &controllers.Params{
		Payload: &controllers.ShapeFrameConfigurationsUpdateParams{
			Frame:   context.StringFlag("frame"),
			Shape:   context.StringFlag("shape"),
			Updates: updates,
		},
	}
	return
}
