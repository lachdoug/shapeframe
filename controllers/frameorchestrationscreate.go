package controllers

import (
	"sf/models"
	"sf/utils"
)

type FrameOrchestrationsCreateParams struct {
	Workspace string
	Frame     string
}

type FrameOrchestrationsCreateResult struct {
	Workspace string
	Frame     string
}

func FrameOrchestrationsCreate(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	params := ParamsFor[FrameOrchestrationsCreateParams](jparams)
	st := utils.StreamCreate()

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, params.Workspace,
		"Frames",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, params.Frame,
		"Configuration", "Shapes.Configuration",
	); err != nil {
		return
	}
	f.Orchestrate(st)

	result := &FrameOrchestrationsCreateResult{
		Workspace: w.Name,
		Frame:     f.Name,
	}

	jbody = jbodyFor(result, nil, st)
	return
}
