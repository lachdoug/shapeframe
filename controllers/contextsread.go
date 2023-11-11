package controllers

import (
	"sf/models"
)

func ContextsRead(jparams []byte) (jbody []byte, err error) {
	uc := models.ResolveUserContext("Workspace", "Frame", "Shape")
	result := uc.Inspect()
	jbody = jbodyFor(result)
	return
}
