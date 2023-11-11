package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func listFramers() (command any) {
	command = &cliapp.Command{
		Name:    "framers",
		Summary: "List framers",
		Aliases: ss("fr"),
		Usage: ss(
			"sf list framers [options]",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"List framers in all workspaces by setting the -all flag",
			"  Otherwise lists framers in workspace context",
			"  Overrides -workspace flag",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
			"bool", "all", "All workspaces",
		),
		Parametizer: listFramersParams,
		Controller:  controllers.FramersIndex,
		Viewer:      listFramersViewer,
	}
	return
}

func listFramersParams(context *cliapp.Context) (jparams []byte, err error) {
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

func listFramersViewer(body map[string]any) (output string, err error) {
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
	output, err = cliapp.View("framers/index")(table.generate())
	return
}
