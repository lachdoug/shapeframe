package main

import (
	"embed"
	"os"
	"sf/app"
	"sf/cli"
	"sf/database"
	"sf/database/migration"
	"strings"
)

// Embed the view templates.
//
//go:embed views
var views embed.FS

func init() {
	app.SetDebug(os.Args)
	app.SetIO(os.Stdout, os.Stdin, os.Stderr)
	app.InitViews(&views)
	app.InitDataDir()
	database.Connect()
	migration.Migrate()
}

func main() {
	if arg(0) == "gui" {
		port := arg(1)
		app.Println("Start the GUI", port)
	} else {
		cli.Router(cliArgs())
	}
}

func arg(i int) (val string) {
	if app.Debug {
		i = i + 1
	}
	if i >= len(os.Args) {
		val = ""
	} else {
		val = os.Args[i+1]
	}
	return
}

func cliArgs() (args []string) {
	if app.Debug {
		// Do not show the -debug flag when printing the command
		args = append(os.Args[0:1], os.Args[2:]...)
		app.PrintfErr("Debug: %s\n", strings.Join(args, " "))
	} else {
		args = os.Args
	}
	return
}
