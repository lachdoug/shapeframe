package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
)

func bye() (commandset any) {
	commandset = &cliapp.CommandSet{
		Name:        "bye",
		Summary:     "Say goodbye",
		Description: ss("See ya"),
		Aliases:     ss("b"),
		Flags:       ss("bool", "w", "WTF"),
		Commands:    cs(byeHappy, byeSad),
	}
	return
}

func byeHappy() (command any) {
	command = &cliapp.Command{
		Name:        "happy",
		Summary:     "Say a happy goodbye",
		Description: ss("Cheerio"),
		Aliases:     ss("h"),
		Flags:       ss(),
		Parametizer: byeHappyParams,
		Controller:  controllers.ByesRead,
		Viewer:      cliapp.View("byes/read"),
	}
	return
}

func byeHappyParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	jparams = jsonParams(map[string]any{
		"Tone":  "happy",
		"Name":  context.Argument(0),
		"Throw": context.BoolFlag("throw"),
	})
	return
}

func byeSad() (command any) {
	command = &cliapp.Command{
		Name:        "sad",
		Summary:     "say a sad goodbye",
		Description: ss("Please don't go"),
		Aliases:     ss("s"),
		Flags:       ss(),
		Parametizer: byeSadParams,
		Controller:  controllers.ByesRead,
		Viewer:      cliapp.View("byes/read"),
	}
	return
}

func byeSadParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	jparams = jsonParams(map[string]any{
		"Tone":  "sad",
		"Name":  context.Argument(0),
		"Throw": context.BoolFlag("throw"),
	})
	return
}
