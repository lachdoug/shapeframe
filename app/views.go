package app

import (
	"embed"
)

var Views *embed.FS

func SetViews(views *embed.FS) {
	Views = views
}
