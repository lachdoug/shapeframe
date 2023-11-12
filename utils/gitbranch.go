package utils

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func GitBranch(dirPath string) (branch string, err error) {
	var g *git.Repository
	var ref *plumbing.Reference
	if g, err = git.PlainOpen(dirPath); err != nil {
		err = fmt.Errorf("git open %s: %s", dirPath, err)
		return
	}
	if ref, err = g.Head(); err != nil {
		err = fmt.Errorf("git read %s: %s", dirPath, err)
		return
	}
	branch = ref.Name().Short()
	return
}
