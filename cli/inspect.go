package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func inspect() (command any) {
	command = &cliapp.Command{
		Name:    "inspect",
		Aliases: ss("i"),
		Summary: "Inspect workspace",
		Usage: ss(
			"sf inspect workspace [options] [name]",
			"Provide an optional workspace name as an argument",
			"  Uses workspace context when not provided",
		),
		Parametizer: inspectParams,
		Controller:  controllers.WorkspaceInspectsRead,
		Viewer: cliapp.View(
			"workspaceinspects/read",
			"workspaceinspects/frames",
			"workspaceinspects/frame",
			"workspaceinspects/frame/shape",
			"workspaceinspects/frame/shapes",
			"workspaceinspects/frame/relations",
			"workspaceinspects/frame/children",
			"workspaceinspects/frame/ancestors",
			"workspaceinspects/frame/descendents",
			"workspaceinspects/repositories",
			"workspaceinspects/repository",
			"workspaceinspects/directories",
			"workspaceinspects/directory",
			"configurations/configuration",
			"configurations/setting",
			"gitrepos/gitrepo",
			"gitrepos/framers",
			"gitrepos/framer",
			"gitrepos/shapers",
			"gitrepos/shaper",
		),
	}
	return
}

func inspectParams(context *cliapp.Context) (jparams []byte, err error) {
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
