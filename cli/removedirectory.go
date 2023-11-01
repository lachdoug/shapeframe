package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func removeDirectory() (command any) {
	command = &cliapp.Command{
		Name:        "directory",
		Summary:     "Remove a directory from the workspace context",
		Usage:       ss("sf remove directory [URI]"),
		Aliases:     ss("d"),
		Parametizer: removeDirectoryParams,
		Controller:  controllers.DirectoriesDestroy,
		Viewer:      cliapp.View("directories/destroy"),
	}
	return
}

func removeDirectoryParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	uc := models.UserContextNew()
	uc.Load("Workspace")
	w := uc.Workspace
	if w == nil {
		err = app.Error(nil, "no workspace context")
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Path":      context.Argument(0),
	})
	return
}
