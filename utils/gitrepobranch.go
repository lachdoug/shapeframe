package utils

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func GitRepoBranch(dirPath string) (branch string) {
	var g *git.Repository
	var gr *plumbing.Reference
	var err error
	if g, err = git.PlainOpen(dirPath); err != nil {
		panic(err)
	}
	if gr, err = g.Head(); err != nil {
		panic(err)
	}
	branch = gr.Name().Short()
	return
}
