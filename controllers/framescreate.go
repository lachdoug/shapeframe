package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type FramesCreateParams struct {
	Workspace string
	Framer    string
	Name      string
	About     string
}

type FramesCreateResult struct {
	Workspace string
	Name      string
}

func FramesCreate(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	params := &FramesCreateParams{}
	utils.JsonUnmarshal(jparams, params)

	v = &app.Validation{}
	if params.Workspace == "" {
		v.Add("Workspace", "must not be blank")
	}
	if params.Framer == "" {
		v.Add("Framer", "must not be blank")
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
	w.Load("Directories.Workspace", "Frames")

	f := models.FrameNew(w, params.Name)
	f.Assign(map[string]any{
		"FramerName": params.Framer,
		"About":      params.About,
	})
	if err = f.SetFramer(); err != nil {
		return
	}
	if err = f.Create(); err != nil {
		return
	}

	result := &FramesCreateResult{Workspace: w.Name, Name: f.Name}
	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
