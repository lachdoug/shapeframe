package cliapp

import (
	"sf/app/errors"
	"sf/app/io"
	"sf/controllers"
)

type Action struct {
	Handler    func(*Context) (*controllers.Params, error)
	Controller func(*controllers.Params) (*controllers.Result, error)
	Viewer     func(*controllers.Result) (string, error)
}

func (ca *Action) run(context *Context) (err error) {
	var params *controllers.Params
	var result *controllers.Result
	var output string

	// Get params from the command handler
	if ca.Handler != nil {
		if params, err = ca.Handler(context); err != nil {
			return
		}
	}

	// Call the controller with params
	if result, err = ca.Controller(params); err != nil {
		return
	}

	// Check for invalidity and if invalid return with formatted error output
	if result != nil && result.Validation != nil && result.Validation.IsInvalid() {
		err = errors.ValidationError(result.Validation.Maps())
		return
	}

	// Otherwise call the viewer
	if output, err = ca.Viewer(result); err != nil {
		return
	}

	// And print the output
	io.Print(output)
	return
}
