package cli

import (
	"fmt"
	"os"
	"sf/app"
	"sf/cli/cliapp"
	"sf/controllers"
	"sf/models"
	"strconv"
	"strings"
)

func pull() (command any) {
	command = &cliapp.Command{
		Name:        "pull",
		Summary:     "Pull a workspace repository",
		Parametizer: pullParams,
		Controller:  controllers.RepositoryPullsUpdate,
		Viewer:      pullViewer,
	}
	return
}

func pullParams(context *cliapp.Context) (jparams []byte, validation *app.Validation, err error) {
	uc := models.UserContextNew()
	uc.Load("Workspace")
	w := uc.Workspace
	if w == nil {
		err = app.Error(nil, "no workspace context")
		return
	}

	uri := context.Argument(0)
	if len(uri) == 0 {
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
	rs := w.RepositoriesWithGitRepos()
	if len(rs) == 0 {
		app.Error(nil, "no repositories in workspace")
		return
	}
	q := ""
	uris := []string{}
	for i, r := range rs {
		uris = append(uris, r.GitRepo.URI)
		q = q + fmt.Sprintf("%d. %s\n", i+1, r.GitRepo.URI)
	}
	s := prompt(q + "Which repository? ")
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err == nil && i <= len(rs) {
		uri = uris[i-1]
	} else {
		fmt.Printf("Invalid: %s", s)
		os.Exit(1)
	}
	return
}

func pullViewer(body map[string]any) (output string, err error) {
	r := resultItem(body)
	fmt.Printf("Pull %s\n", r["GitURL"])
	if err = stream(body); err != nil {
		return
	}
	output = fmt.Sprintf("Updated %s in workspace %s\n", r["URI"], r["Workspace"])
	return
}
