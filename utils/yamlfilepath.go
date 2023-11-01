package utils

import "path/filepath"

func YamlFilePath(dirPath string, fileName string) (filePath string) {
	filePath = filepath.Join(dirPath, fileName+".yaml")
	return
}
