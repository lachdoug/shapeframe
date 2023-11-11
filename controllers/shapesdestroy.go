package controllers

import (
	"sf/models"
)

type ShapesDestroyParams struct {
	Workspace string
	Frame     string
	Shape     string
}

type ShapesDestroyResult struct {
	Workspace string
	Frame     string
	Shape     string
}

func ShapesDestroy(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	params := ParamsFor[ShapesDestroyParams](jparams)

	uc := models.ResolveUserContext(
		"Shape",
		"Workspaces.Frames.Shapes",
	)
	if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, params.Frame); err != nil {
		return
	}
	if s, err = models.ResolveShape(uc, f, params.Shape); err != nil {
		return
	}
	if uc.Shape != nil && uc.Shape.ID == s.ID {
		uc.Clear("Shape")
	}
	s.Destroy()

	result := &ShapesDestroyResult{
		Frame: f.Name,
		Shape: s.Name,
	}

	jbody = jbodyFor(result)
	return
}
