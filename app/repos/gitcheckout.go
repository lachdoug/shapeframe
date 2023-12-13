package repos

import (
	"sf/app/errors"
	"slices"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func GitCheckout(dirPath string, branch string) (err error) {
	var gr *git.Repository
	var gw *git.Worktree
	var branches []string

	if branches, err = GitBranches(dirPath); err != nil {
		return
	}
	if !slices.Contains(branches, branch) {
		err = errors.Errorf("branch %s does not exist in git repo", branch)
		return
	}
	o := &git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
	}
	if gr, err = git.PlainOpen(dirPath); err != nil {
		return
	}
	if gw, err = gr.Worktree(); err != nil {
		return
	}
	err = gw.Checkout(o)
	return
}
