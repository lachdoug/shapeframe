package prompting

import (
	"fmt"
	"sf/app"
	"sf/models"
	"strconv"
	"strings"
)

func RepositoryURI(w *models.Workspace) (uri string, err error) {
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
	s, err := Prompt("Which repository?")
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
