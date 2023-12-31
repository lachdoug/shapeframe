package controllers

import (
	"sf/app/validations"
	"sf/models"
)

type FramesCreateParams struct {
	Framer string
	Frame  string
	About  string
}

type FramesCreateResult struct {
	Workspace string
	Frame     string
}

func FramesCreate(params *Params) (result *Result, err error) {
	p := params.Payload.(*FramesCreateParams)
	var w *models.Workspace
	var f *models.Frame
	var vn *validations.Validation

	if w, err = models.ResolveWorkspace(
		"Frames", "Framers",
	); err != nil {
		return
	}
	if f, vn, err = models.CreateFrame(w, p.Framer, p.Frame, p.About); err != nil {
		return
	}

	result = &Result{
		Payload: &FramesCreateResult{
			Workspace: w.Name,
			Frame:     f.Name,
		},
		Validation: vn,
	}
	return
}
