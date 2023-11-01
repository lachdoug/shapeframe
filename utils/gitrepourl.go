package utils

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func GitRepoURL(dirPath string) (url string) {
	var g *git.Repository
	var gr *git.Remote
	var grc *config.RemoteConfig
	var err error
	if g, err = git.PlainOpen(dirPath); err != nil {
		panic(err)
	} else if gr, err = g.Remote("origin"); err != nil {
		url = ""
		return
	}
	grc = gr.Config()
	url = grc.URLs[0]
	return
}
