package controllers

import (
	"sf/app/validations"
	"sf/models"
)

type ShapesCreateParams struct {
	Workspace string
	Frame     string
	Shaper    string
	Shape     string
	About     string
}

type ShapesCreateResult struct {
	Workspace string
	Frame     string
	Shape     string
}

func ShapesCreate(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	var vn *validations.Validation
	p := params.Payload.(*ShapesCreateParams)

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace", "Frame",
	)
	if w, err = models.ResolveWorkspace(uc, p.Workspace,
		"Frames",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, p.Frame,
		"Shapes", "Workspace.Shapers",
	); err != nil {
		return
	}
	if s, vn, err = models.CreateShape(f, p.Shaper, p.Shape, p.About); err != nil {
		return
	}

	result = &Result{
		Payload: &ShapesCreateResult{
			Workspace: w.Name,
			Frame:     f.Name,
			Shape:     s.Name,
		},
		Validation: vn,
	}
	return
}
