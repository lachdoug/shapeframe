package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func join() (command any) {
	command = &cliapp.Command{
		Name:    "join",
		Summary: "Join a child frame to parent frame",
		Aliases: ss("j"),
		Usage: ss(
			"sf join [options] [URI]",
			"A frame name for the parent frame must be provided as an argument",
			"Provide an optional frame name for the child frame using the -frame flag",
			"  Uses frame context when not provided",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
		),
		Flags: ss(
			"string", "frame", "Frame name",
			"string", "workspace", "Workspace name",
		),
		Parametizer: joinParams,
		Controller:  controllers.FrameJoinsCreate,
		Viewer:      cliapp.View("framejoins/create"),
	}
	return
}

func joinParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	var fp *models.Frame
	workspace := context.StringFlag("workspace")
	frame := context.StringFlag("frame")
	parent := context.Argument(0)

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
	if fp, err = models.ResolveFrame(uc, w, parent); err != nil {
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Frame":     f.Name,
		"Parent":    fp.Name,
	})
	return
}
