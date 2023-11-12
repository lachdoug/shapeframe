package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func orchestrate() (command any) {
	command = &cliapp.Command{
		Name:    "orchestrate",
		Summary: "Orchestrate a frame",
		Aliases: ss("o"),
		Usage: ss(
			"sf orchestrate [options]",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Provide an optional frame name using the -frame flag",
			"  Is required if -workspace flag is set",
			"  Uses frame context when not provided",
		),
		Flags: ss(
			"string", "frame", "Frame name",
			"string", "workspace", "Workspace name",
		),
		Parametizer: orchestrateParams,
		Controller:  controllers.FrameOrchestrationsCreate,
		Viewer:      orchestrateViewer,
	}
	return
}

func orchestrateParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	workspace := context.StringFlag("workspace")
	frame := context.StringFlag("frame")

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace", "Frame",
	)
	if w, err = models.ResolveWorkspace(uc, workspace,
		"Frames",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, frame); err != nil {
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Frame":     f.Name,
	})
	return
}

func orchestrateViewer(body map[string]any) (output string, err error) {
	if err = stream(body); err != nil {
		return
	}
	output, err = cliapp.View("frameorchestrations/create")(body)
	return
}
