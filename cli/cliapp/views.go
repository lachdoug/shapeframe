package cliapp

import (
	"bytes"
	"path/filepath"
	"sf/app"
	"sf/utils"
	"strings"
	"text/template"
)

func View(tplName string, tplPartialNames ...string) (view func(map[string]any) (string, error)) {
	view = templater("cli", tplName, tplPartialNames)
	return
}

func viewFilePath(routerName string, tplName string) (filePath string) {
	elem := append([]string{"views", routerName}, strings.Split(tplName, "/")...)
	filePath = filepath.Join(elem...) + ".go.tpl"
	return
}

func viewFilePaths(routerName string, tplName string, tplPartialNames []string) (filePaths []string) {
	filePaths = []string{}
	filePaths = append(filePaths, viewFilePath(routerName, tplName))
	for _, tplPartialName := range tplPartialNames {
		filePaths = append(filePaths, viewFilePath(routerName, tplPartialName))
	}
	return
}

func funcMap(t *template.Template) (funcMap template.FuncMap) {
	funcMap = template.FuncMap{
		"include": func(name string, data any, indents ...int) (out string, err error) {
			bb := &bytes.Buffer{}
			if err = t.ExecuteTemplate(bb, name, data); err != nil {
				return
			}
			out = utils.IndentLines(strings.TrimSpace(bb.String()), indents...)
			return
		},
	}
	return
}

func templater(routerName string, tplName string, tplPartialNames []string) (view func(map[string]any) (string, error)) {
	view = func(body map[string]any) (output string, err error) {
		var bb bytes.Buffer
		t := template.New(tplName).Option("missingkey=error")
		tt := template.Must(
			t.Funcs(funcMap(t)).
				ParseFS(app.Views, viewFilePaths(routerName, tplName, tplPartialNames)...),
		)
		if err = tt.ExecuteTemplate(&bb, tplName, body); err != nil {
			panic(err)
		}
		output = bb.String()
		return
	}
	return
}
