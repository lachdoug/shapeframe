package cliapp

import (
	"encoding/json"
	"fmt"
	"log"
	"sf/app"
)

type Action struct {
	Parametizer func(*Context) ([]byte, *app.Validation, error)
	Controller  func([]byte) ([]byte, *app.Validation, error)
	Viewer      func(map[string]any) (string, error)
}

func (ca *Action) run(context *Context) {
	var jparams []byte
	var jbody []byte
	var validation *app.Validation
	var output string
	var err error

	if ca.Parametizer == nil {
		jparams = []byte("{}")
	} else if jparams, validation, err = ca.Parametizer(context); err != nil {
		ca.pError(err)
		return
	} else if validation != nil && validation.IsInvalid() {
		ca.pInvalid(validation)
		return
	}

	if jbody, validation, err = ca.Controller(
		jparams,
	); err != nil {
		ca.pError(err)
		return
	} else if validation != nil && validation.IsInvalid() {
		ca.pInvalid(validation)
		return
	}

	body := &map[string]any{}
	json.Unmarshal(jbody, body)

	if output, err = ca.Viewer(*body); err != nil {
		ca.pError(err)
		return
	}
	ca.pOutput(output)
}

func (ca *Action) pError(apperr error) {
	if _, err := fmt.Printf("Error: %s\n", apperr); err != nil {
		log.Fatal(err)
	}
}

func (ca *Action) pInvalid(validation *app.Validation) {
	if _, err := fmt.Printf("Invalid\n%s\n", validation); err != nil {
		log.Fatal(err)
	}
}

func (ca *Action) pOutput(output string) {
	if _, err := fmt.Print(output); err != nil {
		log.Fatal(err)
	}
}
