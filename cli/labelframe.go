package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func labelFrame() (command any) {
	command = &cliapp.Command{
		Name:    "frame",
		Summary: "Change name and/or about for frame",
		Aliases: ss("f"),
		Flags: ss(
			"string", "name", "New name",
			"string", "about", "New about",
		),
		Parametizer: labelFrameParams,
		Controller:  controllers.FramesUpdate,
		Viewer:      cliapp.View("frames/update"),
	}
	return
}

func labelFrameParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
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
		"Name":      context.StringFlag("name"),
		"About":     context.StringFlag("about"),
	})
	return
}
