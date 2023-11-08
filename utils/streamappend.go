package utils

import (
	"os"
	"path/filepath"
)

func StreamAppend(dirPath string, p []byte) {
	filePath := filepath.Join(dirPath, "out")
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	checkErr(err)
	defer func() {
		err = f.Close()
		checkErr(err)
	}()

	_, err = f.WriteString(string(p))
	checkErr(err)
}
