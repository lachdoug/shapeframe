package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type FrameConfigurationsUpdateParams struct {
	Workspace string
	Name      string
	Config    map[string]any
}

type FrameConfigurationsUpdateResult struct {
	Workspace string
	Name      string
	Config    map[string]any
}

func FrameConfigurationsUpdate(jparams []byte) (jbody []byte, validation *app.Validation, err error) {
	params := &FrameConfigurationsUpdateParams{}
	utils.JsonUnmarshal(jparams, params)

	uc := models.UserContextNew()
	uc.Load("Workspaces")
	w := uc.WorkspaceFind(params.Workspace)
	if w == nil {
		err = app.Error(nil, "workspace %s does not exist", params.Workspace)
		return
	}
	f := w.FrameFind(params.Name)
	if f == nil {
		err = app.Error(nil, "frame %s does not exist", params.Name)
		return
	}

	f.Load("Workspace.Directories.Workspace")
	if err = f.SetFramer(); err != nil {
		return
	}
	f.Assign(map[string]any{
		"Config": params.Config,
	})
	if err = f.SaveConfiguration(); err != nil {
		return
	}

	result := &FrameConfigurationsUpdateResult{
		Workspace: w.Name,
		Name:      f.Name,
		Config:    f.ConfigValues(),
	}

	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
