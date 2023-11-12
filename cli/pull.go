package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/cli/prompting"
	"sf/controllers"
	"sf/models"
)

func pull() (command any) {
	command = &cliapp.Command{
		Name:    "pull",
		Summary: "Pull a workspace repository",
		Aliases: ss("p"),
		Usage: ss(
			"sf pull [options] [URI]",
			"A repository URI may be provided as an argument",
			"  Otherwise prompt for URI",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Include a username for HTTPS pull by setting the -username flag",
			"  Otherwise performs git pull without a password when using HTTPS",
			"Include a password (or access token) for HTTPS pull by setting the -password flag",
			"  Otherwise performs git pull without a password when using HTTPS",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
			"string", "username u", "Username for git pull",
			"string", "password p", "Password for git pull",
		),
		Parametizer: pullParams,
		Controller:  controllers.RepositoryPullsCreate,
		Viewer:      pullViewer,
	}
	return
}

func pullParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	uri := context.Argument(0)
	workspace := context.StringFlag("workspace")
	username := context.StringFlag("username")
	password := context.StringFlag("password")

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
		"Username":  username,
		"Password":  password,
	})
	return
}

func pullViewer(body map[string]any) (output string, err error) {
	result := resultItem(body)
	app.Printf("Pull %s\n", result["URL"])
	if err = stream(body); err != nil {
		return
	}
	output, err = cliapp.View("repositorypulls/create")(body)
	return
}
