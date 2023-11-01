package cliapp

import (
	"bytes"
	"log"
	"path/filepath"
	"sf/app"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func View(n string) (view func(map[string]any) (string, error)) {
	view = templater("cli", n)
	return
}

func templater(r string, n string) (view func(map[string]any) (string, error)) {
	view = func(body map[string]any) (output string, err error) {
		elem := append([]string{"views", r}, strings.Split(n, "/")...)
		tp := filepath.Join(elem...) + ".go.tpl"

		tpl, err := app.Views.ReadFile(tp)
		if err != nil {
			log.Fatal(err)
		}

		var bb bytes.Buffer
		tt := template.Must(template.New(n).Funcs(sprig.FuncMap()).Parse(string(tpl)))
		if err = tt.Execute(&bb, body); err != nil {
			log.Fatal(err)
		}
		output = bb.String()
		return
	}
	return
}
