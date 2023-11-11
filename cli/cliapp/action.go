package cliapp

import (
	"encoding/json"
	"fmt"
	"sf/app"
)

type Action struct {
	Parametizer func(*Context) ([]byte, error)
	Controller  func([]byte) ([]byte, error)
	Viewer      func(map[string]any) (string, error)
}

func (ca *Action) run(context *Context) (err error) {
	var jparams []byte
	var jbody []byte
	var output string

	// Get the command params from the parametizer
	if ca.Parametizer == nil {
		jparams = []byte("{}")
	} else if jparams, err = ca.Parametizer(context); err != nil {
		return
	}

	// Call the controller with the command params and get json
	if jbody, err = ca.Controller(
		jparams,
	); err != nil {
		return
	}

	// Unmarshal the controller json
	body := map[string]any{}
	if len(jbody) > 0 {
		if err = json.Unmarshal(jbody, &body); err != nil {
			panic(err)
		}
	}

	// Check for invalidity and if invalid return with formatted error output
	if err = validation(body); err != nil {
		return
	}

	// Otherwise call the viewer
	if output, err = ca.Viewer(body); err != nil {
		return
	}

	// And print the output
	app.Print(output)
	return
}

func validation(body map[string]any) (err error) {
	if body["Invalid"] != nil {
		s := ""
		vmap := body["Invalid"].(map[string]any)
		for k, v := range vmap {
			s = s + fmt.Sprintf("\n  - %s %s", k, v)
		}
		err = app.Error("invalid:" + s)
	}
	return
}
