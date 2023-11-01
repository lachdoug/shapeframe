package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type ContextsUpdateParams struct {
	Workspace string
	Frame     string
	Shape     string
}

type ContextsUpdateResult struct {
	Exit  *models.UserContextInspector
	Enter *models.UserContextInspector
}

func ContextsUpdate(jparams []byte) (jbody []byte, validation *app.Validation, err error) {
	params := &ContextsUpdateParams{}
	utils.JsonUnmarshal(jparams, params)

	uc := models.UserContextNew()
	uc.Load("Workspace", "Frame", "Shape")

	result := &ContextsUpdateResult{Exit: uc.Inspect()}

	s := &models.Shape{}
	f := &models.Frame{}
	w := &models.Workspace{}

	if params.Workspace != "" {
		w = uc.WorkspaceFind(params.Workspace)
		if w == nil {
			err = app.Error(nil, "workspace %s does not exist", params.Workspace)
			return
		} else {
			uc.Workspace = w
		}
	}
	if params.Frame != "" {
		f = w.FrameFind(params.Frame)
		if f == nil {
			err = app.Error(nil, "frame %s does not exist in workspace %s", params.Frame, params.Workspace)
			return
		} else {
			uc.Frame = f
		}
	}
	if params.Shape != "" {
		s = f.ShapeFind(params.Shape)
		if s == nil {
			err = app.Error(nil, "shape %s does not exist in frame %s", params.Shape, params.Frame)
			return
		} else {
			uc.Shape = s
		}
	}

	uc.Save()

	result.Enter = uc.Inspect()

	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
