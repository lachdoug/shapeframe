package controllers

import (
	"sf/app/validations"
	"sf/models"
)

type FramesUpdateParams struct {
	Frame   string
	Updates map[string]string
}

type FramesUpdateResult struct {
	Workspace string
	Frame     string
	From      *FramesUpdateResultDetails
	To        *FramesUpdateResultDetails
}

type FramesUpdateResultDetails struct {
	Name  string
	About string
}

func FramesUpdate(params *Params) (result *Result, err error) {
	p := params.Payload.(*FramesUpdateParams)
	var w *models.Workspace
	var f *models.Frame
	var vn *validations.Validation

	if w, err = models.ResolveWorkspace(
		"Frames", "Frame",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(w, p.Frame,
		"Configuration",
	); err != nil {
		return
	}

	r := &FramesUpdateResult{
		Workspace: w.Name,
		Frame:     f.Name,
		From: &FramesUpdateResultDetails{
			Name:  f.Name,
			About: f.About,
		},
	}

	vn = f.Update(p.Updates)

	r.To = &FramesUpdateResultDetails{
		Name:  f.Name,
		About: f.About,
	}

	result = &Result{
		Payload:    r,
		Validation: vn,
	}
	return
}
