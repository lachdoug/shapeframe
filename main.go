package main

import (
	"embed"
	"fmt"
	"os"
	"regexp"
	"sf/app"
	"sf/app/errors"
	"sf/app/io"
	"sf/app/logs"
	"sf/cli"
	"sf/database"
	"sf/database/migration"
	"sf/gui"
	"sf/tui"
	"sf/views"
	"strings"
)

// Embed the view templates.
//
//go:embed views
var embededViews embed.FS

func init() {
	app.SetDirs()
	logs.SetLogger()
	errors.SetDebug(os.Args)
	io.SetIO(os.Stdout, os.Stdin, os.Stderr)
	views.SetViews(&embededViews)
	database.Connect()
	migration.Migrate()
}

func main() {
	logs.Logf("sf command args: %s", os.Args)

	rargs := routerArgs()
	if errors.Debug {
		debugMessage(rargs)
	}
	if arg(0) == "gui" {
		gui.Run(arg(1))
	} else if arg(0) == "tui" {
		tui.Run()
	} else {
		cli.Run(rargs)
	}
}

func arg(i int) (val string) {
	if errors.Debug {
		i = i + 2 // skip first two args, which are sf --debug
	} else {
		i = i + 1 // skip first arg, which is sf
	}
	if i >= len(os.Args) {
		val = ""
	} else {
		val = os.Args[i]
	}
	return
}

func routerArgs() (rargs []string) {
	if errors.Debug {
		// Remove the --debug flag fram args before passing to router
		rargs = append(os.Args[0:1], os.Args[2:]...)
	} else {
		rargs = os.Args
	}
	return
}

func debugMessage(rargs []string) {
	// Quote args with spaces
	argHasSpaceRegexp := regexp.MustCompile(`\s`)
	for i, rarg := range rargs {
		if argHasSpaceRegexp.Match([]byte(rarg)) {
			rargs[i] = fmt.Sprintf("%q", rarg)
		}
	}
	io.PrintfErr("%sDebug: %s%s\n", io.RedColor, strings.Join(rargs, " "), io.ResetText)
}
