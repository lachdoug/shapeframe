package controllers

import (
	"sf/models"
)

type FramesReadParams struct {
	Frame string
}

func FramesRead(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var f *models.Frame
	p := params.Payload.(*FramesReadParams)

	if w, err = models.ResolveWorkspace(
		"Frames", "Frame",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(w, p.Frame,
		"Workspace", "Configuration.Info", "Shapes", "Parent",
	); err != nil {
		return
	}

	result = &Result{
		Payload: f.Read(),
	}
	return
}
