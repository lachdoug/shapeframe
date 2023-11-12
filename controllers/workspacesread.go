package controllers

import (
	"sf/models"
)

type WorkspacesReadParams struct {
	Workspace string
}

func WorkspacesRead(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	params := ParamsFor[WorkspacesReadParams](jparams)

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, params.Workspace,
		"Frames",
		"Directories",
		"Repositories.GitRepo",
	); err != nil {
		return
	}

	result := w.Read()

	jbody = jbodyFor(result)
	return
}
