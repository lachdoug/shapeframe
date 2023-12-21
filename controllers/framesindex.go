package controllers

import (
	"sf/models"
)

type FramesIndexItemResult struct {
	Workspace string
	Frame     string
	Framer    string
	About     string
	IsContext bool
}

func FramesIndex(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var fs []*models.Frame

	if w, err = models.ResolveWorkspace(
		"Frames.Workspace",
	); err != nil {
		return
	}
	fs = w.Frames

	r := []*FramesIndexItemResult{}
	for _, f := range fs {
		r = append(r, &FramesIndexItemResult{
			Workspace: f.Workspace.Name,
			Frame:     f.Name,
			Framer:    f.FramerName,
			About:     f.About,
			IsContext: w.FrameID == f.ID,
		})
	}

	result = &Result{Payload: r}
	return
}
