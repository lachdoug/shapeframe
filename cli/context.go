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
		Summary: "Report or change context",
		Aliases: ss("cx"),
		Usage: ss(
			"sf context [options] [arguments]",
			"Provide a context path as an argument to change context",
			"  A context path uses / as name seperators",
			"  .. will exit a level of context",
			"  A name will enter a level of context",
			"  For example, if context is",
			"    Workspace: Workspace-A",
			"    Frame: Frame-AB",
			"  sf ../Workspace-B/Frame-BB or sf /Workspace-B/Frame-BB",
			"  Will change context to",
			"    Workspace: Workspace-B",
			"    Frame: Frame-BB",
		),
		Handler:    contextHandler,
		Controller: contextController,
		Viewer:     contextViewer,
	}
	return
}

func contextHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	var abs bool
	var workspace, frame, shape string
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

		uc := models.ResolveUserContext(
			"Workspace", "Frame", "Shape",
		)

		if abs {
			args = args[1:]
		} else {
			if uc.Workspace != nil && uc.Frame == nil {
				contexts = []string{uc.Workspace.Name}
			} else if uc.Frame != nil && uc.Shape == nil {
				contexts = []string{uc.Workspace.Name, uc.Frame.Name}
			} else if uc.Shape != nil {
				contexts = []string{uc.Workspace.Name, uc.Frame.Name, uc.Shape.Name}
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
			workspace = contexts[0]
		}
		if len(contexts) > 1 {
			frame = contexts[1]
		}
		if len(contexts) > 2 {
			shape = contexts[2]
		}

		params = &controllers.Params{
			Payload: &controllers.ContextsUpdateParams{
				Workspace: workspace,
				Frame:     frame,
				Shape:     shape,
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
