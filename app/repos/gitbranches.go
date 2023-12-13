package repos

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

func GitBranches(dirPath string) (branches []string, err error) {
	var g *git.Repository
	var refIter storer.ReferenceIter
	branches = []string{}
	if g, err = git.PlainOpen(dirPath); err != nil {
		err = fmt.Errorf("git open %s: %s", dirPath, err)
		return
	}
	if refIter, err = g.Branches(); err != nil {
		err = fmt.Errorf("git read %s: %s", dirPath, err)
		return
	}
	refIter.ForEach(func(gr *plumbing.Reference) (err error) {
		if gr.Type() == plumbing.HashReference {
			branches = append(branches, gr.Name().Short())
		}
		return
	})
	return
}
