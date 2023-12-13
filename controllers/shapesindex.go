package controllers

import (
	"sf/models"
)

type ShapesIndexParams struct {
	Workspace string
	Frame     string
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

	uc := models.ResolveUserContext("Workspaces")
	if p.Workspace == "" {
		for _, w := range uc.Workspaces {
			w.Load("Frames.Shapes.Frame.Workspace")
			for _, f := range w.Frames {
				ss = append(ss, f.Shapes...)
			}
		}
	} else {
		if p.Frame == "" {
			if w, err = models.ResolveWorkspace(uc, p.Workspace,
				"Frames.Shapes.Frame.Workspace",
			); err != nil {
				return
			}
			for _, f := range w.Frames {
				ss = append(ss, f.Shapes...)
			}
		} else {
			if w, err = models.ResolveWorkspace(uc, p.Workspace,
				"Frames",
			); err != nil {
				return
			}
			if f, err = models.ResolveFrame(uc, w, p.Frame,
				"Shapes.Frame.Workspace",
			); err != nil {
				return
			}
			ss = f.Shapes
		}

	}

	r := []*ShapesIndexItemResult{}
	for _, s := range ss {
		r = append(r, &ShapesIndexItemResult{
			Workspace: s.Frame.Workspace.Name,
			Frame:     s.Frame.Name,
			Shape:     s.Name,
			Shaper:    s.ShaperName,
			About:     s.About,
			IsContext: uc.ShapeID == s.ID,
		})
	}

	result = &Result{Payload: r}
	return
}
