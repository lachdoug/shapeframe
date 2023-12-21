package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func initialize() (command any) {
	command = &cliapp.Command{
		Name:    "initialize",
		Summary: "Ininitialize a workspace",
		Aliases: ss("init"),
		Usage: ss(
			"sf initialize [options]",
			"Provide an optional workspace name using the -name flag",
			"  Otherwise uses present working directory base",
			"Provide an optional workspace update about using the -about flag",
		),
		Flags: ss(
			"string", "name", "Name the workspace",
			"string", "about", "About the workspace",
		),
		Handler:    initializeHandler,
		Controller: controllers.WorkspacesCreate,
		Viewer: cliapp.View(
			"workspaces/create",
		),
	}
	return
}

func initializeHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.WorkspacesCreateParams{
			Name:  context.StringFlag("name"),
			About: context.StringFlag("about"),
		},
	}
	return
}
