package controllers

import (
	"sf/models"
)

type FramesUpdateParams struct {
	Workspace string
	Frame     string
	Update    map[string]any
}

type FramesUpdateResult struct {
	Workspace string
	Frame     string
	From      *FramesUpdateResultDetails
	To        *FramesUpdateResultDetails
}

type FramesUpdateResultDetails struct {
	Name          string
	About         string
	Configuration []map[string]any
}

func FramesUpdate(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	params := ParamsFor[FramesUpdateParams](jparams)

	uc := models.ResolveUserContext(
		"Workspaces",
	)
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
			Name:          f.Name,
			About:         f.About,
			Configuration: f.Configuration.Details(),
		},
	}

	f.Assign(params.Update)
	f.Save()

	if err = f.Load("Configuration"); err != nil {
		return
	}

	result.To = &FramesUpdateResultDetails{
		Name:          f.Name,
		About:         f.About,
		Configuration: f.Configuration.Details(),
	}

	jbody = jbodyFor(result)
	return
}
