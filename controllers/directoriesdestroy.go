package controllers

import (
	"path/filepath"
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

func DirectoriesDestroy(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var d *models.Directory
	var path string
	params := ParamsFor[DirectoriesDestroyParams](jparams)

	if path, err = filepath.Abs(params.Path); err != nil {
		return
	}

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, params.Workspace,
		"Directories",
	); err != nil {
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
