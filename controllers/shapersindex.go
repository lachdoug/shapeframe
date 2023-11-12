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

func ShapersIndex(jparams []byte) (jbody []byte, err error) {
	var srs []*models.Shaper
	params := ParamsFor[ShapersIndexParams](jparams)

	uc := models.ResolveUserContext(
		"Workspaces", "Workspace",
	)
	if params.Workspace == "" {
		for _, w := range uc.Workspaces {
			if err = w.Load("Shapers"); err != nil {
				return
			}
			srs = append(srs, w.Shapers...)
		}
	} else {
		var w *models.Workspace
		if w, err = models.ResolveWorkspace(uc, params.Workspace,
			"Shapers",
		); err != nil {
			return
		}
		srs = w.Shapers
	}

	result := []*ShapersIndexItemResult{}
	var uri string
	for _, sr := range srs {
		if uri, err = sr.URI(); err != nil {
			return
		}
		result = append(result, &ShapersIndexItemResult{
			Workspace: sr.Workspace.Name,
			URI:       uri,
			About:     sr.About,
		})
	}

	jbody = jbodyFor(result)
	return
}
