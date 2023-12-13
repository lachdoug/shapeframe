package controllers

import (
	"sf/app/validations"
	"sf/models"
)

type ShapeConfigurationsUpdateParams struct {
	Workspace string
	Frame     string
	Shape     string
	Updates   map[string]string
}

type ShapeConfigurationsUpdateResult struct {
	Workspace string
	Frame     string
	Shape     string
	From      []map[string]string
	To        []map[string]string
}

func ShapeConfigurationsUpdate(params *Params) (result *Result, err error) {
	p := params.Payload.(*ShapeConfigurationsUpdateParams)
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
	if s, err = models.ResolveShape(uc, f, p.Shape,
		"Configuration",
	); err != nil {
		return
	}

	r := &ShapeConfigurationsUpdateResult{
		Workspace: w.Name,
		Frame:     f.Name,
		Shape:     s.Name,
		From:      s.ShapeConfiguration.Info(),
	}

	vn = s.ShapeConfiguration.Update(p.Updates)

	r.To = s.ShapeConfiguration.Info()

	result = &Result{
		Payload:    r,
		Validation: vn,
	}
	return
}
