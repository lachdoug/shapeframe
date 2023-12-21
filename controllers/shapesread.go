package controllers

import (
	"sf/models"
)

type ShapesReadParams struct {
	Frame string
	Shape string
}

func ShapesRead(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	p := params.Payload.(*ShapesReadParams)

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
	if s, err = models.ResolveShape(w, f, p.Shape,
		"Configuration.Info",
	); err != nil {
		return
	}

	result = &Result{
		Payload: s.Read(),
	}
	return
}
