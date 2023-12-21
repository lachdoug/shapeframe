package controllers

import (
	"sf/models"
)

func WorkspaceInspectsRead(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var wi *models.WorkspaceInspector

	if w, err = models.ResolveWorkspace(
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
