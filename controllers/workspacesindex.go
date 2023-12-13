package controllers

import (
	"sf/models"
)

type WorkspacesIndexItemResult struct {
	Workspace string
	About     string
	IsContext bool
}

func WorkspacesIndex(params *Params) (result *Result, err error) {
	uc := models.ResolveUserContext("Workspaces")

	r := []*WorkspacesIndexItemResult{}
	for _, w := range uc.Workspaces {
		r = append(r, &WorkspacesIndexItemResult{
			Workspace: w.Name,
			About:     w.About,
			IsContext: uc.WorkspaceID == w.ID,
		})
	}

	result = &Result{Payload: r}
	return
}
