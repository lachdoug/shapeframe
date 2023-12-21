package main

import (
	"embed"
	"os"
	"sf/app/io"
	"sf/cli"
	"sf/views"
)

// Embed the view templates.
//
//go:embed views
var embededViews embed.FS

func init() {
	io.SetIO(os.Stdout, os.Stdin, os.Stderr)
	views.SetViews(&embededViews)
}

func main() {
	cli.Run(os.Args)
}
