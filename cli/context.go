package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
	"strings"
)

var isContextReadAction bool

func context() (command any) {
	command = &cliapp.Command{
		Name:    "context",
		Summary: "Report or change workspace CLI context",
		Aliases: ss("cx"),
		Usage: ss(
			"sf context [options] PATH",
			"Provide a context path as an argument to change context",
			"  A context path uses / as name seperators",
			"  .. will exit a level of context",
			"  A name will enter a level of context",
			"  For example, if context is",
			"    Frame: Frame-A",
			"    Shape: Shape-AA",
			"  sf ../Frame-B/Shape-BA or sf /Frame-B/Shape-BA",
			"  Will change context to",
			"    Frame: Frame-B",
			"    Shape: Shape-BA",
		),
		Handler:    contextHandler,
		Controller: contextController,
		Viewer:     contextViewer,
	}
	return
}

func contextHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	var abs bool
	var w *models.Workspace
	var frame, shape string
	contexts := []string{}
	path := context.Argument(0)

	if len(path) == 0 {
		// If no path is given, read the context
		isContextReadAction = true
	} else {
		// Otherwise, update the context using the path
		args := strings.Split(path, "/")
		if string(path[0]) == "/" {
			abs = true
		}

		if w, err = models.ResolveWorkspace(
			"Frame", "Shape",
		); err != nil {
			return
		}

		if abs {
			args = args[1:]
		} else {
			if w.Frame != nil && w.Shape == nil {
				contexts = []string{w.Frame.Name}
			} else if w.Shape != nil {
				contexts = []string{w.Frame.Name, w.Shape.Name}
			}
		}
		for _, arg := range args {
			if arg == ".." {
				contexts = contexts[:len(contexts)-1]
			} else {
				contexts = append(contexts, arg)
			}
		}
		if len(contexts) > 0 {
			frame = contexts[0]
		}
		if len(contexts) > 1 {
			shape = contexts[1]
		}

		params = &controllers.Params{
			Payload: &controllers.ContextsUpdateParams{
				Frame: frame,
				Shape: shape,
			},
		}
	}
	return
}

func contextController(params *controllers.Params) (result *controllers.Result, err error) {
	if isContextReadAction {
		result, err = controllers.ContextsRead(params)
		return
	}
	result, err = controllers.ContextsUpdate(params)
	return
}

func contextViewer(result *controllers.Result) (output string, err error) {
	if isContextReadAction {
		return cliapp.View(
			"contexts/read",
			"contexts/context",
		)(result)
	}
	return cliapp.View(
		"contexts/update",
		"contexts/context",
	)(result)
}
