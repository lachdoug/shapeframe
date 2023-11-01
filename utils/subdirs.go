package utils

import (
	"io/fs"
	"os"
)

func SubDirs(dirPath string) (dirNames []string) {
	var ents []fs.DirEntry
	if !IsDir(dirPath) {
		return
	}
	ents, err := os.ReadDir(dirPath)
	checkErr(err)
	dirNames = []string{}
	for _, ent := range ents {
		if ent.IsDir() {
			dirNames = append(dirNames, ent.Name())
		}
	}
	return
}
