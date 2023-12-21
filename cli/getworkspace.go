package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func getWorkspace() (command any) {
	command = &cliapp.Command{
		Name:    "workspace",
		Summary: "Get workspace",
		Aliases: ss("w"),
		Usage: ss(
			"sf get workspace",
		),
		Controller: controllers.WorkspacesRead,
		Viewer: cliapp.View(
			"workspaces/read",
			"workspaces/frames",
			"workspaces/repositories",
			"workspaces/directories",
		),
	}
	return
}
