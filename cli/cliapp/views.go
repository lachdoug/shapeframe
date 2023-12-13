package cliapp

import (
	"sf/controllers"
	"sf/views"
)

func View(tplName string, tplPartialNames ...string) (view func(*controllers.Result) (string, error)) {
	view = views.Template("cli", tplName, tplPartialNames)
	return
}
