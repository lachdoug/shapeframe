package controllers

import (
	"sf/app"
	"sf/models"
)

type ShapesCreateParams struct {
	Workspace string
	Frame     string
	Shaper    string
	Name      string
	About     string
}

type ShapesCreateResult struct {
	Workspace string
	Frame     string
	Shape     string
}

func ShapesCreate(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	var f *models.Frame
	var s *models.Shape
	params := paramsFor[ShapesCreateParams](jparams)

	vn = &app.Validation{}
	if params.Workspace == "" {
		vn.Add("Workspace", "must not be blank")
	}
	if params.Frame == "" {
		vn.Add("Frame", "must not be blank")
	}
	if params.Shaper == "" {
		vn.Add("Shaper", "must not be blank")
	}
	if vn.IsInvalid() {
		return
	}

	uc := models.ResolveUserContext(
		"Workspace.Frames.Shapes",
		"Workspaces.Frames.Shapes",
	)
	if w, err = models.ResolveWorkspace(uc, params.Workspace, "Shapers"); err != nil {
		return
	}
	if f, err = models.ResolveFrame(uc, w, params.Frame); err != nil {
		return
	}
	if s, err = models.CreateShape(f, params.Shaper, params.Name, params.About); err != nil {
		return
	}

	result := &ShapesCreateResult{
		Workspace: w.Name,
		Frame:     f.Name,
		Shape:     s.Name,
	}

	jbody = jbodyFor(result)
	return
}
