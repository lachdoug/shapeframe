package controllers

import (
	"sf/models"
)

type FramersIndexParams struct {
	Workspace string // Limit list to workspace
}

type FramersIndexItemResult struct {
	Workspace string
	URI       string
	About     string
}

func FramersIndex(jparams []byte) (jbody []byte, err error) {
	var frs []*models.Framer
	params := ParamsFor[FramersIndexParams](jparams)

	uc := models.ResolveUserContext(
		"Workspace",
		"Workspaces",
	)
	if params.Workspace == "" {
		for _, w := range uc.Workspaces {
			if err = w.Load("Framers"); err != nil {
				return
			}
			frs = append(frs, w.Framers...)
		}
	} else {
		var w *models.Workspace
		if w, err = models.ResolveWorkspace(uc, params.Workspace); err != nil {
			return
		}
		if err = w.Load("Framers"); err != nil {
			return
		}
		frs = w.Framers
	}

	result := []*FramersIndexItemResult{}
	var uri string
	for _, fr := range frs {
		if uri, err = fr.URI(); err != nil {
			return
		}
		result = append(result, &FramersIndexItemResult{
			Workspace: fr.Workspace.Name,
			URI:       uri,
			About:     fr.About,
		})
	}

	jbody = jbodyFor(result)
	return
}
