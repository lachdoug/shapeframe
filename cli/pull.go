package cli

import (
	"sf/app/io"
	"sf/cli/cliapp"
	"sf/controllers"
)

func pull() (command any) {
	command = &cliapp.Command{
		Name:    "pull",
		Summary: "Pull a workspace repository",
		Aliases: ss("p"),
		Usage: ss(
			"sf pull [options] URI",
			"A repository URI must be provided as an argument",
			"Include a username for HTTPS pull by setting the -username flag",
			"  Otherwise performs git pull without a password when using HTTPS",
			"Include a password (or access token) for HTTPS pull by setting the -password flag",
			"  Otherwise performs git pull without a password when using HTTPS",
		),
		Flags: ss(
			"string", "username u", "Username for git pull",
			"string", "password p", "Password for git pull",
		),
		Handler:    pullHandler,
		Controller: controllers.RepositoryPullsCreate,
		Viewer:     pullViewer,
	}
	return
}

func pullHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	params = &controllers.Params{
		Payload: &controllers.RepositoryPullsCreateParams{
			URI:      context.Argument(0),
			Username: context.StringFlag("username"),
			Password: context.StringFlag("password"),
		},
	}
	return
}

func pullViewer(result *controllers.Result) (output string, err error) {
	r := result.Payload.(*controllers.RepositoryPullsCreateResult)

	// Stream pull output
	io.Printf("Pull %s\n", r.URL)
	if err = result.Stream.Print(); err != nil {
		return
	}

	output, err = cliapp.View("repositorypulls/create")(result)
	return
}
