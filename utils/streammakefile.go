package utils

import (
	"path/filepath"
)

func StreamMakeFile(dirPath string) {
	filePath := filepath.Join(dirPath, "out")
	MakeFile(filePath)
}
