package controllers

import (
	"path/filepath"
	"sf/app"
	"sf/models"
)

type DirectoriesCreateParams struct {
	Workspace string
	Path      string
}

type DirectoriesCreateResult struct {
	Workspace string
	Path      string
}

func DirectoriesCreate(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var d *models.Directory
	var path string
	var vn *app.Validation
	params := ParamsFor[DirectoriesCreateParams](jparams)

	uc := models.ResolveUserContext(
		"Workspaces.Directories",
	)
	if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
		return
	}
	if path, err = filepath.Abs(params.Path); err != nil {
		return
	}
	if d, vn, err = models.CreateDirectory(w, path); err != nil {
		return
	}

	result := &DirectoriesCreateResult{
		Workspace: w.Name,
		Path:      d.Path,
	}

	jbody = jbodyFor(result, vn)
	return
}
