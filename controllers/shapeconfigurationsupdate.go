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

func ShapeConfigurationsUpdate(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	params := paramsFor[ShapeConfigurationsUpdateParams](jparams)

	vn = &app.Validation{}
	if params.Workspace == "" {
		vn.Add("Workspace", "must not be blank")
	}
	if params.Frame == "" {
		vn.Add("Frame", "must not be blank")
	}
	if params.Shape == "" {
		vn.Add("Shape", "must not be blank")
	}
	if vn.IsInvalid() {
		return
	}

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
			Configuration: s.Configuration.SettingsDetail(),
		},
	}

	s.Assign(params.Update)
	s.Save()

	if err = s.Load("Configuration"); err != nil {
		return
	}
	if vn = s.Configuration.Validate(); vn != nil {
		return
	}

	result.To = &ShapesUpdateResultDetails{
		Configuration: s.Configuration.SettingsDetail(),
	}

	jbody = jbodyFor(result)
	return
}
