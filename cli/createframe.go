package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func createFrame() (command any) {
	command = &cliapp.Command{
		Name:    "frame",
		Summary: "Create a frame",
		Usage:   ss("sf create frame [command options] [name]"),
		Description: ss(
			"A framer name must be provided as the first argument",
			"Provide an optional frame name using the --name flag",
			"Provide an optional about using the --about flag",
		),
		Aliases: ss("f"),
		Flags: ss(
			"string", "name", "Name for the frame",
			"string", "about", "About the frame",
		),
		Parametizer: createFrameParams,
		Controller:  controllers.FramesCreate,
		Viewer:      cliapp.View("frames/create"),
	}
	return
}

func createFrameParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	uc := models.UserContextNew()
	uc.Load("Workspace")
	w := uc.Workspace
	if w == nil {
		err = app.Error(nil, "no workspace context")
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Framer":    context.Argument(0),
		"Name":      context.StringFlag("name"),
		"About":     context.StringFlag("about"),
	})
	return
}
