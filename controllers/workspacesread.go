package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

func WorkspacesRead(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	var w *models.Workspace
	uc := models.UserContextNew()
	uc.Load("Workspace.Directories.Workspace")
	if w = uc.Workspace; w == nil {
		err = app.Error(nil, "no workspace context")
		return
	}
	w.Load("Frames.Shapes")
	var result *models.WorkspaceInspector
	if result, err = w.Inspect(); err != nil {
		return
	}
	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
