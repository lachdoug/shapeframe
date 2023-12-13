package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func listFramers() (command any) {
	command = &cliapp.Command{
		Name:    "framers",
		Summary: "List framers",
		Aliases: ss("fr"),
		Usage: ss(
			"sf list framers [options]",
			"Provide an optional workspace name using the -workspace flag",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
		),
		Handler:    listFramersHandler,
		Controller: controllers.FramersIndex,
		Viewer:     listFramersViewer,
	}
	return
}

func listFramersHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.FramersIndexParams{
			Workspace: context.StringFlag("workspace"),
		},
	}
	return
}

func listFramersViewer(result *controllers.Result) (output string, err error) {
	rs := result.Payload.([]*controllers.FramersIndexItemResult)
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
		Titles: ss("WORKSPACE", "FRAMER", "ABOUT"),
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
	output, err = cliapp.View("framers/index")(result)
	return
}
