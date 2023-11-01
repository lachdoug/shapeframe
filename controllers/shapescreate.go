package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type ShapesCreateParams struct {
	Shaper string
	Name   string
	About  string
}

type ShapesCreateResult struct {
	Frame string
	Name  string
}

func ShapesCreate(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	params := &ShapesCreateParams{}
	utils.JsonUnmarshal(jparams, params)

	v = &app.Validation{}
	if params.Shaper == "" {
		v.Add("Shaper", "must not be blank")
	}
	if v.IsInvalid() {
		return
	}

	uc := models.UserContextNew()
	uc.Load("Frame")
	f := uc.Frame
	if f == nil {
		err = app.Error(nil, "no frame context")
		return
	}
	f.Load("Workspace.Directories.Workspace", "Shapes")

	s := models.ShapeNew(f, params.Name)
	s.Assign(map[string]any{
		"ShaperName": params.Shaper,
		"About":      params.About,
	})
	if err = s.SetShaper(); err != nil {
		return
	}
	if err = s.Create(); err != nil {
		return
	}

	result := &ShapesCreateResult{Frame: f.Name, Name: s.Name}
	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
