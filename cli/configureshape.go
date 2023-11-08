package cli

import (
	"sf/app"
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
		Controller:  controllers.ShapesUpdate,
		Viewer:      configureShapeViewer,
	}
	return
}

func configureShapeParams(context *cliapp.Context) (jparams []byte, vn *app.Validation, err error) {
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
	c := s.Configuration
	if len(values) == 0 {
		form := &Form{Model: c}
		if settings, err = form.prompts(); err != nil {
			return
		}
	} else {
		if settings, err = c.SettingsForValues(values); err != nil {
			return
		}
	}

	update := map[string]any{
		"Configuration": settings,
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Frame":     f.Name,
		"Shape":     s.Name,
		"Update":    update,
	})
	return
}

func configureShapeViewer(body map[string]any) (output string, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	result := resultItem(body)

	uc := models.ResolveUserContext("Workspaces.Frames.Shapes")
	if w, err = models.ResolveWorkspace(uc, result["Workspace"].(string)); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, result["Frame"].(string)); err != nil {
		return
	}
	if s, err = models.ResolveShape(uc, f, result["Shape"].(string), "Configuration"); err != nil {
		return
	}

	if v := s.Configuration.Validate(); v != nil {
		err = app.ErrorWith(v, "invalid")
		return
	} else {
		if output, err = cliapp.View(
			"shapeconfigurations/update",
			"configurations/configuration",
			"configurations/setting",
		)(body); err != nil {
			return
		}
	}
	return
}
