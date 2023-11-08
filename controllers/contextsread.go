package controllers

import (
	"sf/app"
	"sf/models"
)

func ContextsRead(jparams []byte) (jbody []byte, vn *app.Validation, err error) {
	uc := models.ResolveUserContext("Workspace", "Frame", "Shape")
	result := uc.Inspect()
	jbody = jbodyFor(result)
	return
}
