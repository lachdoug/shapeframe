package controllers

import "sf/models"

func WorkspacesRead(params *Params) (result *Result, err error) {
	var w *models.Workspace

	if w, err = models.ResolveWorkspace(
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
