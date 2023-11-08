package controllers

import (
	"sf/app"
	"sf/models"
)

type WorkspacesDestroyParams struct {
	Workspace string
}

type WorkspacesDestroyResult struct {
	Workspace string
}

func WorkspacesDestroy(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	params := paramsFor[WorkspacesDestroyParams](jparams)

	vn = &app.Validation{}
	if params.Workspace == "" {
		vn.Add("Workspace", "must not be blank")
	}
	if vn.IsInvalid() {
		return
	}

	uc := models.ResolveUserContext(
		"Workspace", "Frame", "Shape",
		"Workspaces",
	)
	if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
		return
	}
	if uc.Workspace != nil && uc.Workspace.ID == w.ID {
		uc.Clear("Shape")
		uc.Clear("Frame")
		uc.Clear("Workspace")
	}
	w.Destroy()

	result := &WorkspacesDestroyResult{
		Workspace: w.Name,
	}

	jbody = jbodyFor(result)
	return
}
