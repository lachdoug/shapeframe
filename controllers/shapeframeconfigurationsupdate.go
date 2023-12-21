package controllers

import (
	"sf/app/validations"
	"sf/models"
)

type ShapeFrameConfigurationsUpdateParams struct {
	Frame   string
	Shape   string
	Updates map[string]string
}

type ShapeFrameConfigurationsUpdateResult struct {
	Workspace string
	Frame     string
	Shape     string
	From      models.ConfigurationInspector
	To        models.ConfigurationInspector
}

func ShapeFrameConfigurationsUpdate(params *Params) (result *Result, err error) {
	p := params.Payload.(*ShapeFrameConfigurationsUpdateParams)
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

	r := &ShapeFrameConfigurationsUpdateResult{
		Workspace: w.Name,
		Frame:     f.Name,
		Shape:     s.Name,
		From:      *s.ShapeConfiguration.Frame.Inspect(),
	}

	vn = s.ShapeConfiguration.Frame.Update(p.Updates)

	r.To = *s.ShapeConfiguration.Frame.Inspect()

	result = &Result{
		Payload:    r,
		Validation: vn,
	}
	return
}
