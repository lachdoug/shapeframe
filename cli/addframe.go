package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func addFrame() (command any) {
	command = &cliapp.Command{
		Name:    "frame",
		Summary: "Add a frame to a workspace",
		Aliases: ss("f"),
		Usage: ss(
			"sf add frame [options] [name]",
			"A framer name must be provided as an argument",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame name using the -name flag",
			"  Uses framer name when not provided",
			"Provide an optional about using the -about flag",
			"  Uses framer about when not provided",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
			"string", "name", "Name for the frame",
			"string", "about", "About the frame",
		),
		Parametizer: addFrameParams,
		Controller:  controllers.FramesCreate,
		Viewer:      cliapp.View("frames/create"),
	}
	return
}

func addFrameParams(context *cliapp.Context) (jparams []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	framer := context.Argument(0)
	name := context.StringFlag("name")
	about := context.StringFlag("about")
	workspace := context.StringFlag("workspace")

	uc := models.ResolveUserContext("Workspace", "Workspaces")
	if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Framer":    framer,
		"Name":      name,
		"About":     about,
	})
	return
}
