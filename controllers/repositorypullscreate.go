package controllers

import (
	"sf/app/streams"
	"sf/models"
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

func RepositoryPullsCreate(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var r *models.Repository
	var url string
	p := params.Payload.(*RepositoryPullsCreateParams)
	st := streams.StreamCreate()

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
	r.Update(p.Username, p.Password, st)

	if url, err = r.GitRepo.URL(); err != nil {
		return
	}
	result = &Result{
		Payload: &RepositoryPullsCreateResult{
			URI:       p.URI,
			URL:       url,
			Workspace: r.Workspace.Name,
		},
		Stream: st,
	}
	return
}
