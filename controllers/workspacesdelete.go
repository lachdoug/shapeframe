package controllers

import (
	"sf/models"
)

type WorkspacesDeleteParams struct {
	Workspace string
}

type WorkspacesDeleteResult struct {
	Workspace string
}

func WorkspacesDelete(params *Params) (result *Result, err error) {
	p := params.Payload.(*WorkspacesDeleteParams)
	var w *models.Workspace

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace", "Frame", "Shape",
	)
	if w, err = models.ResolveWorkspace(uc, p.Workspace); err != nil {
		return
	}
	if uc.Workspace != nil && uc.Workspace.ID == w.ID {
		uc.Clear("Shape")
		uc.Clear("Frame")
		uc.Clear("Workspace")
	}
	w.Delete()

	result = &Result{
		Payload: &WorkspacesDeleteResult{
			Workspace: w.Name,
		},
	}
	return
}
