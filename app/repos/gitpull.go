package repos

import (
	"sf/app/streams"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func GitPull(dirPath string, username string, password string, st *streams.Stream) {
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
		// Depth:      1,
		Progress: st.Writer,
	}
	if gr, err = git.PlainOpen(dirPath); err != nil {
		st.Errorf(err.Error())
	} else if gw, err = gr.Worktree(); err != nil {
		st.Errorf(err.Error())
	} else if err = gw.Pull(o); err != nil {
		if err == git.NoErrAlreadyUpToDate {
			st.Writef("Already up to date\n")
		} else {
			st.Errorf(err.Error())
		}
	}
}
