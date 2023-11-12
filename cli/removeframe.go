package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func removeFrame() (command any) {
	command = &cliapp.Command{
		Name:    "frame",
		Summary: "Remove a frame from a workspace",
		Aliases: ss("f"),
		Usage: ss(
			"sf remove frame [options] [name]",
			"Provide an optional frame name as an argument",
			"  Uses frame context when not provided",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
		),
		Parametizer: removeFrameParams,
		Controller:  controllers.FramesDestroy,
		Viewer:      cliapp.View("frames/destroy"),
	}
	return
}

func removeFrameParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	frame := context.Argument(0)
	workspace := context.StringFlag("workspace")

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace", "Frame",
	)
	if w, err = models.ResolveWorkspace(uc, workspace,
		"Frames",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, frame); err != nil {
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Frame":     f.Name,
	})
	return
}
