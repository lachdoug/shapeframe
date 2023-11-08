package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func removeWorkspace() (command any) {
	command = &cliapp.Command{
		Name:    "workspace",
		Summary: "Remove a workspace",
		Aliases: ss("w"),
		Usage: ss(
			"sf remove workspace [options] [name]",
			"Provide an optional workspace name as an argument",
			"  Uses workspace context when not provided",
		),
		Parametizer: removeWorkspaceParams,
		Controller:  controllers.WorkspacesDestroy,
		Viewer:      cliapp.View("workspaces/destroy"),
	}
	return
}

func removeWorkspaceParams(context *cliapp.Context) (jparams []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	workspace := context.Argument(0)

	uc := models.ResolveUserContext(
		"Workspace",
		"Workspaces",
	)
	if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
	})
	return
}
