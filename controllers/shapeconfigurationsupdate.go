package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type ShapeConfigurationsUpdateParams struct {
	Workspace string
	Frame     string
	Name      string
	Config    map[string]any
}

type ShapeConfigurationsUpdateResult struct {
	Workspace string
	Frame     string
	Name      string
	Config    map[string]any
}

func ShapeConfigurationsUpdate(jparams []byte) (jbody []byte, validation *app.Validation, err error) {
	params := &ShapeConfigurationsUpdateParams{}
	utils.JsonUnmarshal(jparams, params)

	uc := models.UserContextNew()
	uc.Load("Workspaces")
	w := uc.WorkspaceFind(params.Workspace)
	if w == nil {
		err = app.Error(nil, "workspace %s does not exist", params.Workspace)
		return
	}
	f := w.FrameFind(params.Frame)
	if f == nil {
		err = app.Error(nil, "frame %s does not exist", params.Frame)
		return
	}
	s := f.ShapeFind(params.Name)
	if s == nil {
		err = app.Error(nil, "shape %s does not exist", params.Name)
		return
	}

	s.Load("Frame.Workspace.Directories.Workspace")
	if err = s.SetShaper(); err != nil {
		return
	}
	s.Assign(map[string]any{
		"Config": params.Config,
	})
	if err = s.SaveConfiguration(); err != nil {
		return
	}

	result := &ShapeConfigurationsUpdateResult{
		Workspace: w.Name,
		Frame:     f.Name,
		Name:      s.Name,
		Config:    s.ConfigValues(),
	}

	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
