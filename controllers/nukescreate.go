package controllers

import (
	"sf/utils"
)

func NukesCreate(params *Params) (result *Result, err error) {
	utils.RemoveDir(utils.DataDir("."))
	utils.RemoveDir(utils.TempDir("."))
	utils.RemoveDir(utils.LogDir("."))
	return
}
