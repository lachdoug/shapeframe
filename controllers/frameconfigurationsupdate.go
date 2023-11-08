package controllers

import (
	"sf/app"
	"sf/models"
)

type FrameConfigurationsUpdateParams struct {
	Workspace     string
	Frame         string
	Configuration map[string]any
}

type FrameConfigurationsUpdateResult struct {
	Workspace     string
	Frame         string
	Configuration []map[string]any
}

func FrameConfigurationsUpdate(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var f *models.Frame
	params := paramsFor[FrameConfigurationsUpdateParams](jparams)

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

	uc := models.ResolveUserContext("Workspaces.Frames")
	if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, params.Frame); err != nil {
		return
	}

	f.Assign(map[string]any{
		"Configuration": params.Configuration,
	})
	f.Save()

	if err = f.Load("Configuration"); err != nil {
		return
	}
	if err = f.Configuration.Validate(); err != nil {
		return
	}

	result := &FrameConfigurationsUpdateResult{
		Workspace:     w.Name,
		Frame:         f.Name,
		Configuration: f.Configuration.SettingsDetail(),
	}

	jbody = jbodyFor(result)
	return
}
