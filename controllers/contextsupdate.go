package controllers

import (
	"sf/models"
)

type ContextsUpdateParams struct {
	Workspace string
	Frame     string
	Shape     string
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

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace", "Frame", "Shape",
	)

	r := &ContextsUpdateResult{From: &ContextsReadResult{
		Workspace: uc.WorkspaceName(),
		Frame:     uc.FrameName(),
		Shape:     uc.ShapeName(),
	}}

	if p.Workspace != "" {
		if w, err = models.ResolveWorkspace(uc, p.Workspace,
			"Frames",
		); err != nil {
			return
		}
		uc.Workspace = w
	} else {
		uc.Clear("Workspace")
	}
	if p.Frame != "" {
		if f, err = models.ResolveFrame(uc, w, p.Frame,
			"Shapes",
		); err != nil {
			return
		}
		uc.Frame = f
	} else {
		uc.Clear("Frame")
	}
	if p.Shape != "" {
		if s, err = models.ResolveShape(uc, f, p.Shape); err != nil {
			return
		}
		uc.Shape = s
	} else {
		uc.Clear("Shape")
	}
	uc.Save()

	r.To = &ContextsReadResult{
		Workspace: uc.WorkspaceName(),
		Frame:     uc.FrameName(),
		Shape:     uc.ShapeName(),
	}
	result = &Result{Payload: r}
	return
}
