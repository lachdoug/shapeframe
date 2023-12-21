package controllers

import (
	"sf/models"
)

type FramersIndexItemResult struct {
	Workspace string
	Framer    string
	URI       string
	About     string
}

func FramersIndex(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var frs []*models.Framer

	if w, err = models.ResolveWorkspace(
		"Framers",
	); err != nil {
		return
	}
	frs = w.Framers

	r := []*FramersIndexItemResult{}
	var uri string
	for _, fr := range frs {
		if uri, err = fr.URI(); err != nil {
			return
		}
		r = append(r, &FramersIndexItemResult{
			Workspace: fr.Workspace.Name,
			Framer:    fr.Name,
			URI:       uri,
			About:     fr.About,
		})
	}

	result = &Result{Payload: r}
	return
}
