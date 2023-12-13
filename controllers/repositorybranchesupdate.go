package controllers

import (
	"sf/models"
)

type RepositoryBranchesUpdateParams struct {
	Workspace string
	URI       string
	Branch    string
}

type RepositoryBranchesUpdateResult struct {
	Workspace string
	URI       string
	Branch    string
}

func RepositoryBranchesUpdate(params *Params) (result *Result, err error) {
	p := params.Payload.(*RepositoryBranchesUpdateParams)
	var w *models.Workspace
	var r *models.Repository

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
	if err = r.Checkout(p.Branch); err != nil {
		return
	}

	result = &Result{
		Payload: &RepositoryBranchesUpdateResult{
			Workspace: r.Workspace.Name,
			URI:       p.URI,
			Branch:    p.Branch,
		},
	}
	return
}
