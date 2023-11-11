package cli

import (
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
		Controller:  controllers.FrameConfigurationsUpdate,
		Viewer: cliapp.View(
			"frameconfigurations/update",
			"configurations/configuration",
			"configurations/setting",
		),
	}
	return
}

func configureFrameParams(context *cliapp.Context) (jparams []byte, err error) {
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
	fm := f.Configuration.Form
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
		"Update":    settings,
	})
	return
}
