package app

import (
	"sf/utils"
)

func SetDirs() {
	utils.MakeDir(utils.DataDir("."))
	utils.MakeDir(utils.TempDir("."))
	utils.MakeDir(utils.LogDir("."))
}
