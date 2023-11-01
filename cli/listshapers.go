package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
)

func listShapers() (command any) {
	command = &cliapp.Command{
		Name:        "shapers",
		Summary:     "List shapers",
		Aliases:     ss("sr"),
		Flags:       ss("bool", "w", "Limit list to workspace"),
		Parametizer: listShapersParams,
		Controller:  controllers.ShapersIndex,
		Viewer:      listShapersViewer,
	}
	return
}

func listShapersParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	jparams = jsonParams(map[string]any{
		"Workspace": context.BoolFlag("w"),
	})
	return
}

func listShapersViewer(body map[string]any) (output string, er error) {
	table := &Table{
		Items:  resultItems(body),
		Titles: ss("WORKSPACE", "SHAPER", "ABOUT"),
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
	return cliapp.View("shapers/index")(table.generate())
}
