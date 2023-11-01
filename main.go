package main

import (
	"embed"
	"fmt"
	"os"
	"sf/app"
	"sf/cli"
	"sf/database"
	"sf/database/migration"
)

// Embed the view templates.
//
//go:embed views
var views embed.FS

func init() {
	app.InitViews(&views)
	app.InitDataDir()
	database.Connect()
	migration.Migrate()
}

func main() {
	if arg(1) == "gui" {
		port := arg(2)
		fmt.Println("Start the GUI", port)
	} else {
		cli.Router(os.Args)
	}
}

func arg(i int) (val string) {
	if i >= len(os.Args) {
		val = ""
	} else {
		val = os.Args[i]
	}
	return
}
