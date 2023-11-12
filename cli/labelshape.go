package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func labelShape() (command any) {
	command = &cliapp.Command{
		Name:    "shape",
		Summary: "Change name and/or about for shape",
		Aliases: ss("s"),
		Usage: ss(
			"sf label shape [options] [name]",
			"Provide an optional shape name as an argument",
			"  Uses shape context when not provided",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Is required if -workspace flag is set",
			"  Uses workspace context when not provided",
			"Provide an optional shape update name using the -name flag",
			"Provide an optional shape update about using the -about flag",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
			"string", "frame", "Name of the frame",
			"string", "name", "New name for the shape",
			"string", "about", "New about for the shape",
		),
		Parametizer: labelShapeParams,
		Controller:  controllers.ShapesUpdate,
		Viewer:      cliapp.View("shapes/update", "labels/label"),
	}
	return
}

func labelShapeParams(context *cliapp.Context) (jparams []byte, err error) {
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
		"Shape":     s.Name,
		"Update":    update,
	}

	jparams = jsonParams(params)
	return
}
