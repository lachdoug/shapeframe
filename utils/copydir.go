package utils

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func CopyDir(fromPath string, toPath string) {
	MakeDir(toPath)
	if !IsDir(fromPath) {
		return
	}
	var walkFunc filepath.WalkFunc = func(sourcePath string, info fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		destPath := filepath.Join(toPath, strings.TrimPrefix(sourcePath, fromPath))
		if IsDir(sourcePath) {
			MakeDir(destPath)
		} else if IsFile(sourcePath) {
			CopyFile(sourcePath, destPath)
		}
		return nil
	}
	filepath.Walk(fromPath, walkFunc)
}
