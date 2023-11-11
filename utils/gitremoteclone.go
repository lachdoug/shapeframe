package utils

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func GitRemoteClone(dirPath string, url string, username string, password string, st *Stream) {
	defer st.Close()
	var err error
	tmp := TempDir("clone")
	RemoveDir(tmp)
	MakeDir(tmp)
	o := &git.CloneOptions{
		URL: url,
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},

		Depth:    1,
		Progress: st,
	}
	if _, err = git.PlainClone(tmp, false, o); err != nil {
		st.Error(err)
	} else {
		MoveDir(tmp, dirPath)
	}
}
