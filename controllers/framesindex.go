package controllers

import (
	"sf/models"
)

type FramesIndexParams struct {
	Workspace string
}

type FramesIndexItemResult struct {
	Workspace string
	Frame     string
	Framer    string
	About     string
	IsContext bool
}

func FramesIndex(params *Params) (result *Result, err error) {
	if params.Payload == nil {
		params.Payload = &FramesIndexParams{}
	}
	p := params.Payload.(*FramesIndexParams)
	var fs []*models.Frame

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace",
	)
	if p.Workspace == "" {
		for _, w := range uc.Workspaces {
			w.Load("Frames.Workspace")
			fs = append(fs, w.Frames...)
		}
	} else {
		var w *models.Workspace
		if w, err = models.ResolveWorkspace(uc, p.Workspace,
			"Frames.Workspace",
		); err != nil {
			return
		}
		fs = w.Frames
	}

	r := []*FramesIndexItemResult{}
	for _, f := range fs {
		r = append(r, &FramesIndexItemResult{
			Workspace: f.Workspace.Name,
			Frame:     f.Name,
			Framer:    f.FramerName,
			About:     f.About,
			IsContext: uc.FrameID == f.ID,
		})
	}

	result = &Result{Payload: r}
	return
}
