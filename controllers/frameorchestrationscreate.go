package controllers

import (
	"sf/app"
	"sf/models"
)

type FrameOrchestrationsCreateParams struct {
	Workspace string
	Frame     string
}

type FrameOrchestrationsCreateResult struct {
	Workspace string
	Frame     string
}

func FrameOrchestrationsCreate(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var f *models.Frame
	params := paramsFor[FrameOrchestrationsCreateParams](jparams)
	st := models.StreamCreate()

	vn = &app.Validation{}
	if params.Workspace == "" {
		vn.Add("Workspace", "must not be blank")
	}
	if params.Frame == "" {
		vn.Add("Frame", "must not be blank")
	}
	if vn.IsInvalid() {
		return
	}

	uc := models.ResolveUserContext(
		"Workspaces.Frames",
	)
	if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, params.Frame,
		"Configuration", "Shapes.Configuration"); err != nil {
		return
	}
	f.Orchestrate(st)

	result := &FrameOrchestrationsCreateResult{
		Workspace: w.Name,
		Frame:     f.Name,
	}

	jbody = jbodyFor(result, st)
	return
}
