package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
	"sf/utils"
)

func addRepository() (command any) {
	command = &cliapp.Command{
		Name:    "repository",
		Summary: "Add a repository to workspace",
		Aliases: ss("r"),
		Usage: ss(
			"sf add repository [options] [URI]",
			"A repository URI must be provided as an argument",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"Clone with https by setting the -https flag",
			"  Otherwise performs git clone using SSH",
			"Include a username for HTTPS clone by setting the -username flag",
			"  Otherwise performs git clone without a token when using HTTPS",
			"Include a password (or personal access token) for HTTPS clone by setting the -password flag",
			"  Otherwise performs git clone without a password when using HTTPS",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
			"bool", "https H", "Use HTTPS for git clone",
			"string", "username u", "Username for git clone",
			"string", "password p", "Password for git clone",
		),
		Parametizer: addRepositoryParams,
		Controller:  controllers.RepositoriesCreate,
		Viewer:      addRepositoryViewer,
	}
	return
}

func addRepositoryParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	var protocol string
	uri := context.Argument(0)
	workspace := context.StringFlag("workspace")
	username := context.StringFlag("username")
	password := context.StringFlag("password")

	if context.BoolFlag("https") {
		protocol = "HTTPS"
	}

	uc := models.ResolveUserContext("Workspace", "Workspaces")
	if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"URI":       uri,
		"Protocol":  protocol,
		"Username":  username,
		"Password":  password,
	})
	return
}

func addRepositoryViewer(body map[string]any) (output string, err error) {
	var w *models.Workspace
	var r *models.Repository
	var gri *models.GitRepoInspector
	result := resultItem(body)
	workspace := result["Workspace"].(string)
	uri := result["URI"].(string)

	// Stream clone output
	app.Printf("Clone %s\n", result["URL"])
	if err = stream(body); err != nil {
		return
	}

	// If clone is successful, show repository contents using GitRepo inspection
	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, workspace, "Repositories"); err != nil {
		return
	}
	if r, err = models.ResolveRepository(w, uri,
		"Shapers", "Framers",
	); err != nil {
		return
	}

	// Convert GitRepoInspector to map for use in view
	if gri, err = r.GitRepo.Inspect(); err != nil {
		return
	}
	result["GitRepo"] = utils.Map(gri)
	output, err = cliapp.View(
		"repositories/create",
		"gitrepos/gitrepo",
		"gitrepos/shapers",
		"gitrepos/shaper",
		"gitrepos/framers",
		"gitrepos/framer",
	)(body)
	return
}
