package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
)

func addWorkspace() (command any) {
	command = &cliapp.Command{
		Name:    "workspace",
		Summary: "Add a workspace",
		Aliases: ss("w"),
		Usage: ss(
			"sf add workspace [options] [name]",
			"A name must be provided as an argument",
			"Provide an optional about using the -about flag",
		),
		Flags: ss(
			"string", "about", "About the workspace",
		),
		Parametizer: addWorkspaceParams,
		Controller:  controllers.WorkspacesCreate,
		Viewer:      cliapp.View("workspaces/create"),
	}
	return
}

func addWorkspaceParams(context *cliapp.Context) (jparams []byte, vn *app.Validation, err error) {
	name := context.Argument(0)
	about := context.StringFlag("about")

	jparams = jsonParams(map[string]any{
		"Name":  name,
		"About": about,
	})
	return
}
