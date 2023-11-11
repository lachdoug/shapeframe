package utils

import (
	"os"
)

func AppendFile(filePath string, p []byte) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	checkErr(err)
	defer func() {
		err = f.Close()
		checkErr(err)
	}()

	_, err = f.WriteString(string(p))
	checkErr(err)
}
