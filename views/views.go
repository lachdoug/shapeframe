package views

import (
	"embed"
)

var Views *embed.FS

func SetViews(views *embed.FS) {
	Views = views
}
