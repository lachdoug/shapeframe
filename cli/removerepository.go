package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func removeRepository() (command any) {
	command = &cliapp.Command{
		Name:    "repository",
		Summary: "Remove a repository from a workspace",
		Aliases: ss("r"),
		Usage: ss(
			"sf remove repository [options] URI",
			"A repository URI must be provided as an argument",
		),
		Handler:    removeRepositoryHandler,
		Controller: controllers.RepositoriesDelete,
		Viewer:     cliapp.View("repositories/delete"),
	}
	return
}

func removeRepositoryHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.RepositoriesDeleteParams{
			URI: context.Argument(0),
		},
	}
	return
}
