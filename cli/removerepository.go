package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func removeRepository() (command any) {
	command = &cliapp.Command{
		Name:    "repository",
		Summary: "Remove a repository from a workspace",
		Aliases: ss("r"),
		Usage: ss(
			"sf remove repository [options] [URI]",
			"A repository URI must be provided as an argument",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
		),
		Parametizer: removeRepositoryParams,
		Controller:  controllers.RepositoriesDestroy,
		Viewer:      cliapp.View("repositories/destroy"),
	}
	return
}

func removeRepositoryParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	uri := context.Argument(0)
	workspace := context.StringFlag("workspace")

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace",
	)
	if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"URI":       uri,
	})
	return
}
