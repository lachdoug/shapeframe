package controllers

import (
	"sf/models"
)

type ContextsReadResult struct {
	Workspace string
	Frame     string
	Shape     string
}

func ContextsRead(params *Params) (result *Result, err error) {
	var w *models.Workspace
	if w, err = models.ResolveWorkspace(
		"Frame", "Shape",
	); err != nil {
		return
	}

	result = &Result{
		Payload: &ContextsReadResult{
			Frame: w.FrameName(),
			Shape: w.ShapeName(),
		},
	}
	return
}
