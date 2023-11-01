package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
)

func hi() (command any) {
	command = &cliapp.Command{
		Name:        "hi",
		Summary:     "Say hello",
		Description: ss("Yo"),
		Aliases:     ss("h"),
		Flags: ss(
			"bool", "x", "Extra exclamations",
			"bool", "throw", "Throw an error",
		),
		Parametizer: hiParams,
		Controller:  controllers.HisRead,
		Viewer:      cliapp.View("his/read"),
	}
	return
}

func hiParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	jparams = jsonParams(map[string]any{
		"Name":  context.Argument(0),
		"Extra": context.BoolFlag("x"),
		"Throw": context.BoolFlag("throw"),
	})
	return
}
