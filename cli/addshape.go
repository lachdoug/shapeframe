package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func addShape() (command any) {
	command = &cliapp.Command{
		Name:    "shape",
		Summary: "Add a shape to a frame",
		Aliases: ss("s"),
		Usage: ss(
			"sf add shape [options] [name]",
			"A shaper name must be provided as an argument",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Is required if -workspace flag is set",
			"  Otherwise uses frame context when not provided",
			"Provide an optional shape name using the -name flag",
			"  Uses shaper name when not provided",
			"Provide an optional about using the -about flag",
			"  Uses shaper about when not provided",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
			"string", "frame", "Name of the frame",
			"string", "name", "Name for the shape",
			"string", "about", "About the shape",
		),
		Parametizer: addShapeParams,
		Controller:  controllers.ShapesCreate,
		Viewer:      cliapp.View("shapes/create"),
	}
	return
}

func addShapeParams(context *cliapp.Context) (jparams []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var f *models.Frame
	shaper := context.Argument(0)
	name := context.StringFlag("name")
	about := context.StringFlag("about")
	frame := context.StringFlag("frame")
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

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Frame":     f.Name,
		"Shaper":    shaper,
		"Name":      name,
		"About":     about,
	})
	return
}
