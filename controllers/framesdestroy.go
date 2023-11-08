package controllers

import (
	"sf/app"
	"sf/models"
)

type FramesDestroyParams struct {
	Workspace string
	Frame     string
}

type FramesDestroyResult struct {
	Workspace string
	Frame     string
}

func FramesDestroy(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var f *models.Frame
	params := paramsFor[FramesDestroyParams](jparams)

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
		"Frame", "Shape",
		"Workspaces.Frames",
	)
	if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, params.Frame); err != nil {
		return
	}
	if uc.Frame != nil && uc.Frame.ID == f.ID {
		uc.Clear("Shape")
		uc.Clear("Frame")
	}
	f.Destroy()

	result := &FramesDestroyResult{
		Workspace: w.Name,
		Frame:     f.Name,
	}

	jbody = jbodyFor(result)
	return
}
