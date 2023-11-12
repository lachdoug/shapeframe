package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
	"sf/utils"
)

func addDirectory() (command any) {
	command = &cliapp.Command{
		Name:    "directory",
		Summary: "Add a directory to workspace",
		Aliases: ss("d"),
		Usage: ss(
			"sf add directory [options] [path]",
			"An absolute or relative (to working directory) path must be provided as an argument",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
		),
		Parametizer: addDirectoryParams,
		Controller:  controllers.DirectoriesCreate,
		Viewer:      addDirectoryViewer,
	}
	return
}

func addDirectoryParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	path := context.Argument(0)
	workspace := context.StringFlag("workspace")

	uc := models.ResolveUserContext(
		"Workspace",
		"Workspaces",
	)
	if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"Path":      path,
	})
	return
}

func addDirectoryViewer(body map[string]any) (output string, err error) {
	var w *models.Workspace
	var d *models.Directory
	var gri *models.GitRepoInspector
	result := resultItem(body)
	workspace := result["Workspace"].(string)
	path := result["Path"].(string)

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, workspace,
		"Directories",
	); err != nil {
		return
	}
	if d, err = models.ResolveDirectory(w, path,
		"Workspace", "Shapers", "Framers",
	); err != nil {
		return
	}

	// Convert GitRepoInspector to map for use in view
	if gri, err = d.GitRepo.Inspect(); err != nil {
		return
	}
	result["GitRepo"] = utils.Map(gri)
	body["Result"] = result
	output, err = cliapp.View(
		"directories/create",
		"gitrepos/gitrepo",
		"gitrepos/shapers",
		"gitrepos/shaper",
		"gitrepos/framers",
		"gitrepos/framer",
	)(body)
	return
}
