package controllers

import (
	"path/filepath"
	"sf/app"
	"sf/models"
)

type DirectoriesDestroyParams struct {
	Workspace string
	Path      string
}

type DirectoriesDestroyResult struct {
	Workspace string
	Path      string
}

func DirectoriesDestroy(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var d *models.Directory
	var path string
	params := paramsFor[DirectoriesDestroyParams](jparams)

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

	if path, err = filepath.Abs(params.Path); err != nil {
		return
	}

	uc := models.ResolveUserContext("Workspaces.Directories")
	if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
		return
	}
	if d, err = models.ResolveDirectory(w, path); err != nil {
		return
	}
	d.Destroy()

	result := &DirectoriesDestroyResult{
		Path:      path,
		Workspace: w.Name,
	}

	jbody = jbodyFor(result)
	return
}
