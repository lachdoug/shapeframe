package controllers

import (
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

func WorkspacesUpdate(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	params := ParamsFor[WorkspacesUpdateParams](jparams)

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
