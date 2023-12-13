package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func listShapes() (command any) {
	command = &cliapp.Command{
		Name:    "shapes",
		Summary: "List shapes",
		Aliases: ss("s"),
		Usage: ss(
			"sf list shapes [options]",
			"Provide an optional workspace name using the -workspace flag",
			"Provide an optional frame name using the -frame flag",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
			"string", "frame", "Frame name",
		),
		Handler:    listShapesHandler,
		Controller: controllers.ShapesIndex,
		Viewer:     listShapesViewer,
	}
	return
}

func listShapesHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.ShapesIndexParams{
			Workspace: context.StringFlag("workspace"),
			Frame:     context.StringFlag("frame"),
		},
	}
	return
}

func listShapesViewer(result *controllers.Result) (output string, err error) {
	rs := result.Payload.([]*controllers.ShapesIndexItemResult)
	items := []map[string]any{}
	for _, r := range rs {
		items = append(items, map[string]any{
			"Workspace": r.Workspace,
			"Frame":     r.Frame,
			"Shape":     r.Shape,
			"Shaper":    r.Shaper,
			"About":     r.About,
			"IsContext": r.IsContext,
		})
	}

	table := &Table{
		Items:  items,
		Titles: ss("WORKSPACE", "FRAME", "SHAPE", "SHAPER", "ABOUT"),
		Keys:   ss("Workspace", "Frame", "Shape", "Shaper", "About"),
		Accents: tas(
			tableCellGreenIfInContextAccentFn,
			tableCellGreenIfInContextAccentFn,
			tableCellGreenIfInContextAccentFn,
			tableCellNoAccentFn,
			tableCellNoAccentFn,
		),
	}

	result = &controllers.Result{
		Payload: table.generate(),
	}
	output, err = cliapp.View("shapes/index")(result)
	return
}
