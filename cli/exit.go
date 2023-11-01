package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func exit() (command any) {
	command = &cliapp.Command{
		Name:        "exit",
		Summary:     "Exit shape, frame or workspace",
		Aliases:     ss("ex"),
		Parametizer: exitParams,
		Controller:  controllers.ContextsUpdate,
		Viewer:      cliapp.View("contexts/update"),
	}
	return
}

func exitParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	params := map[string]any{}

	uc := models.UserContextNew()
	uc.Load("Workspace", "Frame", "Shape")

	if uc.Shape != nil {
		params["Workspace"] = uc.Workspace.Name
		params["Frame"] = uc.Frame.Name
	} else if uc.Frame != nil {
		params["Workspace"] = uc.Workspace.Name
	} else if uc.Workspace == nil {
		err = app.Error(err, "no context so nothing to exit")
		return
	}

	jparams = jsonParams(params)
	return
}
