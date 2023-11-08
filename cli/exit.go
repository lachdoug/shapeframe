package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func exit() (command any) {
	command = &cliapp.Command{
		Name:    "exit",
		Summary: "Exit shape, frame or workspace",
		Aliases: ss("ex"),
		Usage: ss(
			"sf exit [options]",
			"When workspace context, the context will change to no context",
			"When frame context, the context will change to workspace context",
			"When shape context, the context will change to frame context",
		),
		Parametizer: exitParams,
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

func exitParams(context *cliapp.Context) (jparams []byte, vn *app.Validation, err error) {
	params := map[string]any{}

	uc := models.ResolveUserContext(
		"Workspace", "Frame", "Shape",
	)
	if uc.Shape != nil {
		params["Workspace"] = uc.Workspace.Name
		params["Frame"] = uc.Frame.Name
	} else if uc.Frame != nil {
		params["Workspace"] = uc.Workspace.Name
	} else if uc.Workspace == nil {
		err = app.ErrorWith(err, "no context so nothing to exit")
		return
	}

	jparams = jsonParams(params)
	return
}
