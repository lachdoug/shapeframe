package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func removeDirectory() (command any) {
	command = &cliapp.Command{
		Name:    "directory",
		Summary: "Remove a directory from a workspace",
		Aliases: ss("d"),
		Usage: ss(
			"sf remove directory [options] PATH",
			"An absolute or relative (to working directory) path must be provided as an argument",
		),
		Handler:    removeDirectoryHandler,
		Controller: controllers.DirectoriesDelete,
		Viewer:     cliapp.View("directories/delete"),
	}
	return
}

func removeDirectoryHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.DirectoriesDeleteParams{
			Path: context.Argument(0),
		},
	}
	return
}
