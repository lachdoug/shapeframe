package cli

import (
	"fmt"
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
	"sf/utils"
)

func addRepository() (command any) {
	command = &cliapp.Command{
		Name:        "repository",
		Summary:     "Add a repository to workspace",
		Usage:       ss("sf add repository [command options] [URI]"),
		Aliases:     ss("r"),
		Flags:       ss("bool", "ssh", "Use SSH for git clone"),
		Parametizer: addRepositoryParams,
		Controller:  controllers.RepositoriesCreate,
		Viewer:      addRepositoryViewer,
	}
	return
}

func addRepositoryParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	uc := models.UserContextNew()
	uc.Load("Workspace")
	w := uc.Workspace
	if w == nil {
		err = app.Error(nil, "no workspace context")
		return
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"URI":       context.Argument(0),
		"SSH":       context.BoolFlag("ssh"),
	})
	return
}

func addRepositoryViewer(body map[string]any) (output string, err error) {
	result := resultItem(body)
	fmt.Printf("Clone %s\n", result["URL"])
	if err = stream(body); err != nil {
		return
	}
	// If clone is successful, show repository contents using GitRepo inspection
	uc := models.UserContextNew()
	uc.Load("Workspace")
	w := uc.Workspace
	r := w.RepositoryFind(result["Path"].(string))
	r.Load()
	var gri *models.GitRepoInspector
	if gri, err = r.GitRepo.Inspect(); err != nil {
		return
	}
	// Convert GitRepoInspector to map for use in view template
	griMap := &map[string]any{}
	utils.JsonUnmarshal(utils.JsonMarshal(gri), griMap)
	result["GitRepo"] = griMap
	body["Result"] = result
	output, err = cliapp.View("repositories/create")(body)
	return
}
