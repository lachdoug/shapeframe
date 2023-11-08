package controllers

import (
	"sf/app"
	"sf/models"
)

type WorkspacesCreateParams struct {
	Name  string
	About string
}

type WorkspacesCreateResult struct {
	Workspace string
}

func WorkspacesCreate(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	params := paramsFor[WorkspacesCreateParams](jparams)

	vn = &app.Validation{}
	if params.Name == "" {
		vn.Add("Name", "must not be blank")
	}
	if vn.IsInvalid() {
		return
	}

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.CreateWorkspace(uc, params.Name, params.About); err != nil {
		return
	}

	result := &WorkspacesCreateResult{
		Workspace: w.Name,
	}

	jbody = jbodyFor(result)
	return
}
