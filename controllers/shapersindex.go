package controllers

import (
	"sf/models"
)

type ShapersIndexParams struct {
	Workspace string // Limit list to workspace
}

type ShapersIndexItemResult struct {
	Workspace string
	URI       string
	About     string
}

func ShapersIndex(params *Params) (result *Result, err error) {
	var srs []*models.Shaper
	p := params.Payload.(*ShapersIndexParams)

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace",
	)
	if p.Workspace == "" {
		for _, w := range uc.Workspaces {
			if err = w.Load("Shapers"); err != nil {
				return
			}
			srs = append(srs, w.Shapers...)
		}
	} else {
		var w *models.Workspace
		if w, err = models.ResolveWorkspace(uc, p.Workspace,
			"Shapers",
		); err != nil {
			return
		}
		srs = w.Shapers
	}

	r := []*ShapersIndexItemResult{}
	var uri string
	for _, sr := range srs {
		if uri, err = sr.URI(); err != nil {
			return
		}
		r = append(r, &ShapersIndexItemResult{
			Workspace: sr.Workspace.Name,
			URI:       uri,
			About:     sr.About,
		})
	}

	result = &Result{Payload: r}
	return
}
