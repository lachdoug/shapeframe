package guiapp

import (
	"bytes"
	"html/template"
	"log"
	"path/filepath"
	"sf/app"
	"strings"
)

func View(n string) (view func(map[string]any) (string, error)) {
	view = templater("gui", n)
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
		ht := template.Must(template.New(n).Parse(string(tpl)))
		if err = ht.Execute(&bb, body); err != nil {
			log.Fatal(err)
		}
		return
	}
	return
}
