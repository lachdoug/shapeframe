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

func listWorkspacesViewer(result *controllers.Result) (output string, err error) {
	rs := result.Payload.([]*controllers.WorkspacesIndexItemResult)
	items := []map[string]any{}
	for _, r := range rs {
		items = append(items, map[string]any{
			"Workspace": r.Workspace,
			"About":     r.About,
			"IsContext": r.IsContext,
		})
	}

	table := &Table{
		Items:  items,
		Titles: ss("WORKSPACE", "ABOUT"),
		Keys:   ss("Workspace", "About"),
		// Values: tvs(
		// 	tableCellStringValueFn,
		// 	tableCellStringValueFn,
		// ),
		Accents: tas(
			tableCellGreenIfInContextAccentFn,
			tableCellNoAccentFn,
		),
	}

	result = &controllers.Result{
		Payload: table.generate(),
	}
	output, err = cliapp.View("workspaces/index")(result)
	return
}
