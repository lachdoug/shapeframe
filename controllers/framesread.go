package controllers

import (
	"sf/models"
)

type FramesReadParams struct {
	Workspace string
	Frame     string
}

func FramesRead(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	params := ParamsFor[FramesReadParams](jparams)

	uc := models.ResolveUserContext("Workspaces.Frames")
	if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, params.Frame,
		"Workspace",
		"Configuration",
		"Shapes",
		"Parent",
	); err != nil {
		return
	}

	result := f.Read()

	jbody = jbodyFor(result)
	return
}
