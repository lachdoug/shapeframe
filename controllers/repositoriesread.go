package controllers

import (
	"sf/models"
)

type RepositoriesReadParams struct {
	Workspace string
	URI       string
}

func RepositoriesRead(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var r *models.Repository
	var rr *models.RepositoryReader
	params := ParamsFor[RepositoriesReadParams](jparams)

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, params.Workspace,
		"Repositories",
	); err != nil {
		return
	}
	if r, err = models.ResolveRepository(w, params.URI,
		"GitRepo", "Shapers", "Framers",
	); err != nil {
		return
	}

	if rr, err = r.Read(); err != nil {
		return
	}
	result := rr

	jbody = jbodyFor(result)
	return
}
