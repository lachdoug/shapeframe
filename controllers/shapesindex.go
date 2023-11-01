package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type ShapesIndexParams struct {
	Workspace bool // Limit list to workspace context
}

type ShapesIndexItemResult struct {
	Workspace string
	Frame     string
	Shaper    string
	Name      string
	About     string
	IsContext bool
}

func ShapesIndex(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	var ss []*models.Shape

	params := &ShapesIndexParams{}
	utils.JsonUnmarshal(jparams, params)

	uc := models.UserContextNew()
	if params.Workspace {
		uc.Load("Workspace.Frames.Shapes.Frame.Workspace")
		if uc.Workspace == nil {
			err = app.Error(nil, "no workspace context")
			return
		}
		w := uc.Workspace
		ss = w.Shapes()
	} else {
		uc.Load("Workspaces.Frames.Shapes.Frame.Workspace")
		ss = uc.Shapes()
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

	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
