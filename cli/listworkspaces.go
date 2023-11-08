package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func listWorkspaces() (command any) {
	command = &cliapp.Command{
		Name:    "workspaces",
		Summary: "List workspaces",
		Aliases: ss("w"),
		Usage: ss(
			"sf list workspaces [options]",
		),
		Controller: controllers.WorkspacesIndex,
		Viewer:     listWorkspacesViewer,
	}
	return
}

func listWorkspacesViewer(body map[string]any) (output string, er error) {
	table := &Table{
		Items:  resultItems(body),
		Titles: ss("WORKSPACE", "ABOUT"),
		Keys:   ss("Name", "About"),
		Values: tvs(
			tableCellStringValueFn,
			tableCellStringValueFn,
		),
		Accents: tas(
			tableCellGreenIfInContextAccentFn,
			tableCellNoAccentFn,
		),
	}
	return cliapp.View("workspaces/index")(table.generate())
}
