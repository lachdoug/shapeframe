package controllers

import (
	"sf/app/validations"
	"sf/models"
)

type FrameConfigurationsUpdateParams struct {
	Workspace string
	Frame     string
	Updates   map[string]string
}

type FrameConfigurationsUpdateResult struct {
	Workspace string
	Frame     string
	From      []map[string]string
	To        []map[string]string
}

func FrameConfigurationsUpdate(params *Params) (result *Result, err error) {
	p := params.Payload.(*FrameConfigurationsUpdateParams)
	var w *models.Workspace
	var f *models.Frame
	var vn *validations.Validation

	uc := models.ResolveUserContext("Workspaces", "Workspace", "Frame")
	if w, err = models.ResolveWorkspace(uc, p.Workspace,
		"Frames",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, p.Frame,
		"Configuration",
	); err != nil {
		return
	}

	r := &FrameConfigurationsUpdateResult{
		Workspace: w.Name,
		Frame:     f.Name,
		From:      f.Configuration.Info(),
	}

	vn = f.Configuration.Update(p.Updates)

	r.To = f.Configuration.Info()

	result = &Result{
		Payload:    r,
		Validation: vn,
	}
	return
}
