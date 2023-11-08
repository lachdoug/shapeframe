package controllers

import (
	"net/url"
	"sf/app"
	"sf/models"
)

type RepositoryPullsCreateParams struct {
	Workspace string
	URI       string
}

type RepositoryPullsCreateResult struct {
	URI       string
	URL       string
	Workspace string
}

func RepositoryPullsCreate(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var r *models.Repository
	params := paramsFor[RepositoryPullsCreateParams](jparams)
	st := models.StreamCreate()

	vn = &app.Validation{}
	if params.Workspace == "" {
		vn.Add("Workspace", "must not be blank")
	}
	if params.URI == "" {
		vn.Add("URI", "must not be blank")
	}
	if _, err = url.Parse("https://" + params.URI); err != nil {
		vn.Add("URI", "must be valid URI")
	}
	if vn.IsInvalid() {
		return
	}

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, params.Workspace, "Repositories"); err != nil {
		return
	}
	if r, err = models.ResolveRepository(w, params.URI, "GitRepo"); err != nil {
		return
	}
	r.Update(st)

	result := &RepositoryPullsCreateResult{
		URI:       params.URI,
		URL:       r.GitRepo.URL,
		Workspace: r.Workspace.Name,
	}

	jbody = jbodyFor(result, st)
	return
}
