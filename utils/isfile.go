package utils

import "os"

func IsFile(fp string) (is bool) {
	ent, err := os.Stat(fp)
	if os.IsNotExist(err) {
		return
	}
	if !ent.IsDir() {
		is = true
	}
	return
}
