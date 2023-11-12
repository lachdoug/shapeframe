package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
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
		Parametizer: listShapersParams,
		Controller:  controllers.ShapersIndex,
		Viewer:      listShapersViewer,
	}
	return
}

func listShapersParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	all := context.BoolFlag("all")
	workspace := context.StringFlag("workspace")
	params := map[string]any{}

	if !all {
		uc := models.ResolveUserContext(
			"Workspaces", "Workspace",
		)
		if w, err = models.ResolveWorkspace(uc, workspace); err != nil {
			return
		}
		params["Workspace"] = w.Name
	}

	jparams = jsonParams(params)
	return
}

func listShapersViewer(body map[string]any) (output string, err error) {
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
	output, err = cliapp.View("shapers/index")(table.generate())
	return
}
