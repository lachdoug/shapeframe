package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type ShapesUpdateParams struct {
	Workspace string
	Frame     string
	Shape     string
	Name      string
	About     string
}

type ShapesUpdateResult struct {
	From *ShapesUpdateResultDetails
	To   *ShapesUpdateResultDetails
}

type ShapesUpdateResultDetails struct {
	Name  string
	About string
}

func ShapesUpdate(jparams []byte) (jbody []byte, validation *app.Validation, err error) {
	params := &ShapesUpdateParams{}
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
	s := f.ShapeFind(params.Shape)
	if s == nil {
		err = app.Error(nil, "shape %s does not exist", params.Shape)
		return
	}

	result := &ShapesUpdateResult{
		From: &ShapesUpdateResultDetails{
			Name:  s.Name,
			About: s.About,
		},
	}

	s.Assign(map[string]any{
		"Name":  params.Name,
		"About": params.About,
	})
	if err = s.Save(); err != nil {
		return
	}

	result.To = &ShapesUpdateResultDetails{
		Name:  s.Name,
		About: s.About,
	}

	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
