package controllers

import (
	"sf/models"
)

type ShapesDeleteParams struct {
	Workspace string
	Frame     string
	Shape     string
}

type ShapesDeleteResult struct {
	Workspace string
	Frame     string
	Shape     string
}

func ShapesDelete(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	p := params.Payload.(*ShapesDeleteParams)

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
	if uc.Shape != nil && uc.Shape.ID == s.ID {
		uc.Clear("Shape")
	}
	s.Delete()

	result = &Result{
		Payload: &ShapesDeleteResult{
			Workspace: w.Name,
			Frame:     f.Name,
			Shape:     s.Name,
		},
	}
	return
}
