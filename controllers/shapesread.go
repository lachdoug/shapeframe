package controllers

import (
	"sf/models"
)

type ShapesReadParams struct {
	Workspace string
	Frame     string
	Shape     string
}

func ShapesRead(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	params := ParamsFor[ShapesReadParams](jparams)

	uc := models.ResolveUserContext("Workspaces")
	if w, err = models.ResolveWorkspace(uc, params.Workspace,
		"Frames",
	); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, params.Frame,
		"Shapes",
	); err != nil {
		return
	}
	if s, err = models.ResolveShape(uc, f, params.Shape,
		"Configuration",
	); err != nil {
		return
	}

	result := s.Read()

	jbody = jbodyFor(result)
	return
}
