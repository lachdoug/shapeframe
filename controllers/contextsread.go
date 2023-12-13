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
	uc := models.ResolveUserContext("Workspace", "Frame", "Shape")

	result = &Result{
		Payload: &ContextsReadResult{
			Workspace: uc.WorkspaceName(),
			Frame:     uc.FrameName(),
			Shape:     uc.ShapeName(),
		},
	}
	return
}
