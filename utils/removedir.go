package utils

import (
	"os"
)

func RemoveDir(p string) {
	if !IsDir(p) {
		return
	}
	err := os.RemoveAll(p)
	checkErr(err)
}
