package controllers

import (
	"sf/models"
)

type WorkspaceInspectsReadParams struct {
	Workspace string
}

func WorkspaceInspectsRead(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var wi *models.WorkspaceInspector
	params := ParamsFor[WorkspaceInspectsReadParams](jparams)

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

	if wi, err = w.Inspect(); err != nil {
		return
	}

	jbody = jbodyFor(wi)
	return
}
