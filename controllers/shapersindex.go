package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type ShapersIndexParams struct {
	Workspace bool // Limit list to workspace context
}

type ShapersIndexItemResult struct {
	Workspace string
	URI       string
	About     string
}

func ShapersIndex(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	var srs []*models.Shaper

	params := &ShapersIndexParams{}
	utils.JsonUnmarshal(jparams, params)

	uc := models.UserContextNew()
	uc.Load("Workspaces")
	if params.Workspace {
		uc.Load("Workspace.Directories.Workspace")
		if uc.Workspace == nil {
			err = app.Error(nil, "no workspace context")
			return
		}
		w := uc.Workspace
		if srs, err = w.Shapers(); err != nil {
			return
		}
	} else {
		uc.Load("Workspaces.Directories.Workspace")
		if srs, err = uc.Shapers(); err != nil {
			return
		}
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

	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
