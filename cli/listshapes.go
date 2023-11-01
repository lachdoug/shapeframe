package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
)

func listShapes() (command any) {
	command = &cliapp.Command{
		Name:        "shapes",
		Summary:     "List shapes",
		Aliases:     ss("s"),
		Flags:       ss("bool", "w", "Limit list to workspace"),
		Parametizer: listShapesParams,
		Controller:  controllers.ShapesIndex,
		Viewer:      listShapesViewer,
	}
	return
}

func listShapesParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	jparams = jsonParams(map[string]any{
		"Workspace": context.BoolFlag("w"),
	})
	return
}

func listShapesViewer(body map[string]any) (output string, er error) {
	table := &Table{
		Items:  resultItems(body),
		Titles: ss(" ", "SHAPE", "FRAME", "WORKSPACE", "SHAPER", "ABOUT"),
		Keys:   ss("IsContext", "Name", "Frame", "Workspace", "Shaper", "About"),
		Values: tvs(
			tableCellAsteriskIfTrueValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
		),
		Accents: tas(
			tableCellNoAccentFn,
			tableCellGreenIfInContextAccentFn,
			tableCellGreenIfInContextAccentFn,
			tableCellGreenIfInContextAccentFn,
			tableCellNoAccentFn,
			tableCellNoAccentFn,
		),
	}
	return cliapp.View("shapes/index")(table.generate())
}
