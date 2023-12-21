package controllers

import (
	"sf/app/validations"
	"sf/models"
)

type ShapeConfigurationsUpdateParams struct {
	Frame   string
	Shape   string
	Updates map[string]string
}

type ShapeConfigurationsUpdateResult struct {
	Workspace string
	Frame     string
	Shape     string
	From      models.ConfigurationInspector
	To        models.ConfigurationInspector
}

func ShapeConfigurationsUpdate(params *Params) (result *Result, err error) {
	p := params.Payload.(*ShapeConfigurationsUpdateParams)
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	var vn *validations.Validation

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

	r := &ShapeConfigurationsUpdateResult{
		Workspace: w.Name,
		Frame:     f.Name,
		Shape:     s.Name,
		From:      *s.ShapeConfiguration.Shape.Inspect(),
	}

	vn = s.ShapeConfiguration.Shape.Update(p.Updates)

	r.To = *s.ShapeConfiguration.Shape.Inspect()

	result = &Result{
		Payload:    r,
		Validation: vn,
	}
	return
}
