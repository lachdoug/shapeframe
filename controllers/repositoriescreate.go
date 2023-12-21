package controllers

import (
	"sf/app/streams"
	"sf/app/validations"
	"sf/models"
)

type RepositoriesCreateParams struct {
	URI      string
	Protocol string
	Username string
	Password string
}

type RepositoriesCreateResult struct {
	Workspace string
	URI       string
	URL       string
}

func RepositoriesCreate(params *Params) (result *Result, err error) {
	p := params.Payload.(*RepositoriesCreateParams)
	var w *models.Workspace
	var r *models.Repository
	var vn *validations.Validation
	st := streams.StreamCreate()

	if w, err = models.ResolveWorkspace(
		"Repositories",
	); err != nil {
		return
	}
	if r, vn, err = models.CreateRepository(
		w,
		p.URI,
		p.Protocol,
		p.Username,
		p.Password,
		st,
	); err != nil {
		return
	}

	result = &Result{
		Payload: &RepositoriesCreateResult{
			Workspace: r.Workspace.Name,
			URI:       p.URI,
			URL:       r.OriginURL(p.Protocol),
		},
		Validation: vn,
		Stream:     st,
	}
	return
}
