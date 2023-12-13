package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func removeWorkspace() (command any) {
	command = &cliapp.Command{
		Name:    "workspace",
		Summary: "Remove a workspace",
		Aliases: ss("w"),
		Usage: ss(
			"sf remove workspace [options] [workspace]",
			"Provide an optional workspace name as an argument",
			"  Uses workspace context when not provided",
		),
		Handler:    removeWorkspaceHandler,
		Controller: controllers.WorkspacesDelete,
		Viewer:     cliapp.View("workspaces/delete"),
	}
	return
}

func removeWorkspaceHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.WorkspacesDeleteParams{
			Workspace: context.Argument(0),
		},
	}
	return
}
