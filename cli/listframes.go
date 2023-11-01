package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
)

func listFrames() (command any) {
	command = &cliapp.Command{
		Name:        "frames",
		Summary:     "List frames",
		Aliases:     ss("f"),
		Flags:       ss("bool", "w", "Limit list to workspace"),
		Parametizer: listFramesParams,
		Controller:  controllers.FramesIndex,
		Viewer:      listFramesViewer,
	}
	return
}

func listFramesParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	jparams = jsonParams(map[string]any{
		"Workspace": context.BoolFlag("w"),
	})
	return
}

func listFramesViewer(body map[string]any) (output string, er error) {
	table := &Table{
		Items:  resultItems(body),
		Titles: ss(" ", "FRAME", "WORKSPACE", "FRAMER", "ABOUT"),
		Keys:   ss("IsContext", "Name", "Workspace", "Framer", "About"),
		Values: tvs(
			tableCellAsteriskIfTrueValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
		),
		Accents: tas(
			tableCellNoAccentFn,
			tableCellGreenIfInContextAccentFn,
			tableCellGreenIfInContextAccentFn,
			tableCellNoAccentFn,
			tableCellNoAccentFn,
		),
	}
	return cliapp.View("frames/index")(table.generate())
}
