package controllers

import (
	"sf/models"
)

type ShapersIndexItemResult struct {
	Workspace string
	Shaper    string
	URI       string
	About     string
}

func ShapersIndex(params *Params) (result *Result, err error) {
	var w *models.Workspace
	var srs []*models.Shaper

	if w, err = models.ResolveWorkspace(
		"Shapers",
	); err != nil {
		return
	}
	srs = w.Shapers

	r := []*ShapersIndexItemResult{}
	var uri string
	for _, sr := range srs {
		if uri, err = sr.URI(); err != nil {
			return
		}
		r = append(r, &ShapersIndexItemResult{
			Workspace: sr.Workspace.Name,
			Shaper:    sr.Name,
			URI:       uri,
			About:     sr.About,
		})
	}

	result = &Result{Payload: r}
	return
}
