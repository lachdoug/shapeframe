package controllers

import (
	"sf/app"
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

func FramesUpdate(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var f *models.Frame
	params := paramsFor[FramesUpdateParams](jparams)

	vn = &app.Validation{}
	if params.Workspace == "" {
		vn.Add("Workspace", "must not be blank")
	}
	if params.Frame == "" {
		vn.Add("Frame", "must not be blank")
	}
	if vn.IsInvalid() {
		return
	}

	uc := models.ResolveUserContext(
		"Workspaces.Frames",
	)
	if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, params.Frame, "Configuration"); err != nil {
		return
	}

	result := &FramesUpdateResult{
		Workspace: w.Name,
		Frame:     f.Name,
		From: &FramesUpdateResultDetails{
			Name:          f.Name,
			About:         f.About,
			Configuration: f.Configuration.SettingsDetail(),
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
		Configuration: f.Configuration.SettingsDetail(),
	}

	jbody = jbodyFor(result)
	return
}
