package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

type ByesReadParams struct {
	Tone  string
	Name  string
	Throw bool
}

type ByesReadResult struct {
	Tone string
	Name string
}

func (params *ByesReadParams) validation() (validation *app.Validation) {
	validation = &app.Validation{}
	if params.Name == "" {
		validation.Add("Name", "must not be blank")
	}
	return
}

func ByesRead(jparams []byte) (jbody []byte, validation *app.Validation, err error) {
	params := &ByesReadParams{}
	utils.JsonUnmarshal(jparams, params)
	if validation = params.validation(); validation.IsInvalid() {
		return
	}
	bye := models.NewBye(params.Name, params.Tone)

	result := &ByesReadResult{
		Tone: bye.Tone,
		Name: bye.Name,
	}
	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
