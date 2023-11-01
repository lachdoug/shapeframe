package controllers

import (
	"net/url"
	"sf/app"
	"sf/models"
	"sf/utils"
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

func RepositoriesCreate(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	params := &RepositoriesCreateParams{}
	utils.JsonUnmarshal(jparams, params)

	v = &app.Validation{}
	if params.Workspace == "" {
		v.Add("Workspace", "must not be blank")
	}
	if params.URI == "" {
		v.Add("URI", "must not be blank")
	}
	if _, err = url.Parse("https://" + params.URI); err != nil {
		v.Add("URI", "must be valid URI")
	}
	if v.IsInvalid() {
		return
	}

	uc := models.UserContextNew()
	w := uc.WorkspaceFind(params.Workspace)
	if w == nil {
		err = app.Error(nil, "workspace %s does not exist", params.Workspace)
		return
	}
	path := w.GitRepoDirectoryFor(params.URI)
	r := models.RepositoryNew(w, path)
	if r.IsExists() {
		err = app.Error(nil, "repository %s already exists in workspace %s", params.URI, w.Name)
		return
	}

	s := models.StreamCreate()
	r.Create(s, params.URI, params.SSH)

	result := &RepositoriesCreateResult{
		Workspace: r.Workspace.Name,
		Path:      path,
		URI:       params.URI,
		URL:       r.GitRepo.OriginURL(params.URI, params.SSH),
	}
	body := &app.Body{Result: result, Stream: s.Identifier}
	jbody = utils.JsonMarshal(body)
	return
}
