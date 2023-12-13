package controllers

import (
	"path/filepath"
	"sf/models"
)

type DirectoriesDeleteParams struct {
	Workspace string
	Path      string
}

type DirectoriesDeleteResult struct {
	Workspace string
	Path      string
}

func DirectoriesDelete(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var d *models.Directory
	var path string
	p := params.Payload.(*DirectoriesDeleteParams)

	if path, err = filepath.Abs(p.Path); err != nil {
		return
	}

	uc := models.ResolveUserContext("Workspaces", "Workspace")
	if w, err = models.ResolveWorkspace(uc, p.Workspace,
		"Directories",
	); err != nil {
		return
	}
	if d, err = models.ResolveDirectory(w, path); err != nil {
		return
	}
	d.Delete()

	result = &Result{
		Payload: &DirectoriesDeleteResult{
			Path:      path,
			Workspace: w.Name,
		},
	}
	return
}
