package utils

import (
	"io/fs"
	"path/filepath"
)

func GitDirs(dirPath string) (gitRepoDirs []string) {
	gitRepoDirs = []string{}
	var walker fs.WalkDirFunc = func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.Name() == ".git" {
			gitRepoDirs = append(gitRepoDirs, filepath.Join(filepath.Dir(path)))
		}
		return nil
	}
	err := filepath.WalkDir(dirPath, walker)
	checkErr(err)
	return
}
