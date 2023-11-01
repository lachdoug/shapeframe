package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type ShapesDestroyParams struct {
	Workspace string
	Frame     string
	Name      string
}

type ShapesDestroyResult struct {
	Frame string
	Name  string
}

func ShapesDestroy(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	params := &ShapesDestroyParams{}
	utils.JsonUnmarshal(jparams, params)

	v = &app.Validation{}
	if params.Workspace == "" {
		v.Add("Workspace", "must not be blank")
	}
	if params.Frame == "" {
		v.Add("Frame", "must not be blank")
	}
	if params.Name == "" {
		v.Add("Name", "must not be blank")
	}
	if v.IsInvalid() {
		return
	}

	uc := models.UserContextNew()
	w := uc.WorkspaceFind(params.Workspace)
	if w == nil {
		err = app.Error(nil, "workspace %s does not exist", params.Workspace)
		return
	}
	f := w.FrameFind(params.Frame)
	if f == nil {
		err = app.Error(nil, "frame %s does not exist in workspace %s", params.Frame, params.Workspace)
		return
	}
	s := f.ShapeFind(params.Name)
	if s == nil {
		err = app.Error(nil, "shape %s does not exist in frame %s", params.Name, f.Name)
		return
	}

	uc.Load("Shape")
	if uc.Shape.ID == s.ID {
		uc.Clear("Shape")
	}

	s.Destroy()

	result := &ShapesDestroyResult{Frame: f.Name, Name: s.Name}
	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
