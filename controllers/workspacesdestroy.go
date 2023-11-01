package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type WorkspacesDestroyParams struct {
	Name string
}

type WorkspacesDestroyResult struct {
	Name string
}

func WorkspacesDestroy(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	params := &WorkspacesDestroyParams{}
	utils.JsonUnmarshal(jparams, params)

	v = &app.Validation{}
	if params.Name == "" {
		v.Add("Name", "must not be blank")
	}
	if v.IsInvalid() {
		return
	}

	uc := models.UserContextNew()
	w := uc.WorkspaceFind(params.Name)
	if w == nil {
		err = app.Error(nil, "workspace %s does not exist", params.Name)
		return
	}

	uc.Load("Workspace")
	if uc.Workspace.ID == w.ID {
		uc.Clear("Shape")
		uc.Clear("Frame")
		uc.Clear("Workspace")
	}

	w.Destroy()

	result := &WorkspacesDestroyResult{
		Name: w.Name,
	}

	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
