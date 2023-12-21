package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/utils"
)

func configureShape() (command any) {
	command = &cliapp.Command{
		Name:    "shape",
		Summary: "Configure settings for a shape",
		Aliases: ss("s"),
		Usage: ss(
			"sf configure shape [options] SETTINGS",
			"Configuration settings must be provided as the argument",
			"  Encode settings as YAML (accepts JSON)",
			"Provide an optional frame name using the -frame flag",
			"  Is required if -workspace flag is set",
			"  Otherwise uses frame context when not provided",
			"Provide an optional shape name using the -shape flag",
			"  Is required if -frame flag is set",
			"  Uses shape context when not provided",
			"Provide an optional password for secrets encryption using the -password flag",
			"  Prompt when when not provided",
		),
		Flags: ss(
			"string", "frame", "Name of the frame",
			"string", "shape", "Name of the shape",
			"string", "password", "Password for secrets encryption key",
		),
		Handler:    configureShapeParams,
		Controller: controllers.ShapeConfigurationsUpdate,
		Viewer: cliapp.View(
			"shapeconfigurations/update",
			"configurations/configuration",
			"configurations/settings",
			"configurations/setting",
		),
	}
	return
}

func configureShapeParams(context *cliapp.Context) (params *controllers.Params, err error) {
	var updates map[string]string
	yaml := context.Argument(0)

	utils.YamlUnmarshal([]byte(yaml), &updates)

	params = &controllers.Params{
		Payload: &controllers.ShapeConfigurationsUpdateParams{
			Frame:   context.StringFlag("frame"),
			Shape:   context.StringFlag("shape"),
			Updates: updates,
		},
	}
	return
}
