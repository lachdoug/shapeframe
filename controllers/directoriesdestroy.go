package controllers

import (
	"path/filepath"
	"sf/app"
	"sf/models"
	"sf/utils"
)

type DirectoriesDestroyParams struct {
	Workspace string
	Path      string
}

type DirectoriesDestroyResult struct {
	Workspace string
	Path      string
}

func DirectoriesDestroy(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	params := &DirectoriesDestroyParams{}
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
	d := w.DirectoryFind(path)
	if d == nil {
		err = app.Error(nil, "directory %s does not exist in workspace %s", path, w.Name)
		return
	}

	d.Destroy()

	result := &DirectoriesDestroyResult{
		Path:      path,
		Workspace: w.Name,
	}
	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
