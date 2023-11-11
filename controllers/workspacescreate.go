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

func WorkspacesCreate(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var vn *app.Validation
	params := ParamsFor[WorkspacesCreateParams](jparams)

	uc := models.ResolveUserContext("Workspaces")
	if w, vn, err = models.CreateWorkspace(uc, params.Name, params.About); err != nil {
		return
	}

	result := &WorkspacesCreateResult{
		Workspace: w.Name,
	}

	jbody = jbodyFor(result, vn)
	return
}
