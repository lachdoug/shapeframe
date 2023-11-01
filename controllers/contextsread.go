package controllers

import (
	"sf/app"
	"sf/models"
	"sf/utils"
)

func ContextsRead(jparams []byte) (jbody []byte, validation *app.Validation, err error) {
	uc := models.UserContextNew()
	uc.Load("Workspace", "Frame", "Shape")
	result := uc.Inspect()

	body := &app.Body{Result: result}
	jbody = utils.JsonMarshal(body)
	return
}
