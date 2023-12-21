package controllers

import (
	"sf/models"
)

type ShapesIndexParams struct {
	Frame string
}

type ShapesIndexItemResult struct {
	Workspace string
	Frame     string
	Shape     string
	Shaper    string
	About     string
	IsContext bool
}

func ShapesIndex(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var f *models.Frame
	var ss []*models.Shape
	if params.Payload == nil {
		params.Payload = &ShapesIndexParams{}
	}
	p := params.Payload.(*ShapesIndexParams)

	if p.Frame == "" {
		if w, err = models.ResolveWorkspace(
			"Frames.Shapes.Frame.Workspace",
		); err != nil {
			return
		}
		for _, f := range w.Frames {
			ss = append(ss, f.Shapes...)
		}
	} else {
		if w, err = models.ResolveWorkspace(
			"Frames", "Frame",
		); err != nil {
			return
		}
		if f, err = models.ResolveFrame(w, p.Frame,
			"Shapes.Frame.Workspace",
		); err != nil {
			return
		}
		ss = f.Shapes
	}

	r := []*ShapesIndexItemResult{}
	for _, s := range ss {
		r = append(r, &ShapesIndexItemResult{
			Workspace: s.Frame.Workspace.Name,
			Frame:     s.Frame.Name,
			Shape:     s.Name,
			Shaper:    s.ShaperName,
			About:     s.About,
			IsContext: w.ShapeID == s.ID,
		})
	}

	result = &Result{Payload: r}
	return
}
