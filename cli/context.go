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
		Summary: "Report context",
		Aliases: ss("cx"),
		Usage: ss(
			"sf context [options] [arguments]",
			"Provide a context path as an argument to change context",
			"  A context path uses / as name seperators",
			"  .. will exit a level of context",
			"  A name will enter a level of context",
			"  For example, if context is",
			"    Workspace: Workspace-A",
			"	 Frame: Frame-AB",
			"  sf ../Workspace-B/Frame-BB or sf /Workspace-B/Frame-BB",
			"  Will change context to",
			"    Workspace: Workspace-B",
			"	 Frame: Frame-BB",
		),
		Parametizer: contextParams,
		Controller:  contextController,
		Viewer:      contextViewer,
	}
	return
}

func contextParams(context *cliapp.Context) (jparams []byte, err error) {
	var abs bool
	path := context.Argument(0)
	params := map[string]any{}

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
		contexts := []string{}

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
			params["Workspace"] = contexts[0]
		}
		if len(contexts) > 1 {
			params["Frame"] = contexts[1]
		}
		if len(contexts) > 2 {
			params["Shape"] = contexts[2]
		}

	}

	jparams = jsonParams(params)
	return
}

func contextController(jparams []byte) (jbody []byte, err error) {
	if isContextReadAction {
		return controllers.ContextsRead(jparams)
	}
	return controllers.ContextsUpdate(jparams)
}

func contextViewer(body map[string]any) (output string, err error) {
	if isContextReadAction {
		return cliapp.View(
			"contexts/read",
			"contexts/context",
		)(body)
	}
	return cliapp.View(
		"contexts/update",
		"contexts/context",
	)(body)
}
