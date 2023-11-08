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

func DirectoriesCreate(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var d *models.Directory
	var path string
	params := paramsFor[DirectoriesCreateParams](jparams)

	vn = &app.Validation{}
	if params.Workspace == "" {
		vn.Add("Workspace", "must not be blank")
	}
	if params.Path == "" {
		vn.Add("Path", "must not be blank")
	}
	if vn.IsInvalid() {
		return
	}

	uc := models.ResolveUserContext(
		"Workspaces.Directories",
	)
	if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
		return
	}
	if path, err = filepath.Abs(params.Path); err != nil {
		return
	}
	if d, err = models.CreateDirectory(w, path); err != nil {
		return
	}

	result := &DirectoriesCreateResult{
		Workspace: w.Name,
		Path:      d.Path,
	}

	jbody = jbodyFor(result)
	return
}
