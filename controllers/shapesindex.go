package controllers

import (
	"sf/models"
)

type ShapesIndexParams struct {
	Workspace string // Limit list to workspace
}

type ShapesIndexItemResult struct {
	Workspace string
	Frame     string
	Shaper    string
	Name      string
	About     string
	IsContext bool
}

func ShapesIndex(jparams []byte) (jbody []byte, err error) {
	var ss []*models.Shape
	params := ParamsFor[ShapesIndexParams](jparams)

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace",
	)
	if params.Workspace == "" {
		for _, w := range uc.Workspaces {
			w.Load("Frames.Shapes.Frame.Workspace")
			for _, f := range w.Frames {
				ss = append(ss, f.Shapes...)
			}
		}
	} else {
		var w *models.Workspace
		if w, err = models.ResolveWorkspace(uc, params.Workspace,
			"Frames.Shapes.Frame.Workspace",
		); err != nil {
			return
		}
		for _, f := range w.Frames {
			ss = append(ss, f.Shapes...)
		}
	}

	result := []*ShapesIndexItemResult{}
	for _, s := range ss {
		result = append(result, &ShapesIndexItemResult{
			Workspace: s.Frame.Workspace.Name,
			Frame:     s.Frame.Name,
			Shaper:    s.ShaperName,
			Name:      s.Name,
			About:     s.About,
			IsContext: uc.ShapeID == s.ID,
		})
	}

	jbody = jbodyFor(result)
	return
}
