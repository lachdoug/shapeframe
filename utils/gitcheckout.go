package utils

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func GitCheckout(dirPath string, branch string) (err error) {
	var gr *git.Repository
	var gw *git.Worktree
	o := &git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
	}
	if gr, err = git.PlainOpen(dirPath); err != nil {
		return
	} else if gw, err = gr.Worktree(); err != nil {
		return
	} else {
		err = gw.Checkout(o)
	}
	return
}
