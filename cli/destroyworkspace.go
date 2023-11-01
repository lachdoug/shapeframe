package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
)

func destroyWorkspace() (command any) {
	command = &cliapp.Command{
		Name:        "workspace",
		Summary:     "Destroy a workspace",
		Usage:       ss("sf destroy workspace [command options] [name]"),
		Aliases:     ss("w"),
		Parametizer: destroyWorkspaceParams,
		Controller:  controllers.WorkspacesDestroy,
		Viewer:      cliapp.View("workspaces/destroy"),
	}
	return
}

func destroyWorkspaceParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	jparams = jsonParams(map[string]any{
		"Name": context.Argument(0),
	})
	return
}
