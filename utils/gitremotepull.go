package utils

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func GitRemotePull(dirPath string, username string, password string, st *Stream) {
	defer st.Close()
	var gr *git.Repository
	var gw *git.Worktree
	var err error
	o := &git.PullOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
		RemoteName: "origin",
		Depth:      1,
		Progress:   st,
	}
	if gr, err = git.PlainOpen(dirPath); err != nil {
		st.Error(err)
	} else if gw, err = gr.Worktree(); err != nil {
		st.Error(err)
	} else if err = gw.Pull(o); err != nil {
		if err == git.NoErrAlreadyUpToDate {
			st.Writef("Already up to date\n")
		} else {
			st.Error(err)
		}
	}
}
