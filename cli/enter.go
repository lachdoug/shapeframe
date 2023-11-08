package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func enter() (command any) {
	command = &cliapp.Command{
		Name:    "enter",
		Summary: "Enter shape, frame or workspace",
		Aliases: ss("en"),
		Usage: ss(
			"sf enter [options] [name]",
			"A name must be provided as an argument",
			"  When no context, the name must be a workspace name",
			"  When workspace context, the name must be a frame name",
			"  When frame context, the name must be a shape name",
		),
		Parametizer: enterParams,
		Controller:  controllers.ContextsUpdate,
		Viewer: cliapp.View(
			"contexts/update",
			"contexts/from",
			"contexts/to",
			"contexts/context",
		),
	}
	return
}

func enterParams(context *cliapp.Context) (jparams []byte, vn *app.Validation, err error) {
	name := context.Argument(0)
	params := map[string]any{}

	vn = &app.Validation{}
	if name == "" {
		vn.Add("Name", "must not be blank")
	}
	if vn.IsInvalid() {
		return
	}

	uc := models.ResolveUserContext(
		"Workspace", "Frame", "Shape",
	)
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
		err = app.ErrorWith(err, "has shape context so nothing to enter")
		return
	}

	jparams = jsonParams(params)
	return
}
