package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type FramesIndexParams struct {
	Workspace bool // Limit list to workspace context
}

type FramesIndexItemResult struct {
	Workspace string
	Framer    string
	Name      string
	About     string
	IsContext bool
}

func FramesIndex(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	var fs []*models.Frame

	params := &FramesIndexParams{}
	utils.JsonUnmarshal(jparams, params)

	uc := models.UserContextNew()
	if params.Workspace {
		uc.Load("Workspace.Frames.Workspace")
		if uc.Workspace == nil {
			err = app.Error(nil, "no workspace context")
			return
		}
		w := uc.Workspace
		fs = w.Frames
	} else {
		uc.Load("Workspaces.Frames.Workspace")
		fs = uc.Frames()
	}

	result := []*FramesIndexItemResult{}
	for _, f := range fs {
		result = append(result, &FramesIndexItemResult{
			Workspace: f.Workspace.Name,
			Framer:    f.FramerName,
			Name:      f.Name,
			About:     f.About,
			IsContext: uc.FrameID == f.ID,
		})
	}

	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
