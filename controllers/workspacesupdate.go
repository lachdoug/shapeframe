package controllers

import (
	"sf/app/validations"
	"sf/models"
)

type WorkspacesUpdateParams struct {
	Updates map[string]string
}

type WorkspacesUpdateResult struct {
	From *WorkspacesUpdateResultDetails
	To   *WorkspacesUpdateResultDetails
}

type WorkspacesUpdateResultDetails struct {
	Name  string
	About string
}

func WorkspacesUpdate(params *Params) (result *Result, err error) {
	p := params.Payload.(*WorkspacesUpdateParams)
	var w *models.Workspace
	var vn *validations.Validation

	if w, err = models.ResolveWorkspace(); err != nil {
		return
	}

	r := &WorkspacesUpdateResult{
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
