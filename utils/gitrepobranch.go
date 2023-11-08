package utils

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func GitRepoBranch(dirPath string) (branch string, err error) {
	var g *git.Repository
	var gr *plumbing.Reference
	if g, err = git.PlainOpen(dirPath); err != nil {
		panic("WTF")
		err = fmt.Errorf("git repo directory %s: %s", dirPath, err)
		return
	}
	if gr, err = g.Head(); err != nil {
		err = fmt.Errorf("git repo directory %s: %s", dirPath, err)
		return
	}
	branch = gr.Name().Short()
	return
}
