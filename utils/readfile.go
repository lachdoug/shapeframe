package utils

import "os"

func ReadFile(filePath string) (content []byte) {
	content, err := os.ReadFile(filePath)
	checkErr(err)
	return
}
