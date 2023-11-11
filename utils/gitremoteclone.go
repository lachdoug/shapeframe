package utils

import (
	"github.com/go-git/go-git/v5"
)

func GitRemoteClone(dirPath string, url string, st *Stream) {
	var err error
	tmp := TempDir("clone")
	RemoveDir(tmp)
	MakeDir(tmp)
	o := &git.CloneOptions{
		URL:      url,
		Depth:    1,
		Progress: st,
	}
	if _, err = git.PlainClone(tmp, false, o); err != nil {
		st.Error(err)
	}
	MoveDir(tmp, dirPath)
	st.Close()
}
