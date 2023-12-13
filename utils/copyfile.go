package utils

import "path/filepath"

func CopyFile(fromPath string, toPath string) {
	if !IsFile(fromPath) {
		return
	}
	MakeDir(filepath.Dir(toPath))
	WriteFile(toPath, ReadFile(fromPath))
}
