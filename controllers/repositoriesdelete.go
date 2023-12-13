package controllers

import (
	"sf/models"
)

type RepositoriesDeleteParams struct {
	Workspace string
	URI       string
}

type RepositoriesDeleteResult struct {
	Workspace string
	URI       string
}

func RepositoriesDelete(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var r *models.Repository
	p := params.Payload.(*RepositoriesDeleteParams)

	uc := models.ResolveUserContext("Workspaces", "Workspace")
	if w, err = models.ResolveWorkspace(uc, p.Workspace,
		"Repositories",
	); err != nil {
		return
	}
	if r, err = models.ResolveRepository(w, p.URI,
		"GitRepo",
	); err != nil {
		return
	}
	r.Delete()

	result = &Result{
		Payload: &RepositoriesDeleteResult{
			URI:       p.URI,
			Workspace: r.Workspace.Name,
		},
	}
	return
}
