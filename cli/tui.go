package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/tui"
)

func terminalUserInterface() (command any) {
	command = &cliapp.Command{
		Name:    "tui",
		Summary: "Run the Terminal User Interface",
		Usage: ss(
			"sf tui",
		),
		Handler: terminalUserInterfaceHandler,
	}
	return
}

func terminalUserInterfaceHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	tui.Run()
	return
}
