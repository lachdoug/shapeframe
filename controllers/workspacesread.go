package controllers

import "sf/models"

type WorkspacesReadParams struct {
	Workspace string
}

func WorkspacesRead(params *Params) (result *Result, err error) {
	var w *models.Workspace
	p := params.Payload.(*WorkspacesReadParams)

	uc := models.ResolveUserContext("Workspaces", "Workspace")
	if w, err = models.ResolveWorkspace(uc, p.Workspace,
		"Frames",
		"Directories",
		"Repositories.GitRepo",
	); err != nil {
		return
	}

	result = &Result{
		Payload: w.Read(),
	}
	return
}
