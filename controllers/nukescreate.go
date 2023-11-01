package controllers

import (
	"sf/app"
	"sf/utils"
)

func NukesCreate(jparams []byte) (jbody []byte, validation *app.Validation, err error) {
	utils.RemoveDir(utils.DataDir("."))
	return
}
