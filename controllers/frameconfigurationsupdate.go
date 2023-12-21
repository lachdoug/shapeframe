package controllers

import (
	"sf/app/validations"
	"sf/models"
)

type FrameConfigurationsUpdateParams struct {
	Frame   string
	Updates map[string]string
}

type FrameConfigurationsUpdateResult struct {
	Workspace string
	Frame     string
	From      models.ConfigurationInspector
	To        models.ConfigurationInspector
}

func FrameConfigurationsUpdate(params *Params) (result *Result, err error) {
	p := params.Payload.(*FrameConfigurationsUpdateParams)
	var w *models.Workspace
	var f *models.Frame
	var vn *validations.Validation

	if w, err = models.ResolveWorkspace(
		"Frames", "Frame",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(w, p.Frame,
		"Configuration.Info",
	); err != nil {
		return
	}

	r := &FrameConfigurationsUpdateResult{
		Workspace: w.Name,
		Frame:     f.Name,
		From:      *f.Configuration.Inspect(),
	}

	vn = f.Configuration.Update(p.Updates)

	r.To = *f.Configuration.Inspect()

	result = &Result{
		Payload:    r,
		Validation: vn,
	}
	return
}
