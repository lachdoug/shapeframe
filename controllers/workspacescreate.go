package controllers

import (
	"sf/app/validations"
	"sf/models"
)

type WorkspacesCreateParams struct {
	Workspace string
	About     string
}

type WorkspacesCreateResult struct {
	Workspace string
}

func WorkspacesCreate(params *Params) (result *Result, err error) {
	p := params.Payload.(*WorkspacesCreateParams)
	var w *models.Workspace
	var vn *validations.Validation

	uc := models.ResolveUserContext("Workspaces")
	if w, vn, err = models.CreateWorkspace(uc, p.Workspace, p.About); err != nil {
		return
	}

	result = &Result{
		Payload: &WorkspacesCreateResult{
			Workspace: w.Name,
		},
		Validation: vn,
	}
	return
}
