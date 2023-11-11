package controllers

import (
	"sf/app"
	"sf/models"
)

type ShapeConfigurationsUpdateParams struct {
	Workspace string
	Frame     string
	Shape     string
	Update    map[string]any
}

type ShapeConfigurationsUpdateResult struct {
	Workspace     string
	Frame         string
	Shape         string
	Configuration []map[string]any
}

func ShapeConfigurationsUpdate(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	var vn *app.Validation
	params := ParamsFor[ShapeConfigurationsUpdateParams](jparams)

	uc := models.ResolveUserContext("Workspaces.Frames.Shapes")
	if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, params.Frame); err != nil {
		return
	}
	if s, err = models.ResolveShape(uc, f, params.Shape, "Configuration"); err != nil {
		return
	}

	result := &ShapesUpdateResult{
		Workspace: w.Name,
		Frame:     f.Name,
		Shape:     s.Name,
		From: &ShapesUpdateResultDetails{
			Configuration: s.Configuration.Details(),
		},
	}

	if vn, err = s.Configuration.Update(params.Update); err != nil {
		return
	}

	result.To = &ShapesUpdateResultDetails{
		Configuration: s.Configuration.Details(),
	}

	jbody = jbodyFor(result, vn)
	return
}
