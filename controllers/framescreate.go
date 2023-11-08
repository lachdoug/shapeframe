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

func FramesCreate(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var f *models.Frame
	params := paramsFor[FramesCreateParams](jparams)

	vn = &app.Validation{}
	if params.Workspace == "" {
		vn.Add("Workspace", "must not be blank")
	}
	if params.Framer == "" {
		vn.Add("Framer", "must not be blank")
	}
	if vn.IsInvalid() {
		return
	}

	uc := models.ResolveUserContext("Workspaces.Frames")
	if w, err = models.ResolveWorkspace(uc, params.Workspace, "Framers"); err != nil {
		return
	}
	if f, err = models.CreateFrame(w, params.Framer, params.Name, params.About); err != nil {
		return
	}

	result := &FramesCreateResult{
		Workspace: w.Name,
		Frame:     f.Name,
	}

	jbody = jbodyFor(result)
	return
}
