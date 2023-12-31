package utils

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func GitURL(dirPath string) (url string, err error) {
	var g *git.Repository
	var gr *git.Remote
	var grc *config.RemoteConfig
	if g, err = git.PlainOpen(dirPath); err != nil {
		err = fmt.Errorf("git url open %s: %s", dirPath, err)
		return
	} else if gr, err = g.Remote("origin"); err != nil {
		url = ""
		return
	}
	grc = gr.Config()
	url = grc.URLs[0]
	if url == "" {
		err = fmt.Errorf("git url origin %s", dirPath)
	}
	return
}
