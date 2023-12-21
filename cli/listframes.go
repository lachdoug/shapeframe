package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func listFrames() (command any) {
	command = &cliapp.Command{
		Name:    "frames",
		Summary: "List frames",
		Aliases: ss("f"),
		Usage: ss(
			"sf list frames [options]",
		),
		Controller: controllers.FramesIndex,
		Viewer:     listFramesViewer,
	}
	return
}

func listFramesViewer(result *controllers.Result) (output string, err error) {
	rs := result.Payload.([]*controllers.FramesIndexItemResult)
	items := []map[string]any{}
	for _, r := range rs {
		items = append(items, map[string]any{
			"Workspace": r.Workspace,
			"Frame":     r.Frame,
			"Framer":    r.Framer,
			"About":     r.About,
			"IsContext": r.IsContext,
		})
	}

	table := &Table{
		Items:  items,
		Titles: ss("WORKSPACE", "FRAME", "FRAMER", "ABOUT"),
		Keys:   ss("Workspace", "Frame", "Framer", "About"),
		Accents: tas(
			tableCellGreenIfInContextAccentFn,
			tableCellGreenIfInContextAccentFn,
			tableCellNoAccentFn,
			tableCellNoAccentFn,
		),
	}

	result = &controllers.Result{
		Payload: table.generate(),
	}
	output, err = cliapp.View("frames/index")(result)
	return
}
