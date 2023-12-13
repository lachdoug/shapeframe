package utils

import (
	"os"
	"path/filepath"
)

func DataDir(relativePath string) (dirPath string) {
	dirPath = filepath.Join(DataDirLocation(), ".shapeframe", relativePath)
	return
}

func TempDir(relativePath string) (dirPath string) {
	dirPath = filepath.Join(TempDirLocation(), ".shapeframe", relativePath)
	return
}

func LogDir(relativePath string) (dirPath string) {
	dirPath = filepath.Join(LogDirLocation(), ".shapeframe", relativePath)
	return
}

func DataDirLocation() (dirPath string) {
	dirPath = os.Getenv("SHAPEFRAME_DATA_DIR")
	if dirPath == "" {
		if execFilePath, err := os.Executable(); err != nil {
			panic(err)
		} else {
			dirPath = filepath.Dir(execFilePath)
		}
	}
	return
}

func TempDirLocation() (dirPath string) {
	dirPath = os.Getenv("SHAPEFRAME_TEMP_DIR")
	if dirPath == "" {
		dirPath = filepath.Dir(os.TempDir())
	}
	if dirPath == "/" {
		dirPath = "/tmp"
	}
	return
}

func LogDirLocation() (dirPath string) {
	dirPath = os.Getenv("SHAPEFRAME_LOG_DIR")
	if dirPath == "" {
		dirPath = filepath.Dir(os.TempDir())
	}
	if dirPath == "/" {
		dirPath = "/tmp"
	}
	return
}
