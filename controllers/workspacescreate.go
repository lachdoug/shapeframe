package controllers

import (
	"fmt"
	"sf/app"
	"sf/models"
	"sf/utils"
)

type WorkspacesCreateParams struct {
	Name  string
	About string
}

type WorkspacesCreateResult struct {
	Name string
}

func WorkspacesCreate(jparams []byte) (jbody []byte, v *app.Validation, err error) {
	params := &WorkspacesCreateParams{}
	utils.JsonUnmarshal(jparams, params)

	name := utils.StringTidy(params.Name)

	v = &app.Validation{}
	if name == "" {
		v.Add("Name", "must not be blank")
	}
	if v.IsInvalid() {
		return
	}

	uc := models.UserContextNew()
	fmt.Println("User Context", uc)
	w := models.WorkspaceNew(uc, name)
	w.Assign(map[string]any{
		"About": params.About,
	})
	if err = w.Create(); err != nil {
		return
	}

	result := &WorkspacesCreateResult{Name: w.Name}
	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
