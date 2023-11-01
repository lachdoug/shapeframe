package utils

import (
	"os"
	"path/filepath"
)

func StreamAppend(dirPath string, p []byte) {
	var f *os.File
	var err error

	filePath := filepath.Join(dirPath, "out")
	f, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	checkErr(err)
	defer f.Close()

	_, err = f.WriteString(string(p))
	checkErr(err)
}
