package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func addWorkspace() (command any) {
	command = &cliapp.Command{
		Name:    "workspace",
		Summary: "Add a workspace",
		Aliases: ss("w"),
		Usage: ss(
			"sf add workspace [options] [workspace]",
			"A workspace name must be provided as an argument",
			"Provide an optional about using the -about flag",
		),
		Flags: ss(
			"string", "about", "About the workspace",
		),
		Handler:    addWorkspaceHandler,
		Controller: controllers.WorkspacesCreate,
		Viewer:     cliapp.View("workspaces/create"),
	}
	return
}

func addWorkspaceHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.WorkspacesCreateParams{
			Workspace: context.Argument(0),
			About:     context.StringFlag("about"),
		},
	}
	return
}
