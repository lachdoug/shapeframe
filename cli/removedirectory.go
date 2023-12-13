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
			"sf remove directory [options] [path]",
			"An absolute or relative (to working directory) path must be provided as an argument",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
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
			Workspace: context.StringFlag("workspace"),
			Path:      context.Argument(0),
		},
	}
	return
}
