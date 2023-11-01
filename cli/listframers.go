package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
)

func listFramers() (command any) {
	command = &cliapp.Command{
		Name:        "framers",
		Summary:     "List framers",
		Aliases:     ss("fr"),
		Flags:       ss("bool", "w", "Limit list to workspace context"),
		Parametizer: listFramersParams,
		Controller:  controllers.FramersIndex,
		Viewer:      listFramersViewer,
	}
	return
}

func listFramersParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	jparams = jsonParams(map[string]any{
		"Workspace": context.BoolFlag("w"),
	})
	return
}

func listFramersViewer(body map[string]any) (output string, er error) {
	table := &Table{
		Items:  resultItems(body),
		Titles: ss("WORKSPACE", "FRAMER", "ABOUT"),
		Keys:   ss("Workspace", "URI", "About"),
		Values: tvs(
			tableCellStringValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
		),
		Accents: tas(
			tableCellNoAccentFn,
			tableCellNoAccentFn,
			tableCellNoAccentFn,
		),
	}
	return cliapp.View("framers/index")(table.generate())
}
