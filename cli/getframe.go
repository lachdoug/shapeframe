package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func getFrame() (command any) {
	command = &cliapp.Command{
		Name:    "frame",
		Summary: "Get a frame",
		Aliases: ss("f"),
		Usage: ss(
			"sf get frame [options] [name]",
			"Provide an optional frame name as an argument",
			"  Uses frame context when not provided",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
		),
		Parametizer: getFrameParams,
		Controller:  controllers.FramesRead,
		Viewer: cliapp.View(
			"frames/read",
			"frames/shapes",
			"configurations/configuration",
			"configurations/setting",
		),
	}
	return
}

func getFrameParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	frame := context.Argument(0)
	workspace := context.StringFlag("workspace")

	uc := models.ResolveUserContext(
		"Frame",
		"Workspace.Frames",
		"Workspaces.Frames",
	)
	if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, frame, "Configuration"); err != nil {
		return
	}

	params := map[string]any{
		"Workspace": w.Name,
		"Frame":     f.Name,
	}

	jparams = jsonParams(params)
	return
}
