package controllers

import (
	"sf/models"
)

type FramesDeleteParams struct {
	Frame string
}

type FramesDeleteResult struct {
	Workspace string
	Frame     string
}

func FramesDelete(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var f *models.Frame
	p := params.Payload.(*FramesDeleteParams)

	if w, err = models.ResolveWorkspace(
		"Frames", "Frame",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(w, p.Frame); err != nil {
		return
	}
	if w.Frame != nil && w.Frame.ID == f.ID {
		w.Clear("Shape")
		w.Clear("Frame")
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
