package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type FramesDestroyParams struct {
	Workspace string
	Name      string
}

type FramesDestroyResult struct {
	Workspace string
	Name      string
}

func FramesDestroy(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	params := &FramesDestroyParams{}
	utils.JsonUnmarshal(jparams, params)

	v = &app.Validation{}
	if params.Workspace == "" {
		v.Add("Workspace", "must not be blank")
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
	f := w.FrameFind(params.Name)
	if f == nil {
		err = app.Error(nil, "frame %s does not exist in workspace %s", params.Name, w.Name)
		return
	}

	uc.Load("Frame")
	if uc.Frame.ID == f.ID {
		uc.Clear("Shape")
		uc.Clear("Frame")
	}

	f.Destroy()

	result := &FramesDestroyResult{Workspace: w.Name, Name: f.Name}
	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
