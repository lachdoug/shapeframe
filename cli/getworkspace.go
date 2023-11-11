package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
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
		Parametizer: getWorkspaceParams,
		Controller:  controllers.WorkspacesRead,
		Viewer: cliapp.View(
			"workspaces/read",
			"workspaces/frames",
			"workspaces/repositories",
			"workspaces/directories",
		),
	}
	return
}

func getWorkspaceParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	workspace := context.Argument(0)

	uc := models.ResolveUserContext(
		"Workspace",
		"Workspaces",
	)
	if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
		return
	}

	params := map[string]any{
		"Workspace": w.Name,
	}

	jparams = jsonParams(params)
	return
}
