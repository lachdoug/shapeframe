package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type FramesUpdateParams struct {
	Workspace string
	Frame     string
	Name      string
	About     string
}

type FramesUpdateResult struct {
	From *FramesUpdateResultDetails
	To   *FramesUpdateResultDetails
}

type FramesUpdateResultDetails struct {
	Name  string
	About string
}

func FramesUpdate(jparams []byte) (jbody []byte, validation *app.Validation, err error) {
	params := &FramesUpdateParams{}
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

	result := &FramesUpdateResult{
		From: &FramesUpdateResultDetails{
			Name:  f.Name,
			About: f.About,
		},
	}

	f.Assign(map[string]any{
		"Name":  params.Name,
		"About": params.About,
	})
	if err = f.Save(); err != nil {
		return
	}

	result.To = &FramesUpdateResultDetails{
		Name:  f.Name,
		About: f.About,
	}

	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
