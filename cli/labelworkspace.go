package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func labelWorkspace() (command any) {
	command = &cliapp.Command{
		Name:    "workspace",
		Summary: "Change name and/or about for workspace",
		Aliases: ss("w"),
		Flags: ss(
			"string", "name", "New name",
			"string", "about", "New about",
		),
		Parametizer: labelWorkspaceParams,
		Controller:  controllers.WorkspacesUpdate,
		Viewer:      cliapp.View("workspaces/update"),
	}
	return
}

func labelWorkspaceParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	uc := models.UserContextNew()
	uc.Load("Workspace")

	w := uc.Workspace
	if w == nil {
		err = app.Error(nil, "no workspace context")
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Name":      context.StringFlag("name"),
		"About":     context.StringFlag("about"),
	})
	return
}
