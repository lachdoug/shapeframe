package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func configureShape() (command any) {
	command = &cliapp.Command{
		Name:    "shape",
		Summary: "Configure a shape",
		Aliases: ss("s"),
		Usage: ss(
			"sf configure shape [options] [value1] [value2] [value3]...",
			"Configuration values may be provided as arguments",
			"  Values are mapped to configuration settings in the order provided",
			"  Prompts will be shown if no arguments are provided",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Is required if -workspace flag is set",
			"  Otherwise uses frame context when not provided",
			"Provide an optional shape name using the -shape flag",
			"  Is required if -frame flag is set",
			"  Uses shape context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
			"string", "frame", "Name of the frame",
			"string", "shape", "Name of the shape",
		),
		Parametizer: configureShapeParams,
		Controller:  controllers.ShapeConfigurationsUpdate,
		Viewer: cliapp.View(
			"shapeconfigurations/update",
			"configurations/configuration",
			"configurations/setting",
		),
	}
	return
}

func configureShapeParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	shape := context.StringFlag("shape")
	frame := context.StringFlag("frame")
	workspace := context.StringFlag("workspace")
	values := context.Arguments()

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

	var settings map[string]any
	fm := s.Configuration.Form
	if len(values) == 0 {
		form := &Form{Model: fm}
		if settings, err = form.prompts(); err != nil {
			return
		}
	} else {
		if settings, err = fm.SettingsForValues(values); err != nil {
			return
		}
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Frame":     f.Name,
		"Shape":     s.Name,
		"Update":    settings,
	})
	return
}
