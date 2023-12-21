package cli

import (
	"sf/cli/cliapp"
	"sf/controllers"
)

func checkout() (command any) {
	command = &cliapp.Command{
		Name:    "checkout",
		Summary: "Checkout a repository branch",
		Aliases: ss("ck"),
		Usage: ss(
			"sf branch [options] URI BRANCH",
			"A repository URI must be provided as the first argument",
			"A branch name must be provided as the second argument",
		),
		Handler:    checkoutHandler,
		Controller: controllers.RepositoryBranchesUpdate,
		Viewer:     cliapp.View("repositorybranches/update"),
	}
	return
}

func checkoutHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.RepositoryBranchesUpdateParams{
			URI:    context.Argument(0),
			Branch: context.Argument(1),
		},
	}
	return
}
