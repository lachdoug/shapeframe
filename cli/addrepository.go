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
			"  Otherwise uses SSS",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
			"bool", "https", "Use HTTPS for git clone and pull",
		),
		Parametizer: addRepositoryParams,
		Controller:  controllers.RepositoriesCreate,
		Viewer:      addRepositoryViewer,
	}
	return
}

func addRepositoryParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	uri := context.Argument(0)
	protocol := "ssh"
	workspace := context.StringFlag("workspace")

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
	})
	return
}

func addRepositoryViewer(body map[string]any) (output string, err error) {
	var w *models.Workspace
	var r *models.Repository
	var gri *models.GitRepoInspector
	result := resultItem(body)

	// Stream clone output
	app.Printf("Clone %s\n", result["URL"])
	if err = stream(body); err != nil {
		return
	}

	// If clone is successful, show repository contents using GitRepo inspection
	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, result["Workspace"].(string), "Repositories"); err != nil {
		return
	}
	if r, err = models.ResolveRepository(w, result["URI"].(string)); err != nil {
		return
	}
	if err = r.Load("Shapers", "Framers"); err != nil {
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
