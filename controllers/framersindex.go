package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type FramersIndexParams struct {
	Workspace bool // Limit list to workspace context
}

type FramersIndexItemResult struct {
	Workspace string
	URI       string
	About     string
}

func FramersIndex(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	var frs []*models.Framer

	params := &FramersIndexParams{}
	utils.JsonUnmarshal(jparams, params)

	uc := models.UserContextNew()
	uc.Load("Workspace")
	if params.Workspace {
		uc.Load("Workspace.Directories.Workspace")
		if uc.Workspace == nil {
			err = app.Error(nil, "no workspace context")
			return
		}
		w := uc.Workspace
		if frs, err = w.Framers(); err != nil {
			return
		}
	} else {
		uc.Load("Workspaces.Directories.Workspace")
		if frs, err = uc.Framers(); err != nil {
			return
		}
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

	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
