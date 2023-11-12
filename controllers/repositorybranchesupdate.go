package controllers

import (
	"sf/models"
	"sf/utils"
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

func RepositoryBranchesUpdate(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var r *models.Repository
	params := ParamsFor[RepositoryBranchesUpdateParams](jparams)
	workspace := params.Workspace
	uri := params.URI
	branch := params.Branch
	st := utils.StreamCreate()

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, workspace,
		"Repositories",
	); err != nil {
		return
	}
	if r, err = models.ResolveRepository(w, uri,
		"GitRepo",
	); err != nil {
		return
	}
	if err = r.Checkout(branch); err != nil {
		return
	}

	if branch, err = r.GitRepo.Branch(); err != nil {
		return
	}
	result := &RepositoryBranchesUpdateResult{
		Workspace: r.Workspace.Name,
		URI:       params.URI,
		Branch:    branch,
	}

	jbody = jbodyFor(result, nil, st)
	return
}
