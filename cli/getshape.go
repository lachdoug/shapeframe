package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func getShape() (command any) {
	command = &cliapp.Command{
		Name:    "inspect",
		Summary: "Inspect shape",
		Aliases: ss("s"),
		Usage: ss(
			"sf inspect shape [options] [name]",
			"Provide an optional shape name as an argument",
			"  Uses shape context when not provided",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Uses frame context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
			"string", "frame", "Frame name",
		),
		Parametizer: getShapeParams,
		Controller:  controllers.ShapesRead,
		Viewer: cliapp.View(
			"shapes/read",
			"configurations/configuration",
			"configurations/setting",
		),
	}
	return
}

func getShapeParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	shape := context.Argument(0)
	workspace := context.StringFlag("workspace")
	frame := context.StringFlag("frame")

	uc := models.ResolveUserContext(
		"Shape",
		"Frame.Shapes",
		"Workspace.Frames.Shapes",
		"Workspaces.Frames.Shapes",
	)
	if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, frame); err != nil {
		return
	}
	if s, err = models.ResolveShape(uc, f, shape, "Configuration"); err != nil {
		return
	}

	params := map[string]any{
		"Workspace": w.Name,
		"Frame":     f.Name,
		"Shape":     s.Name,
	}

	jparams = jsonParams(params)
	return
}
