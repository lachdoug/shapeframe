package utils

import (
	"os"
	"path/filepath"
)

func MakeDir(elem ...string) {
	dirPath := filepath.Join(elem...)
	if IsDir(dirPath) {
		return
	}
	err := os.MkdirAll(dirPath, os.ModePerm)
	checkErr(err)
}
