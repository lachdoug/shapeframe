package cli

import (
	"sf/cli/cliapp"
	"sf/cli/prompting"
	"sf/controllers"
	"sf/models"
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
		Parametizer: getRepositoryParams,
		Controller:  controllers.RepositoriesRead,
		Viewer: cliapp.View(
			"repositories/read",
			"repositories/branches",
			"repositories/shapers",
			"repositories/framers",
		),
	}
	return
}

func getRepositoryParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	uri := context.Argument(0)
	workspace := context.StringFlag("workspace")

	uc := models.ResolveUserContext("Workspace", "Workspaces")
	if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
		return
	}

	if uri == "" {
		if uri, err = prompting.RepositoryURI(w); err != nil {
			return
		}
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"URI":       uri,
	})
	return
}
