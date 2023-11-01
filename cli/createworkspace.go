package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
)

func createWorkspace() (command any) {
	command = &cliapp.Command{
		Name:    "workspace",
		Summary: "Create a workspace",
		Usage:   ss("sf create workspace [command options] [name]"),
		Description: ss(
			"A name must be provided as the first argument",
			"Provide an optional about using the --about flag",
		),
		Aliases:     ss("w"),
		Flags:       ss("string", "about", "About the workspace"),
		Parametizer: createWorkspaceParams,
		Controller:  controllers.WorkspacesCreate,
		Viewer:      cliapp.View("workspaces/create"),
	}
	return
}

func createWorkspaceParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	jparams = jsonParams(map[string]any{
		"Name":  context.Argument(0),
		"About": context.StringFlag("about"),
	})
	return
}
