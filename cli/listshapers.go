package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func listShapers() (command any) {
	command = &cliapp.Command{
		Name:    "shapers",
		Summary: "List shapers",
		Aliases: ss("sr"),
		Usage: ss(
			"sf list shapers [options]",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"List shapers in all workspaces by setting the -all flag",
			"  Otherwise lists shapers in workspace context",
			"  Overrides -workspace flag",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
			"bool", "all", "All workspaces",
		),
		Handler:    listShapersHandler,
		Controller: controllers.ShapersIndex,
		Viewer:     listShapersViewer,
	}
	return
}

func listShapersHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.ShapersIndexParams{
			Workspace: context.StringFlag("workspace"),
		},
	}
	return
}

func listShapersViewer(result *controllers.Result) (output string, err error) {
	rs := result.Payload.([]*controllers.ShapersIndexItemResult)
	items := []map[string]any{}
	for _, r := range rs {
		items = append(items, map[string]any{
			"Workspace": r.Workspace,
			"URI":       r.About,
			"About":     r.About,
		})
	}

	table := &Table{
		Items:  items,
		Titles: ss("WORKSPACE", "SHAPER", "ABOUT"),
		Keys:   ss("Workspace", "URI", "About"),
		Accents: tas(
			tableCellNoAccentFn,
			tableCellNoAccentFn,
			tableCellNoAccentFn,
		),
	}

	result = &controllers.Result{
		Payload: table.generate(),
	}
	output, err = cliapp.View("shapers/index")(result)
	return
}
