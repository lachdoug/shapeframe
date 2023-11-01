package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func addDirectory() (command any) {
	command = &cliapp.Command{
		Name:        "directory",
		Summary:     "Add a directory to workspace",
		Usage:       ss("sf add directory [command options] [URI]"),
		Aliases:     ss("d"),
		Parametizer: addDirectoryParams,
		Controller:  controllers.DirectoriesCreate,
		Viewer:      cliapp.View("directories/create"),
	}
	return
}

func addDirectoryParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
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
