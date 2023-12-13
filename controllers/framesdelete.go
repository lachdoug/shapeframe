package controllers

import (
	"sf/models"
)

type FramesDeleteParams struct {
	Workspace string
	Frame     string
}

type FramesDeleteResult struct {
	Workspace string
	Frame     string
}

func FramesDelete(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var f *models.Frame
	p := params.Payload.(*FramesDeleteParams)

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace", "Frame", "Shape",
	)
	if w, err = models.ResolveWorkspace(uc, p.Workspace,
		"Frames",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, p.Frame); err != nil {
		return
	}
	if uc.Frame != nil && uc.Frame.ID == f.ID {
		uc.Clear("Shape")
		uc.Clear("Frame")
	}
	f.Delete()

	result = &Result{
		Payload: &FramesDeleteResult{
			Workspace: w.Name,
			Frame:     f.Name,
		},
	}
	return
}
