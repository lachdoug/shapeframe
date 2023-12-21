package controllers

import (
	"sf/models"
)

type ShapesDeleteParams struct {
	Frame string
	Shape string
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

	if w, err = models.ResolveWorkspace(
		"Frames", "Frame", "Shape",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(w, p.Frame,
		"Shapes",
	); err != nil {
		return
	}
	if s, err = models.ResolveShape(w, f, p.Shape); err != nil {
		return
	}
	if w.Shape != nil && w.Shape.ID == s.ID {
		w.Clear("Shape")
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
