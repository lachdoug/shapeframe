package repos

import (
	"sf/app/streams"
	"sf/utils"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func GitClone(dirPath string, url string, username string, password string, st *streams.Stream) {
	defer st.Close()
	var err error
	tmp := utils.TempDir("clone")
	utils.RemoveDir(tmp)
	utils.MakeDir(tmp)
	o := &git.CloneOptions{
		URL: url,
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},

		// Depth:    1,
		Progress: st.Writer,
	}
	if _, err = git.PlainClone(tmp, false, o); err != nil {
		st.Errorf(err.Error())
	} else {
		utils.MoveDir(tmp, dirPath)
	}
}
