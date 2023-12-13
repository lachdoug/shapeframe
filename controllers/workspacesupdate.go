package controllers

import (
	"sf/app/validations"
	"sf/models"
)

type WorkspacesUpdateParams struct {
	Workspace string
	Updates   map[string]any
}

type WorkspacesUpdateResult struct {
	Workspace string
	From      *WorkspacesUpdateResultDetails
	To        *WorkspacesUpdateResultDetails
}

type WorkspacesUpdateResultDetails struct {
	Name  string
	About string
}

func WorkspacesUpdate(params *Params) (result *Result, err error) {
	p := params.Payload.(*WorkspacesUpdateParams)
	var w *models.Workspace
	var vn *validations.Validation

	uc := models.ResolveUserContext("Workspaces", "Workspace")
	if w, err = models.ResolveWorkspace(uc, p.Workspace); err != nil {
		return
	}

	r := &WorkspacesUpdateResult{
		Workspace: w.Name,
		From: &WorkspacesUpdateResultDetails{
			Name:  w.Name,
			About: w.About,
		},
	}

	vn = w.Update(p.Updates)

	r.To = &WorkspacesUpdateResultDetails{
		Name:  w.Name,
		About: w.About,
	}

	result = &Result{
		Payload:    r,
		Validation: vn,
	}
	return
}
