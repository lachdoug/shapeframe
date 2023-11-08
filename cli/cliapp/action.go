package cliapp

import (
	"encoding/json"
	"sf/app"
)

type Action struct {
	Parametizer func(*Context) ([]byte, *app.Validation, error)
	Controller  func([]byte) ([]byte, *app.Validation, error)
	Viewer      func(map[string]any) (string, error)
}

func (ca *Action) run(context *Context) (err error) {
	var jparams []byte
	var jbody []byte
	var vn *app.Validation
	var output string

	if ca.Parametizer == nil {
		jparams = []byte("{}")
	} else if jparams, vn, err = ca.Parametizer(context); err != nil {
		return
	} else if vn != nil && vn.IsInvalid() {
		err = app.ErrorWith(vn, "invalid")
		return
	}

	if jbody, vn, err = ca.Controller(
		jparams,
	); err != nil {
		return
	} else if vn != nil && vn.IsInvalid() {
		err = app.ErrorWith(vn, "invalid")
		return
	}

	body := &map[string]any{}
	if len(jbody) > 0 {
		if err = json.Unmarshal(jbody, body); err != nil {
			panic(err)
		}
	}

	if output, err = ca.Viewer(*body); err != nil {
		return
	}
	app.Printf(output)
	return
}
