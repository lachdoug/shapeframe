package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func getRepository() (command any) {
	command = &cliapp.Command{
		Name:    "repository",
		Summary: "Get a repository",
		Aliases: ss("r"),
		Usage: ss(
			"sf get repository [options] [URI]",
			"A repository URI may be provided as an argument",
			"  Otherwise prompt for URI",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Name of the workspace",
		),
		Handler:    getRepositoryHandler,
		Controller: controllers.RepositoriesRead,
		Viewer: cliapp.View(
			"repositories/read",
			"repositories/branches",
			"repositories/shapers",
			"repositories/framers",
		),
	}
	return
}

func getRepositoryHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.RepositoriesReadParams{
			Workspace: context.StringFlag("workspace"),
			URI:       context.Argument(0),
		},
	}
	return
}
