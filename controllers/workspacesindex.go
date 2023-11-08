package controllers

import (
	"sf/app"
	"sf/models"
)

type WorkspacesIndexItemResult struct {
	Name      string
	About     string
	IsContext bool
}

func WorkspacesIndex(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	var ws []*models.Workspace

	uc := models.ResolveUserContext(
		"Workspace",
		"Workspaces",
	)
	ws = uc.Workspaces

	result := []*WorkspacesIndexItemResult{}
	for _, w := range ws {
		result = append(result, &WorkspacesIndexItemResult{
			Name:      w.Name,
			About:     w.About,
			IsContext: uc.WorkspaceID == w.ID,
		})
	}

	jbody = jbodyFor(result)
	return
}
