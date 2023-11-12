package controllers

import (
	"sf/app"
	"sf/models"
)

type FrameConfigurationsUpdateParams struct {
	Workspace string
	Frame     string
	Update    map[string]any
}

type FrameConfigurationsUpdateResult struct {
	Workspace     string
	Frame         string
	Configuration []map[string]any
}

func FrameConfigurationsUpdate(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	var vn *app.Validation
	params := ParamsFor[FrameConfigurationsUpdateParams](jparams)

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, params.Workspace,
		"Frames",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, params.Frame,
		"Configuration",
	); err != nil {
		return
	}

	result := &FramesUpdateResult{
		Workspace: w.Name,
		Frame:     f.Name,
		From: &FramesUpdateResultDetails{
			Configuration: f.Configuration.Details(),
		},
	}

	if vn, err = f.Configuration.Update(params.Update); err != nil {
		return
	}

	result.To = &FramesUpdateResultDetails{
		Configuration: f.Configuration.Details(),
	}

	jbody = jbodyFor(result, vn)
	return
}
