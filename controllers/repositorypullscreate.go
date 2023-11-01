package controllers

import (
	"net/url"
	"sf/app"
	"sf/models"
	"sf/utils"
)

type RepositoryPullsUpdateParams struct {
	Workspace string
	URI       string
}

type RepositoryPullsUpdateResult struct {
	URI       string
	GitURL    string
	Workspace string
}

func RepositoryPullsUpdate(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	params := &RepositoryPullsUpdateParams{}
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
	uc.Load("Workspaces")
	w := uc.WorkspaceFind(params.Workspace)
	if w == nil {
		err = app.Error(nil, "workspace %s does not exist", params.Workspace)
		return
	}
	r := w.RepositoryFind(w.GitRepoDirectoryFor(params.URI))
	if r == nil {
		err = app.Error(nil, "repository %s does not exist in workspace %s", params.URI, w.Name)
		return
	}
	r.Load()

	s := models.StreamCreate()
	r.Update(s)

	result := &RepositoryPullsUpdateResult{
		URI:       params.URI,
		GitURL:    r.GitRepo.URL,
		Workspace: r.Workspace.Name,
	}

	body := &app.Body{Result: result, Stream: s.Identifier}
	jbody = utils.JsonMarshal(body)
	return
}
