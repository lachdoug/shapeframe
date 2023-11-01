package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func destroyFrame() (command any) {
	command = &cliapp.Command{
		Name:        "frame",
		Summary:     "Destroy a frame",
		Usage:       ss("sf destroy frame [command options] [name]"),
		Aliases:     ss("f"),
		Parametizer: destroyFrameParams,
		Controller:  controllers.FramesDestroy,
		Viewer:      cliapp.View("frames/destroy"),
	}
	return
}

func destroyFrameParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	uc := models.UserContextNew()
	uc.Load("Workspace")
	w := uc.Workspace
	if w == nil {
		err = app.Error(nil, "no workspace context")
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Name":      context.Argument(0),
	})
	return
}
