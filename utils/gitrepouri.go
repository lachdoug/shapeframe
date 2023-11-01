package utils

import (
	"strings"
)

func GitRepoURI(dirPath string) (uri string) {
	url := GitRepoURL(dirPath)
	if url[0:4] == "git@" {
		uri = strings.Replace(
			strings.TrimSuffix(
				strings.TrimPrefix(url, "git@"), ".git"), ":", "/", 1)
	} else {
		uri = strings.TrimPrefix(url, "https://")
	}
	return
}
