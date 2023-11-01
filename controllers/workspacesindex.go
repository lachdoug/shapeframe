package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type WorkspacesIndexParams struct {
	Workspace bool // Limit list to workspace context
}

type WorkspacesIndexItemResult struct {
	Name      string
	About     string
	IsContext bool
}

func WorkspacesIndex(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	var ws []*models.Workspace

	params := &WorkspacesIndexParams{}
	utils.JsonUnmarshal(jparams, params)

	uc := models.UserContextNew()
	if params.Workspace {
		uc.Load("Workspace")
		if uc.Workspace == nil {
			err = app.Error(nil, "no workspace context")
			return
		}
		w := uc.Workspace
		ws = append(ws, w)
	} else {
		uc.Load("Workspaces")
		ws = uc.Workspaces
	}

	result := []*WorkspacesIndexItemResult{}
	for _, w := range ws {
		result = append(result, &WorkspacesIndexItemResult{
			Name:      w.Name,
			About:     w.About,
			IsContext: uc.WorkspaceID == w.ID,
		})
	}

	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
