package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func configureFrame() (command any) {
	command = &cliapp.Command{
		Name:    "frame",
		Summary: "Configure a frame",
		Aliases: ss("f"),
		Usage: ss(
			"sf configure frame [options] [value1] [value2] [value3]...",
			"Configuration values may be provided as arguments",
			"  Values are mapped to configuration settings in the order provided",
			"  Prompts will be shown if no arguments are provided",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Is required if -workspace flag is set",
			"  Uses frame context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
			"string", "frame", "Name of the frame",
		),
		Parametizer: configureFrameParams,
		Controller:  controllers.FramesUpdate,
		Viewer:      configureFrameViewer,
	}
	return
}

func configureFrameParams(context *cliapp.Context) (jparams []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var f *models.Frame
	frame := context.StringFlag("frame")
	workspace := context.StringFlag("workspace")
	values := context.Arguments()

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

	var settings map[string]any
	c := f.Configuration
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
		"Update":    update,
	})
	return
}

func configureFrameViewer(body map[string]any) (output string, err error) {
	var w *models.Workspace
	var f *models.Frame
	result := resultItem(body)

	if output, err = cliapp.View(
		"frameconfigurations/update",
		"configurations/configuration",
		"configurations/setting",
	)(body); err != nil {
		return
	}

	uc := models.ResolveUserContext("Workspaces.Frames")
	if w, err = models.ResolveWorkspace(uc, result["Workspace"].(string)); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, result["Frame"].(string), "Configuration"); err != nil {
		return
	}

	if v := f.Configuration.Validate(); v != nil {
		output = output + "\n" + app.ErrorWith(v, "invalid").Error() + "\n"
	}
	return
}
