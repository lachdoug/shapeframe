package controllers

import (
	"sf/models"
)

type RepositoriesDestroyParams struct {
	Workspace string
	URI       string
}

type RepositoriesDestroyResult struct {
	URI       string
	Workspace string
}

func RepositoriesDestroy(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var r *models.Repository
	params := ParamsFor[RepositoriesDestroyParams](jparams)

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, params.Workspace, "Repositories"); err != nil {
		return
	}
	if r, err = models.ResolveRepository(w, params.URI, "GitRepo"); err != nil {
		return
	}
	r.Destroy()

	result := &RepositoriesDestroyResult{
		URI:       params.URI,
		Workspace: r.Workspace.Name,
	}

	jbody = jbodyFor(result)
	return
}
