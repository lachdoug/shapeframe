package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func destroyShape() (command any) {
	command = &cliapp.Command{
		Name:        "shape",
		Summary:     "Destroy a shape",
		Usage:       ss("sf destroy shape [command options] [name]"),
		Aliases:     ss("s"),
		Parametizer: destroyShapeParams,
		Controller:  controllers.ShapesDestroy,
		Viewer:      cliapp.View("shapes/destroy"),
	}
	return
}

func destroyShapeParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	uc := models.UserContextNew()
	uc.Load("Workspace", "Frame")
	w := uc.Workspace
	if w == nil {
		err = app.Error(nil, "no workspace context")
		return
	}
	f := uc.Frame
	if f == nil {
		err = app.Error(nil, "no frame context")
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Frame":     f.Name,
		"Name":      context.Argument(0),
	})
	return
}
