package controllers

import (
	"path/filepath"
	"sf/app"
	"sf/models"
	"sf/utils"
)

type DirectoriesCreateParams struct {
	Workspace string
	Path      string
}

type DirectoriesCreateResult struct {
	Workspace string
	Directory *models.DirectoryInspector
}

func DirectoriesCreate(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	params := &DirectoriesCreateParams{}
	utils.JsonUnmarshal(jparams, params)

	v = &app.Validation{}
	if params.Workspace == "" {
		v.Add("Workspace", "must not be blank")
	}
	if params.Path == "" {
		v.Add("Path", "must not be blank")
	}
	if v.IsInvalid() {
		return
	}

	uc := models.UserContextNew()
	w := uc.WorkspaceFind(params.Workspace)
	if w == nil {
		err = app.Error(nil, "workspace %s does not exist", params.Workspace)
		return
	}
	var path string
	if path, err = filepath.Abs(params.Path); err != nil {
		return
	}
	d := models.DirectoryNew(w, path)
	if d.IsExists() {
		err = app.Error(nil, "directory %s already exists in workspace %s", path, w.Name)
		return
	}

	d.Create()
	d.Load()

	var di *models.DirectoryInspector
	if di, err = d.Inspect(); err != nil {
		return
	}
	result := &DirectoriesCreateResult{
		Workspace: w.Name,
		Directory: di,
	}
	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
