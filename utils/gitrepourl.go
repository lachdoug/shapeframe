package utils

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func GitRepoURL(dirPath string) (url string, err error) {
	var g *git.Repository
	var gr *git.Remote
	var grc *config.RemoteConfig
	if g, err = git.PlainOpen(dirPath); err != nil {
		err = fmt.Errorf("git repo directory %s: %s", dirPath, err)
		return
	} else if gr, err = g.Remote("origin"); err != nil {
		url = ""
		return
	}
	grc = gr.Config()
	url = grc.URLs[0]
	if url == "" {
		err = fmt.Errorf("git repo directory %s: no origin url", dirPath)
	}
	return
}
