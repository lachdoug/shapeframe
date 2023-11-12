package controllers

import (
	"sf/models"
)

type WorkspacesIndexItemResult struct {
	Name      string
	About     string
	IsContext bool
}

func WorkspacesIndex(jparams []byte) (jbody []byte, err error) {
	uc := models.ResolveUserContext("Workspaces")

	result := []*WorkspacesIndexItemResult{}
	for _, w := range uc.Workspaces {
		result = append(result, &WorkspacesIndexItemResult{
			Name:      w.Name,
			About:     w.About,
			IsContext: uc.WorkspaceID == w.ID,
		})
	}

	jbody = jbodyFor(result)
	return
}
