package cli

import (
	"sf/cli/cliapp"
	"sf/cli/prompting"
	"sf/controllers"
	"sf/models"
)

func checkout() (command any) {
	command = &cliapp.Command{
		Name:    "checkout",
		Summary: "Checkout a repository branch",
		Aliases: ss("ck"),
		Usage: ss(
			"sf branch [options] [URI]",
			"A repository URI may be provided as an argument",
			"  Otherwise prompt for URI",
			"Provide an optional branch name using the -branch flag",
			"  Otherwise prompt for branch",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
		),
		Flags: ss(
			"string", "branch", "Branch name to checkout",
			"string", "workspace", "Workspace name",
		),
		Parametizer: checkoutParams,
		Controller:  controllers.RepositoryBranchesUpdate,
		Viewer:      cliapp.View("repositorybranches/update"),
	}
	return
}

func checkoutParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	uri := context.Argument(0)
	branch := context.StringFlag("branch")
	workspace := context.StringFlag("workspace")

	uc := models.ResolveUserContext("Workspace", "Workspaces")
	if w, err = models.ResolveWorkspace(uc, workspace,
		"Repositories",
	); err != nil {
		return
	}

	if uri == "" {
		if uri, err = prompting.RepositoryURI(w); err != nil {
			return
		}
	}

	if branch == "" {
		if branch, err = prompting.Prompt("Branch name to checkout?"); err != nil {
			return
		}
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"URI":       uri,
		"Branch":    branch,
	})
	return
}
