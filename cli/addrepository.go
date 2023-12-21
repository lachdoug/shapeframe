package cli

import (
	"sf/app/io"
	"sf/cli/cliapp"
	"sf/controllers"
)

func addRepository() (command any) {
	command = &cliapp.Command{
		Name:    "repository",
		Summary: "Add a repository to workspace",
		Aliases: ss("r"),
		Usage: ss(
			"sf add repository [options] URI",
			"A repository URI must be provided as an argument",
			"Clone with https by setting the -https flag",
			"  Otherwise performs git clone using SSH",
			"Include a username for HTTPS clone by setting the -username flag",
			"  Otherwise performs git clone without a token when using HTTPS",
			"Include a password (or personal access token) for HTTPS clone by setting the -password flag",
			"  Otherwise performs git clone without a password when using HTTPS",
		),
		Flags: ss(
			"bool", "https H", "Use HTTPS for git clone",
			"string", "username u", "Username for git clone",
			"string", "password p", "Password for git clone",
		),
		Handler:    addRepositoryHandler,
		Controller: controllers.RepositoriesCreate,
		Viewer:     addRepositoryViewer,
	}
	return
}

func addRepositoryHandler(context *cliapp.Context) (params *controllers.Params, err error) {
	var protocol string

	if context.BoolFlag("https") {
		protocol = "HTTPS"
	}

	params = &controllers.Params{
		Payload: &controllers.RepositoriesCreateParams{
			URI:      context.Argument(0),
			Protocol: protocol,
			Username: context.StringFlag("username"),
			Password: context.StringFlag("password"),
		},
	}
	return
}

func addRepositoryViewer(result *controllers.Result) (output string, err error) {
	r := result.Payload.(*controllers.RepositoriesCreateResult)

	// Stream clone output
	io.Printf("Clone %s\n", r.URL)
	if err = result.Stream.Print(); err != nil {
		return
	}

	output, err = cliapp.View("repositories/create")(result)
	return
}
