package controllers

import "sf/models"

type RepositoriesReadParams struct {
	URI string
}

func RepositoriesRead(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var r *models.Repository
	var rr *models.RepositoryReader
	p := params.Payload.(*RepositoriesReadParams)

	if w, err = models.ResolveWorkspace(
		"Repositories",
	); err != nil {
		return
	}
	if r, err = models.ResolveRepository(w, p.URI,
		"GitRepo", "Shapers", "Framers",
	); err != nil {
		return
	}

	if rr, err = r.Read(); err != nil {
		return
	}

	result = &Result{
		Payload: rr,
	}
	return
}
