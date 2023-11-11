package utils

import (
	"os"
	"path/filepath"
)

func MoveDir(from string, to string) {
	// Remove existing dir if exists
	RemoveDir(to)
	// Make parent dir if not exists
	MakeDir(filepath.Dir(to))
	// And move the dir
	err := os.Rename(from, to)
	checkErr(err)
}
