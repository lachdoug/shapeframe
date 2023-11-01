package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func labelShape() (command any) {
	command = &cliapp.Command{
		Name:    "shape",
		Summary: "Change name and/or about for shape",
		Aliases: ss("s"),
		Flags: ss(
			"string", "name", "New name",
			"string", "about", "New about",
		),
		Parametizer: labelShapeParams,
		Controller:  controllers.ShapesUpdate,
		Viewer:      cliapp.View("shapes/update"),
	}
	return
}

func labelShapeParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	uc := models.UserContextNew()
	uc.Load("Workspace", "Frame", "Shape")

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
	s := uc.Shape
	if s == nil {
		err = app.Error(nil, "no shape context")
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Frame":     f.Name,
		"Shape":     s.Name,
		"Name":      context.StringFlag("name"),
		"About":     context.StringFlag("about"),
	})
	return
}
