package controllers

import (
	"net/url"
	"sf/app"
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

func RepositoriesDestroy(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var r *models.Repository
	params := paramsFor[RepositoriesDestroyParams](jparams)

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
	r.Destroy()

	result := &RepositoriesDestroyResult{
		URI:       params.URI,
		Workspace: r.Workspace.Name,
	}

	jbody = jbodyFor(result)
	return
}
