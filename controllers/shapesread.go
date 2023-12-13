package controllers

import (
	"sf/models"
)

type ShapesReadParams struct {
	Workspace string
	Frame     string
	Shape     string
}

func ShapesRead(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	p := params.Payload.(*ShapesReadParams)

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
	if s, err = models.ResolveShape(uc, f, p.Shape,
		"Configuration",
	); err != nil {
		return
	}

	result = &Result{
		Payload: s.Read(),
	}
	return
}
