package main

import (
	"embed"
	"fmt"
	"os"
	"regexp"
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
	app.SetViews(&views)
	database.Connect()
	migration.Migrate()
}

func main() {
	if arg(0) == "gui" {
		port := arg(1)
		app.Println("Start the GUI", port)
	} else {
		rargs := cliRouterArgs()
		if app.Debug {
			debugMessage(rargs)
		}
		cli.Router(rargs)
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

func cliRouterArgs() (rargs []string) {
	if app.Debug {
		// Remove the -debug flag fram args before passing to router
		rargs = append(os.Args[0:1], os.Args[2:]...)
	} else {
		rargs = os.Args
	}
	return
}

func debugMessage(rargs []string) {
	redColor := "\033[0;31m"
	noColor := "\033[0m"
	// Quote args with spaces
	argHasSpaceRegexp := regexp.MustCompile(`\s`)
	for i, rarg := range rargs {
		if argHasSpaceRegexp.Match([]byte(rarg)) {
			rargs[i] = fmt.Sprintf("%q", rarg)
		}
	}
	app.PrintfErr("%sDebug: %s%s\n", redColor, strings.Join(rargs, " "), noColor)
}
