package cli

import (
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func listShapes() (command any) {
	command = &cliapp.Command{
		Name:    "shapes",
		Summary: "List shapes",
		Aliases: ss("s"),
		Usage: ss(
			"sf list shapes [options]",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"List shapes in all workspaces by setting the -all flag",
			"  Otherwise lists shapes in workspace context",
			"  Overrides -workspace flag",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
			"bool", "all", "All workspaces",
		),
		Parametizer: listShapesParams,
		Controller:  controllers.ShapesIndex,
		Viewer:      listShapesViewer,
	}
	return
}

func listShapesParams(context *cliapp.Context) (jparams []byte, vn *app.Validation, err error) {
	var w *models.Workspace
	all := context.BoolFlag("all")
	workspace := context.StringFlag("workspace")
	params := map[string]any{}

	if !all {
		uc := models.ResolveUserContext(
			"Workspace",
			"Workspaces",
		)
		if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
			return
		}
		params["Workspace"] = w.Name
	}

	jparams = jsonParams(params)
	return
}

func listShapesViewer(body map[string]any) (output string, er error) {
	table := &Table{
		Items:  resultItems(body),
		Titles: ss("SHAPE", "FRAME", "WORKSPACE", "SHAPER", "ABOUT"),
		Keys:   ss("Name", "Frame", "Workspace", "Shaper", "About"),
		Values: tvs(
			tableCellStringValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
		),
		Accents: tas(
			tableCellGreenIfInContextAccentFn,
			tableCellGreenIfInContextAccentFn,
			tableCellGreenIfInContextAccentFn,
			tableCellNoAccentFn,
			tableCellNoAccentFn,
		),
	}
	return cliapp.View("shapes/index")(table.generate())
}
