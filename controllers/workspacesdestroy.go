package controllers

import (
	"sf/models"
)

type WorkspacesDestroyParams struct {
	Workspace string
}

type WorkspacesDestroyResult struct {
	Workspace string
}

func WorkspacesDestroy(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	params := ParamsFor[WorkspacesDestroyParams](jparams)

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace", "Frame", "Shape",
	)
	if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
		return
	}
	if uc.Workspace != nil && uc.Workspace.ID == w.ID {
		uc.Clear("Shape")
		uc.Clear("Frame")
		uc.Clear("Workspace")
	}
	w.Destroy()

	result := &WorkspacesDestroyResult{
		Workspace: w.Name,
	}

	jbody = jbodyFor(result)
	return
}
