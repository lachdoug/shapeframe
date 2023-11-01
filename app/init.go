package app

import (
	"embed"
	"sf/utils"
)

var Views *embed.FS

func InitViews(views *embed.FS) {
	Views = views
}

func InitDataDir() {
	utils.MakeDir(utils.DataDir("."))
}
