package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func labelWorkspace() (command any) {
	command = &cliapp.Command{
		Name:    "workspace",
		Summary: "Change name and/or about for workspace",
		Aliases: ss("w"),
		Usage: ss(
			"sf label workspace [options] [name]",
			"Provide an optional workspace name as an argument",
			"  Uses workspace context when not provided",
			"Provide an optional workspace update name using the -name flag",
			"Provide an optional workspace update about using the -about flag",
		),
		Flags: ss(
			"string", "name", "New name for the workspace",
			"string", "about", "New about for the workspace",
		),
		Parametizer: labelWorkspaceParams,
		Controller:  controllers.WorkspacesUpdate,
		Viewer:      cliapp.View("workspaces/update", "labels/label"),
	}
	return
}

func labelWorkspaceParams(context *cliapp.Context) (jparams []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	workspace := context.Argument(0)

	uc := models.ResolveUserContext(
		"Workspace",
		"Workspaces",
	)
	if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
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
		"Update":    update,
	}

	jparams = jsonParams(params)
	return
}
