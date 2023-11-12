package controllers

import (
	"sf/models"
	"sf/utils"
)

type RepositoryPullsCreateParams struct {
	Workspace string
	URI       string
	Username  string
	Password  string
}

type RepositoryPullsCreateResult struct {
	Workspace string
	URI       string
	URL       string
}

func RepositoryPullsCreate(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var r *models.Repository
	var url string
	params := ParamsFor[RepositoryPullsCreateParams](jparams)
	st := utils.StreamCreate()

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, params.Workspace,
		"Repositories",
	); err != nil {
		return
	}
	if r, err = models.ResolveRepository(w, params.URI,
		"GitRepo",
	); err != nil {
		return
	}
	r.Update(params.Username, params.Password, st)

	if url, err = r.GitRepo.URL(); err != nil {
		return
	}
	result := &RepositoryPullsCreateResult{
		URI:       params.URI,
		URL:       url,
		Workspace: r.Workspace.Name,
	}

	jbody = jbodyFor(result, nil, st)
	return
}
