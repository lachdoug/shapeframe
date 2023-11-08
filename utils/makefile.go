package utils

import (
	"os"
	"path/filepath"
)

func MakeFile(elem ...string) {
	filePath := filepath.Join(elem...)
	if IsFile(filePath) {
		return
	}
	MakeDir(filepath.Dir(filePath))
	f, err := os.Create(filePath)
	checkErr(err)
	err = f.Close()
	checkErr(err)
	return
}
