package controllers

import (
	"net/url"
	"sf/app"
	"sf/models"
)

type RepositoriesCreateParams struct {
	Workspace string
	URI       string
	SSH       bool
}

type RepositoriesCreateResult struct {
	Path      string
	URI       string
	URL       string
	Workspace string
}

func RepositoriesCreate(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var r *models.Repository
	params := paramsFor[RepositoriesCreateParams](jparams)
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

	uc := models.ResolveUserContext(
		"Workspaces",
	)
	if w, err = models.ResolveWorkspace(uc, params.Workspace, "Repositories"); err != nil {
		return
	}
	if r, err = models.CreateRepository(w, params.URI, params.SSH, st); err != nil {
		return
	}

	result := &RepositoriesCreateResult{
		Workspace: r.Workspace.Name,
		URI:       params.URI,
		URL:       r.GitRepo.OriginURL(params.URI, params.SSH),
	}

	jbody = jbodyFor(result, st)
	return
}
