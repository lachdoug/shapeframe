package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func inspectWorkspace() (command any) {
	command = &cliapp.Command{
		Name:    "inspect",
		Summary: "Inspect workspace",
		Aliases: ss("w"),
		Usage: ss(
			"sf inspect workspace [options] [name]",
			"Provide an optional workspace name as an argument",
			"  Uses workspace context when not provided",
		),
		Parametizer: inspectWorkspaceParams,
		Controller:  controllers.WorkspacesRead,
		Viewer: cliapp.View(
			"workspaces/read",
			"workspaces/frame",
			"workspaces/frames",
			"workspaces/shape",
			"workspaces/shapes",
			"workspaces/repositories",
			"workspaces/repository",
			"workspaces/directories",
			"workspaces/directory",
			"configurations/configuration",
			"configurations/setting",
			"gitrepos/gitrepo",
			"gitrepos/framer",
			"gitrepos/framers",
			"gitrepos/shaper",
			"gitrepos/shapers",
		),
	}
	return
}

func inspectWorkspaceParams(context *cliapp.Context) (jparams []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	workspace := context.Argument(0)

	uc := models.ResolveUserContext(
		"Workspace",
		"Workspaces",
	)
	if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
		return
	}

	params := map[string]any{
		"Workspace": w.Name,
	}

	jparams = jsonParams(params)
	return
}
