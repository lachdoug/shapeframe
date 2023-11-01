package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
)

func listWorkspaces() (command any) {
	command = &cliapp.Command{
		Name:        "workspaces",
		Summary:     "List workspaces",
		Aliases:     ss("w"),
		Flags:       ss("bool", "w", "Limit list to workspace"),
		Parametizer: listWorkspacesParams,
		Controller:  controllers.WorkspacesIndex,
		Viewer:      listWorkspacesViewer,
	}
	return
}

func listWorkspacesParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	jparams = jsonParams(map[string]any{
		"Workspace": context.BoolFlag("w"),
	})
	return
}

func listWorkspacesViewer(body map[string]any) (output string, er error) {
	table := &Table{
		Items:  resultItems(body),
		Titles: ss(" ", "WORKSPACE", "ABOUT"),
		Keys:   ss("IsContext", "Name", "About"),
		Values: tvs(
			tableCellAsteriskIfTrueValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
		),
		Accents: tas(
			tableCellNoAccentFn,
			tableCellGreenIfInContextAccentFn,
			tableCellNoAccentFn,
		),
	}
	return cliapp.View("workspaces/index")(table.generate())
}
