package controllers

import (
	"sf/app/validations"
	"sf/models"
)

type ShapesUpdateParams struct {
	Workspace string
	Frame     string
	Shape     string
	Updates   map[string]any
}

type ShapesUpdateResult struct {
	Workspace string
	Frame     string
	Shape     string
	From      *ShapesUpdateResultDetails
	To        *ShapesUpdateResultDetails
}

type ShapesUpdateResultDetails struct {
	Name          string
	About         string
	Configuration []map[string]any
}

func ShapesUpdate(params *Params) (result *Result, err error) {
	p := params.Payload.(*ShapesUpdateParams)
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	var vn *validations.Validation

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace", "Frame", "Shape",
	)
	if w, err = models.ResolveWorkspace(uc, p.Workspace,
		"Frames",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, p.Frame,
		"Shapes",
	); err != nil {
		return
	}
	if s, err = models.ResolveShape(uc, f, p.Shape); err != nil {
		return
	}

	r := &ShapesUpdateResult{
		Workspace: w.Name,
		Frame:     f.Name,
		Shape:     s.Name,
		From: &ShapesUpdateResultDetails{
			Name:  s.Name,
			About: s.About,
		},
	}

	vn = s.Update(p.Updates)

	r.To = &ShapesUpdateResultDetails{
		Name:  s.Name,
		About: s.About,
	}

	result = &Result{
		Payload:    r,
		Validation: vn,
	}
	return
}
