package controllers

import (
	"sf/models"
)

type ContextsUpdateParams struct {
	Workspace string
	Frame     string
	Shape     string
}

type ContextsUpdateResult struct {
	From *models.UserContextInspector
	To   *models.UserContextInspector
}

func ContextsUpdate(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	params := ParamsFor[ContextsUpdateParams](jparams)

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace", "Frame", "Shape",
	)

	result := &ContextsUpdateResult{From: uc.Inspect()}

	if params.Workspace != "" {
		if w, err = models.ResolveWorkspace(uc, params.Workspace,
			"Frames",
		); err != nil {
			return
		}
		uc.Workspace = w
	} else {
		uc.Clear("Workspace")
	}
	if params.Frame != "" {
		if f, err = models.ResolveFrame(uc, w, params.Frame,
			"Shapes",
		); err != nil {
			return
		}
		uc.Frame = f
	} else {
		uc.Clear("Frame")
	}
	if params.Shape != "" {
		if s, err = models.ResolveShape(uc, f, params.Shape); err != nil {
			return
		}
		uc.Shape = s
	} else {
		uc.Clear("Shape")
	}
	uc.Save()

	result.To = uc.Inspect()

	jbody = jbodyFor(result)
	return
}
