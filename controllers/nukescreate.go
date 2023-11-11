package controllers

import (
	"sf/utils"
)

func NukesCreate(jparams []byte) (jbody []byte, err error) {
	utils.RemoveDir(utils.DataDir("."))
	return
}
