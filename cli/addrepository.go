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
			"Clone with ssh by setting the -ssh flag",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
			"bool", "ssh", "Use SSH for git clone",
		),
		Parametizer: addRepositoryParams,
		Controller:  controllers.RepositoriesCreate,
		Viewer:      addRepositoryViewer,
	}
	return
}

func addRepositoryParams(context *cliapp.Context) (jparams []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	uri := context.Argument(0)
	ssh := context.BoolFlag("ssh")
	workspace := context.StringFlag("workspace")

	uc := models.ResolveUserContext("Workspace", "Workspaces")
	if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"URI":       uri,
		"SSH":       ssh,
	})
	return
}

func addRepositoryViewer(body map[string]any) (output string, err error) {
	var w *models.Workspace
	var r *models.Repository
	result := resultItem(body)

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, result["Workspace"].(string), "Repositories"); err != nil {
		return
	}
	if r, err = models.ResolveRepository(w, result["URI"].(string)); err != nil {
		return
	}

	// Stream clone output
	app.Printf("Clone %s\n", result["URL"])
	if err = stream(body); err != nil {
		return
	}

	// If clone is successful, show repository contents using GitRepo inspection
	if err = r.Load("Shapers", "Framers"); err != nil {
		return
	}

	// Convert GitRepoInspector to map for use in view
	result["GitRepo"] = utils.Map(r.GitRepo.Inspect())
	body["Result"] = result
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
