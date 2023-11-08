package utils

import (
	"os"
	"path/filepath"
)

func WriteFile(filePath string, content []byte) {
	MakeDir(filepath.Dir(filePath))
	err := os.WriteFile(filePath, content, 0666)
	checkErr(err)
	return
}
