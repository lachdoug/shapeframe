package controllers

import (
	"sf/models"
)

type FramesReadParams struct {
	Workspace string
	Frame     string
}

func FramesRead(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var f *models.Frame
	p := params.Payload.(*FramesReadParams)

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace", "Frame",
	)
	if w, err = models.ResolveWorkspace(uc, p.Workspace,
		"Frames",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, p.Frame,
		"Workspace", "Configuration", "Shapes", "Parent",
	); err != nil {
		return
	}

	result = &Result{
		Payload: f.Read(),
	}
	return
}
