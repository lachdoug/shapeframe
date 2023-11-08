package controllers

import (
	"sf/app"
	"sf/models"
)

type WorkspacesReadParams struct {
	Workspace string
}

func WorkspacesRead(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	params := paramsFor[WorkspacesReadParams](jparams)

	uc := models.ResolveUserContext(
		"Workspaces",
	)
	if w, err = models.ResolveWorkspace(uc, params.Workspace,
		"Frames.Configuration",
		"Frames.Shapes.Configuration",
		"Directories.Workspace",
		"Directories.Framers",
		"Directories.Shapers",
		"Repositories.Framers",
		"Repositories.Shapers",
	); err != nil {
		return
	}

	result := w.Inspect()

	jbody = jbodyFor(result)
	return
}
