package cli

import (
	"fmt"
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
	"strconv"
	"strings"
)

func pull() (command any) {
	command = &cliapp.Command{
		Name:    "pull",
		Summary: "Pull a workspace repository",
		Aliases: ss("p"),
		Usage: ss(
			"sf pull [options] [URI]",
			"A repository URI may be provided as an argument",
			"  Otherwise prompt for URI",
			"Provide an optional workspace name using the -workspace flag",
			"  Uses workspace context when not provided",
		),
		Flags: ss(
			"string", "workspace", "Workspace name",
		),
		Parametizer: pullParams,
		Controller:  controllers.RepositoryPullsCreate,
		Viewer:      pullViewer,
	}
	return
}

func pullParams(context *cliapp.Context) (jparams []byte, err error) {
	var w *models.Workspace
	workspace := context.StringFlag("workspace")
	uri := context.Argument(0)

	uc := models.ResolveUserContext("Workspace")
	if w, err = models.ResolveWorkspace(uc, workspace, "Repositories"); err != nil {
		return
	}

	if uri == "" {
		if uri, err = pullPrompt(w); err != nil {
			return
		}
	}

	jparams = jsonParams(map[string]any{
		"Workspace": w.Name,
		"URI":       uri,
	})
	return
}

func pullPrompt(w *models.Workspace) (uri string, err error) {
	rs := w.Repositories
	if len(rs) == 0 {
		err = app.Error("no repositories in workspace")
		return
	}
	list := ""
	uris := []string{}
	for i, r := range rs {
		uris = append(uris, r.URI)
		list = list + fmt.Sprintf("%d. %s\n", i+1, r.URI)
	}
	app.Printf(list)
	s, err := prompt("Which repository?")
	if err != nil {
		return
	}
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err == nil && i <= len(rs) {
		uri = uris[i-1]
	} else {
		err = app.Error("invalid: %s", s)
		return
	}
	return
}

func pullViewer(body map[string]any) (output string, err error) {
	result := resultItem(body)
	app.Printf("Pull %s\n", result["URL"])
	if err = stream(body); err != nil {
		return
	}
	output, err = cliapp.View("repositorypulls/create")(body)
	return
}
