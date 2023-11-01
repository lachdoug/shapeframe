package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type WorkspacesUpdateParams struct {
	Workspace string
	Name      string
	About     string
}

type WorkspacesUpdateResult struct {
	From *WorkspacesUpdateResultDetails
	To   *WorkspacesUpdateResultDetails
}

type WorkspacesUpdateResultDetails struct {
	Name  string
	About string
}

func WorkspacesUpdate(jparams []byte) (jbody []byte, validation *app.Validation, err error) {
	params := &WorkspacesUpdateParams{}
	utils.JsonUnmarshal(jparams, params)

	uc := models.UserContextNew()
	uc.Load("Workspaces")
	w := uc.WorkspaceFind(params.Workspace)
	if w == nil {
		err = app.Error(nil, "workspace %s does not exist", params.Workspace)
		return
	}

	result := &WorkspacesUpdateResult{
		From: &WorkspacesUpdateResultDetails{
			Name:  w.Name,
			About: w.About,
		},
	}

	w.Assign(map[string]any{
		"Name":  params.Name,
		"About": params.About,
	})
	if err = w.Save(); err != nil {
		return
	}

	result.To = &WorkspacesUpdateResultDetails{
		Name:  w.Name,
		About: w.About,
	}

	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
