package controllers

import (
	"path/filepath"
	"sf/app/validations"
	"sf/models"
)

type DirectoriesCreateParams struct {
	Path string
}

type DirectoriesCreateResult struct {
	Workspace string
	Path      string
}

func DirectoriesCreate(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var d *models.Directory
	var path string
	var vn *validations.Validation
	p := params.Payload.(*DirectoriesCreateParams)

	if w, err = models.ResolveWorkspace(
		"Directories",
	); err != nil {
		return
	}
	if path, err = filepath.Abs(p.Path); err != nil {
		return
	}
	if d, vn, err = models.CreateDirectory(w, path); err != nil {
		return
	}

	result = &Result{
		Payload: &DirectoriesCreateResult{
			Workspace: w.Name,
			Path:      d.Path,
		},
		Validation: vn,
	}
	return
}
