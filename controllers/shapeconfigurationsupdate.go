package controllers

import (
	"sf/app"
	"sf/models"
)

type ShapeConfigurationsUpdateParams struct {
	Workspace     string
	Frame         string
	Shape         string
	Configuration map[string]any
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
	if s, err = models.ResolveShape(uc, f, params.Shape); err != nil {
		return
	}

	s.Assign(map[string]any{
		"Configuration": params.Configuration,
	})
	s.Save()

	if err = s.Load("Configuration"); err != nil {
		return
	}
	if err = s.Configuration.Validate(); err != nil {
		return
	}

	result := &ShapeConfigurationsUpdateResult{
		Workspace:     w.Name,
		Frame:         f.Name,
		Shape:         s.Name,
		Configuration: s.Configuration.SettingsDetail(),
	}

	jbody = jbodyFor(result)
	return
}
