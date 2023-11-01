package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func enter() (command any) {
	command = &cliapp.Command{
		Name:        "enter",
		Summary:     "Enter shape, frame or workspace",
		Aliases:     ss("en"),
		Parametizer: enterParams,
		Controller:  controllers.ContextsUpdate,
		Viewer:      cliapp.View("contexts/update"),
	}
	return
}

func enterParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	name := context.Argument(0)
	validation = &app.Validation{}
	if name == "" {
		validation.Add("Name", "must not be blank")
	}
	if validation.IsInvalid() {
		return
	}

	params := map[string]any{}

	uc := models.UserContextNew()
	uc.Load("Workspace", "Frame", "Shape")

	if uc.Workspace == nil {
		params["Workspace"] = name
	} else if uc.Frame == nil {
		params["Workspace"] = uc.Workspace.Name
		params["Frame"] = name
	} else if uc.Shape == nil {
		params["Workspace"] = uc.Workspace.Name
		params["Frame"] = uc.Frame.Name
		params["Shape"] = name
	} else {
		err = app.Error(err, "has shape context so nothing to enter")
		return
	}

	jparams = jsonParams(params)
	return
}
