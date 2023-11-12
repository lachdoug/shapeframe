package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func removeShape() (command any) {
	command = &cliapp.Command{
		Name:    "shape",
		Summary: "Remove a shape from a frame",
		Aliases: ss("s"),
		Usage: ss(
			"sf remove shape [options] [name]",
			"Provide an optional shape name as an argument",
			"  Uses shape context when not provided",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Is required if -workspace flag is set",
			"  Otherwise uses frame context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
			"string", "frame", "Name of the frame",
		),
		Parametizer: removeShapeParams,
		Controller:  controllers.ShapesDestroy,
		Viewer:      cliapp.View("shapes/destroy"),
	}
	return
}

func removeShapeParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	shape := context.Argument(0)
	frame := context.StringFlag("frame")
	workspace := context.StringFlag("workspace")

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace", "Frame", "Shape",
	)
	if w, err = models.ResolveWorkspace(uc, workspace,
		"Frames",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, frame,
		"Shapes",
	); err != nil {
		return
	}
	if s, err = models.ResolveShape(uc, f, shape); err != nil {
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Frame":     f.Name,
		"Shape":     s.Name,
	})
	return
}
