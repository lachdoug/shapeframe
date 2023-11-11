package controllers

import (
	"sf/app"
	"sf/models"
)

type FramesCreateParams struct {
	Workspace string
	Framer    string
	Name      string
	About     string
}

type FramesCreateResult struct {
	Workspace string
	Frame     string
}

func FramesCreate(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	var vn *app.Validation
	params := ParamsFor[FramesCreateParams](jparams)

	uc := models.ResolveUserContext("Workspaces.Frames")
	if w, err = models.ResolveWorkspace(uc, params.Workspace, "Framers"); err != nil {
		return
	}
	if f, vn, err = models.CreateFrame(w, params.Framer, params.Name, params.About); err != nil {
		return
	}

	result := &FramesCreateResult{
		Workspace: w.Name,
		Frame:     f.Name,
	}

	jbody = jbodyFor(result, vn)
	return
}
