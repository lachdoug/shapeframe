package controllers

import (
	"sf/models"
)

type FramersIndexParams struct {
	Workspace string
}

type FramersIndexItemResult struct {
	Workspace string
	URI       string
	About     string
}

func FramersIndex(params *Params) (result *Result, err error) {
	var frs []*models.Framer
	p := params.Payload.(*FramersIndexParams)

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace",
	)
	if p.Workspace == "" {
		for _, w := range uc.Workspaces {
			if err = w.Load("Framers"); err != nil {
				return
			}
			frs = append(frs, w.Framers...)
		}
	} else {
		var w *models.Workspace
		if w, err = models.ResolveWorkspace(uc, p.Workspace,
			"Framers",
		); err != nil {
			return
		}
		frs = w.Framers
	}

	r := []*FramersIndexItemResult{}
	var uri string
	for _, fr := range frs {
		if uri, err = fr.URI(); err != nil {
			return
		}
		r = append(r, &FramersIndexItemResult{
			Workspace: fr.Workspace.Name,
			URI:       uri,
			About:     fr.About,
		})
	}

	result = &Result{Payload: r}
	return
}
