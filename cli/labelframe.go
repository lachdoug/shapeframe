package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func labelFrame() (command any) {
	command = &cliapp.Command{
		Name:    "frame",
		Summary: "Change name and/or about for frame",
		Aliases: ss("f"),
		Usage: ss(
			"sf label frame [options] [name]",
			"Provide an optional frame name as an argument",
			"  Uses frame context when not provided",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame update name using the -name flag",
			"Provide an optional frame update about using the -about flag",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
			"string", "name", "New name for the frame",
			"string", "about", "New about for the frame",
		),
		Parametizer: labelFrameParams,
		Controller:  controllers.FramesUpdate,
		Viewer:      cliapp.View("frames/update", "labels/label"),
	}
	return
}

func labelFrameParams(context *cliapp.Context) (jparams []byte, err error) {
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
	if f, err = models.ResolveFrame(uc, w, frame); err != nil {
		return
	}

	update := map[string]any{}
	if context.IsSet("name") {
		update["Name"] = context.StringFlag("name")
	}
	if context.IsSet("about") {
		update["About"] = context.StringFlag("about")
	}

	params := map[string]any{
		"Workspace": w.Name,
		"Frame":     f.Name,
		"Update":    update,
	}

	jparams = jsonParams(params)
	return
}
