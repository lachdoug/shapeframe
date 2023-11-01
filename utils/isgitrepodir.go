package utils

import "path/filepath"

func IsGitRepoDir(dirPath string) (is bool) {
	if IsDir(filepath.Join(dirPath, ".git")) {
		is = true
		return
	}
	return
}
