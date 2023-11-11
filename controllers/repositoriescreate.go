package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type RepositoriesCreateParams struct {
	Workspace string
	URI       string
	Protocol  string
}

type RepositoriesCreateResult struct {
	Path      string
	URI       string
	URL       string
	Workspace string
}

func RepositoriesCreate(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var r *models.Repository
	var vn *app.Validation
	params := ParamsFor[RepositoriesCreateParams](jparams)
	st := utils.StreamCreate()

	uc := models.ResolveUserContext(
		"Workspaces",
	)
	if w, err = models.ResolveWorkspace(uc, params.Workspace, "Repositories"); err != nil {
		return
	}
	if r, vn, err = models.CreateRepository(w, params.URI, params.Protocol, st); err != nil {
		return
	}

	result := &RepositoriesCreateResult{
		Workspace: r.Workspace.Name,
		URI:       params.URI,
		URL:       r.OriginURL(),
	}

	jbody = jbodyFor(result, vn, st)
	return
}
