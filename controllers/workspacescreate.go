package controllers

import (
	"sf/app/validations"
	"sf/models"
)

type WorkspacesCreateParams struct {
	Directory string
	Name      string
	About     string
}

type WorkspacesCreateResult struct {
	Directory string
	Workspace string
}

func WorkspacesCreate(params *Params) (result *Result, err error) {
	p := params.Payload.(*WorkspacesCreateParams)
	var w *models.Workspace
	var vn *validations.Validation

	if w, vn, err = models.CreateWorkspace(p.Directory, p.Name, p.About); err != nil {
		return
	}

	result = &Result{
		Payload: &WorkspacesCreateResult{
			Workspace: w.Name,
			Directory: w.Directory(),
		},
		Validation: vn,
	}
	return
}
