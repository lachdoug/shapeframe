package controllers

import (
	"sf/models"
)

type WorkspaceInspectsReadParams struct {
	Workspace string
}

func WorkspaceInspectsRead(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var wi *models.WorkspaceInspector
	p := params.Payload.(*WorkspaceInspectsReadParams)

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace",
	)
	if w, err = models.ResolveWorkspace(uc, p.Workspace,
		"Framers", "Shapers",
		"Frames.Parent", "Frames.Children",
		"Frames.Configuration", "Frames.Shapes.Configuration",
	); err != nil {
		return
	}

	if wi, err = w.Inspect(); err != nil {
		return
	}

	result = &Result{
		Payload: wi,
	}
	return
}
