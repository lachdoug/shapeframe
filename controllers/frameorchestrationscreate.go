package controllers

import (
	"sf/app/streams"
	"sf/models"
)

type FrameOrchestrationsCreateParams struct {
	Frame string
}

type FrameOrchestrationsCreateResult struct {
	Workspace string
	Frame     string
}

func FrameOrchestrationsCreate(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var f *models.Frame
	p := params.Payload.(*FrameOrchestrationsCreateParams)
	st := streams.StreamCreate()

	if w, err = models.ResolveWorkspace(
		"Frames", "Frame",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(w, p.Frame,
		"Configuration", "Shapes.Configuration",
	); err != nil {
		return
	}
	f.Apply(st)

	result = &Result{
		Payload: &FrameOrchestrationsCreateResult{
			Workspace: w.Name,
			Frame:     f.Name,
		},
		Stream: st,
	}
	return
}
