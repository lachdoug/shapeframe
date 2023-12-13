package cli

import (
	"fmt"
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
			"sf configure frame [options] [settings]",
			"Configuration settings must be provided as the argument",
			"  Encode settings as YAML (accepts JSON)",
			"Provide an optional frame name using the -frame flag",
			"  Uses frame context when not provided",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
		),
		Flags: ss(
			"string", "frame", "Name of the frame",
			"string", "workspace", "Name of the workspace",
		),
		Handler:    configureFrameHandler,
		Controller: controllers.FrameConfigurationsUpdate,
		Viewer: cliapp.View(
			"frameconfigurations/update",
			"configurations/configuration",
			"configurations/datum",
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

	fmt.Println("yaml, &updates", yaml, updates)

	params = &controllers.Params{
		Payload: &controllers.FrameConfigurationsUpdateParams{
			Workspace: context.StringFlag("workspace"),
			Frame:     context.StringFlag("frame"),
			Updates:   updates,
		},
	}
	return
}
