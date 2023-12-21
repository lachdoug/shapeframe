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
			"sf get repository [options] URI",
			"A repository URI must be provided as an argument",
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
			URI: context.Argument(0),
		},
	}
	return
}
