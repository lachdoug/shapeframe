package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func addDirectory() (command any) {
	command = &cliapp.Command{
		Name:    "directory",
		Summary: "Add a directory to workspace",
		Aliases: ss("d"),
		Usage: ss(
			"sf add directory [options] PATH",
			"A path must be provided as an argument",
		),
		Handler:    addDirectoryHandler,
		Controller: controllers.DirectoriesCreate,
		Viewer:     cliapp.View("directories/create"),
	}
	return
}

func addDirectoryHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.DirectoriesCreateParams{
			Path: context.Argument(0),
		},
	}
	return
}
