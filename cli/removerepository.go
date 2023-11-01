package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func removeRepository() (command any) {
	command = &cliapp.Command{
		Name:        "repository",
		Summary:     "Remove a repository from workspace",
		Usage:       ss("sf remove repository [URI]"),
		Aliases:     ss("r"),
		Parametizer: removeRepositoryParams,
		Controller:  controllers.RepositoriesDestroy,
		Viewer:      cliapp.View("repositories/destroy"),
	}
	return
}

func removeRepositoryParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	uc := models.UserContextNew()
	uc.Load("Workspace")
	w := uc.Workspace
	if w == nil {
		err = app.Error(nil, "no workspace context")
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"URI":       context.Argument(0),
	})
	return
}
