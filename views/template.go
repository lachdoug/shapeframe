package views

import (
	"bytes"
	htemplate "html/template"
	"path/filepath"
	"sf/controllers"
	"sf/utils"
	"strings"
	ttemplate "text/template"
)

func Template(routerName string, tplName string, tplPartialNames []string) (view func(*controllers.Result) (string, error)) {
	view = func(result *controllers.Result) (output string, err error) {
		var bb bytes.Buffer
		if routerName == "gui" {
			htmlTemplate(&bb, routerName, tplName, tplPartialNames, result)
		} else {
			textTemplate(&bb, routerName, tplName, tplPartialNames, result)
		}
		output = bb.String()
		return
	}
	return
}

func htmlTemplate(
	bb *bytes.Buffer,
	routerName string,
	tplName string,
	tplPartialNames []string,
	result any,
) {
	t := htemplate.New(tplName).Option("missingkey=error")
	tt := htemplate.Must(
		t.Funcs(hFuncMap(t)).
			ParseFS(
				Views,
				viewFilePaths(routerName, tplName, tplPartialNames)...,
			),
	)
	if err := tt.ExecuteTemplate(bb, tplName, result); err != nil {
		panic(err)
	}
}

func textTemplate(
	bb *bytes.Buffer,
	routerName string,
	tplName string,
	tplPartialNames []string,
	result any,
) {
	t := ttemplate.New(tplName).Option("missingkey=error")
	tt := ttemplate.Must(
		t.Funcs(tFuncMap(t)).
			ParseFS(
				Views,
				viewFilePaths(routerName, tplName, tplPartialNames)...,
			),
	)
	if err := tt.ExecuteTemplate(bb, tplName, result); err != nil {
		panic(err)
	}
}

func viewFilePaths(routerName string, tplName string, tplPartialNames []string) (filePaths []string) {
	filePaths = []string{}
	filePaths = append(filePaths, viewFilePath(routerName, tplName))
	for _, tplPartialName := range tplPartialNames {
		filePaths = append(filePaths, viewFilePath(routerName, tplPartialName))
	}
	return
}

func viewFilePath(routerName string, tplName string) (filePath string) {
	elem := append([]string{"views", routerName}, strings.Split(tplName, "/")...)
	filePath = filepath.Join(elem...) + ".go.tpl"
	return
}

func hFuncMap(t *htemplate.Template) (funcMap htemplate.FuncMap) {
	funcMap = htemplate.FuncMap{
		"stringContains": strings.Contains,
		"indentLines":    utils.IndentLines,
		"include": func(name string, data any, indents ...int) (out string, err error) {
			bb := &bytes.Buffer{}
			if err = t.ExecuteTemplate(bb, name, data); err != nil {
				return
			}
			out = utils.IndentLines(strings.TrimRight(bb.String(), " \n"), indents...)
			return
		},
	}
	return
}

func tFuncMap(t *ttemplate.Template) (funcMap ttemplate.FuncMap) {
	funcMap = ttemplate.FuncMap{
		"stringContains": strings.Contains,
		"indentLines":    utils.IndentLines,
		"include": func(name string, data any, indents ...int) (out string, err error) {
			bb := &bytes.Buffer{}
			if err = t.ExecuteTemplate(bb, name, data); err != nil {
				return
			}
			out = utils.IndentLines(strings.TrimRight(bb.String(), " \n"), indents...)
			return
		},
	}
	return
}
