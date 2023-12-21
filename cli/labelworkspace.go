package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func labelWorkspace() (command any) {
	command = &cliapp.Command{
		Name:    "workspace",
		Summary: "Change name and/or about for the workspace",
		Aliases: ss("w"),
		Usage: ss(
			"sf label workspace [options]",
			"Provide an optional workspace update name using the -name flag",
			"Provide an optional workspace update about using the -about flag",
		),
		Flags: ss(
			"string", "name", "New name for the workspace",
			"string", "about", "New about for the workspace",
		),
		Handler:    labelWorkspaceHandler,
		Controller: controllers.WorkspacesUpdate,
		Viewer:     cliapp.View("workspaces/update", "labels/label"),
	}
	return
}

func labelWorkspaceHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	updates := map[string]string{}
	if context.IsSet("name") {
		updates["Name"] = context.StringFlag("name")
	}
	if context.IsSet("about") {
		updates["About"] = context.StringFlag("about")
	}

	params = &controllers.Params{
		Payload: &controllers.WorkspacesUpdateParams{
			Updates: updates,
		},
	}
	return
}
