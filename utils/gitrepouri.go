package utils

import (
	"strings"
)

func GitRepoURI(dirPath string) (uri string, err error) {
	var url string
	if url, err = GitRepoURL(dirPath); err != nil {
		return
	}
	if url[0:4] == "git@" {
		uri = strings.Replace(
			strings.TrimSuffix(
				strings.TrimPrefix(url, "git@"), ".git"), ":", "/", 1)
	} else {
		uri = strings.TrimPrefix(url, "https://")
	}
	return
}
