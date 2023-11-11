package controllers

import (
	"sf/app"
	"sf/models"
)

type ShapesUpdateParams struct {
	Workspace string
	Frame     string
	Shape     string
	Update    map[string]any
}

type ShapesUpdateResult struct {
	Workspace string
	Frame     string
	Shape     string
	From      *ShapesUpdateResultDetails
	To        *ShapesUpdateResultDetails
}

type ShapesUpdateResultDetails struct {
	Name          string
	About         string
	Configuration []map[string]any
}

func ShapesUpdate(jparams []byte) (jbody []byte, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	var vn *app.Validation
	params := ParamsFor[ShapesUpdateParams](jparams)

	uc := models.ResolveUserContext(
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

	result := &ShapesUpdateResult{
		Workspace: w.Name,
		Frame:     f.Name,
		Shape:     s.Name,
		From: &ShapesUpdateResultDetails{
			Name:  s.Name,
			About: s.About,
		},
	}

	vn = s.Update(params.Update)

	result.To = &ShapesUpdateResultDetails{
		Name:  s.Name,
		About: s.About,
	}

	jbody = jbodyFor(result, vn)
	return
}
