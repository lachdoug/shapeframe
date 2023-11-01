package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type HisReadParams struct {
	Name  string
	Extra bool
	Throw bool
}

type HisReadResult struct {
	Name  string
	Extra bool
}

func (params *HisReadParams) validation() (validation *app.Validation) {
	validation = &app.Validation{}
	if params.Name == "" {
		validation.Add("Name", "must not be blank")
	}
	return
}

func HisRead(jparams []byte) (jbody []byte, validation *app.Validation, err error) {
	params := &HisReadParams{}
	utils.JsonUnmarshal(jparams, params)
	if validation = params.validation(); validation.IsInvalid() {
		return
	}
	if params.Throw {
		err = app.Error(nil, "WTF %s", params.Name)
		return
	}
	hi := &models.Hi{
		Name: params.Name,
	}
	result := &HisReadResult{
		Name:  hi.Name,
		Extra: params.Extra,
	}
	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
