package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func removeDirectory() (command any) {
	command = &cliapp.Command{
		Name:    "directory",
		Summary: "Remove a directory from a workspace",
		Aliases: ss("d"),
		Usage: ss(
			"sf remove directory [options] [path]",
			"An absolute or relative (to working directory) path must be provided as an argument",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
		),
		Parametizer: removeDirectoryParams,
		Controller:  controllers.DirectoriesDestroy,
		Viewer:      cliapp.View("directories/destroy"),
	}
	return
}

func removeDirectoryParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	path := context.Argument(0)
	workspace := context.StringFlag("workspace")

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace",
	)
	if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Path":      path,
	})
	return
}
