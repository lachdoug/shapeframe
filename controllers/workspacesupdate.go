package controllers

import (
	"sf/app"
	"sf/models"
)

type WorkspacesUpdateParams struct {
	Workspace string
	Update    map[string]any
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

func WorkspacesUpdate(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	params := paramsFor[WorkspacesUpdateParams](jparams)

	vn = &app.Validation{}
	if params.Workspace == "" {
		vn.Add("Workspace", "must not be blank")
	}
	if vn.IsInvalid() {
		return
	}

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
		return
	}

	result := &WorkspacesUpdateResult{
		Workspace: w.Name,
		From: &WorkspacesUpdateResultDetails{
			Name:  w.Name,
			About: w.About,
		},
	}

	w.Assign(params.Update)
	w.Save()

	result.To = &WorkspacesUpdateResultDetails{
		Name:  w.Name,
		About: w.About,
	}

	jbody = jbodyFor(result)
	return
}
