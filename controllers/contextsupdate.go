package controllers

import (
	"sf/models"
)

type ContextsUpdateParams struct {
	Frame string
	Shape string
}

type ContextsUpdateResult struct {
	From *ContextsReadResult
	To   *ContextsReadResult
}

func ContextsUpdate(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	p := params.Payload.(*ContextsUpdateParams)

	if w, err = models.ResolveWorkspace(
		"Frames", "Frame", "Shape",
	); err != nil {
		return
	}

	r := &ContextsUpdateResult{From: &ContextsReadResult{
		Frame: w.FrameName(),
		Shape: w.ShapeName(),
	}}

	if p.Frame != "" {
		if f, err = models.ResolveFrame(w, p.Frame,
			"Shapes",
		); err != nil {
			return
		}
		w.Frame = f
	} else {
		w.Clear("Frame")
	}
	if p.Shape != "" {
		if s, err = models.ResolveShape(w, f, p.Shape); err != nil {
			return
		}
		w.Shape = s
	} else {
		w.Clear("Shape")
	}
	w.Save()

	r.To = &ContextsReadResult{
		Frame: w.FrameName(),
		Shape: w.ShapeName(),
	}
	result = &Result{Payload: r}
	return
}
