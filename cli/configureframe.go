package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/utils"
)

func configureFrame() (command any) {
	command = &cliapp.Command{
		Name:    "frame",
		Summary: "Configure settings for a frame",
		Aliases: ss("f"),
		Usage: ss(
			"sf configure frame [options] SETTINGS",
			"Configuration settings must be provided as the argument",
			"  Encode settings as YAML (accepts JSON)",
			"Provide an optional frame name using the -frame flag",
			"  Uses frame context when not provided",
		),
		Flags: ss(
			"string", "frame", "Name of the frame",
		),
		Handler:    configureFrameHandler,
		Controller: controllers.FrameConfigurationsUpdate,
		Viewer: cliapp.View(
			"frameconfigurations/update",
			"configurations/configuration",
			"configurations/settings",
			"configurations/setting",
		),
	}
	return
}

func configureFrameHandler(context *cliapp.Context) (
	params *controllers.Params,
	err error,
) {
	updates := make(map[string]string)
	yaml := context.Argument(0)

	utils.YamlUnmarshal([]byte(yaml), &updates)

	params = &controllers.Params{
		Payload: &controllers.FrameConfigurationsUpdateParams{
			Frame:   context.StringFlag("frame"),
			Updates: updates,
		},
	}
	return
}
