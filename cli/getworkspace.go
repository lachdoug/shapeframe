package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func getWorkspace() (command any) {
	command = &cliapp.Command{
		Name:    "inspect",
		Summary: "Inspect workspace",
		Aliases: ss("w"),
		Usage: ss(
			"sf inspect workspace [options] [name]",
			"Provide an optional workspace name as an argument",
			"  Uses workspace context when not provided",
		),
		Handler:    getWorkspaceHandler,
		Controller: controllers.WorkspacesRead,
		Viewer: cliapp.View(
			"workspaces/read",
			"workspaces/frames",
			"workspaces/repositories",
			"workspaces/directories",
		),
	}
	return
}

func getWorkspaceHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.WorkspacesReadParams{
			Workspace: context.Argument(0),
		},
	}
	return
}
