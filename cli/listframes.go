package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
)

func listFrames() (command any) {
	command = &cliapp.Command{
		Name:    "frames",
		Summary: "List frames",
		Aliases: ss("f"),
		Usage: ss(
			"sf list frames [options]",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
			"List frames in all workspaces by setting the -all flag",
			"  Otherwise lists frames in workspace context",
			"  Overrides -workspace flag",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
			"bool", "all", "All workspaces",
		),
		Parametizer: listFramesParams,
		Controller:  controllers.FramesIndex,
		Viewer:      listFramesViewer,
	}
	return
}

func listFramesParams(context *cliapp.Context) (jparams []byte, err error) {
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

func listFramesViewer(body map[string]any) (output string, err error) {
	table := &Table{
		Items:  resultItems(body),
		Titles: ss("FRAME", "WORKSPACE", "FRAMER", "ABOUT"),
		Keys:   ss("Name", "Workspace", "Framer", "About"),
		Values: tvs(
			tableCellStringValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
			tableCellStringValueFn,
		),
		Accents: tas(
			tableCellGreenIfInContextAccentFn,
			tableCellGreenIfInContextAccentFn,
			tableCellNoAccentFn,
			tableCellNoAccentFn,
		),
	}
	output, err = cliapp.View("frames/index")(table.generate())
	return
}
