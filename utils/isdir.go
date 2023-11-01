package utils

import "os"

func IsDir(dirPath string) (is bool) {
	ent, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return
	}
	if ent.IsDir() {
		is = true
	}
	return
}
