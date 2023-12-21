package dirs

import (
	"os"
	"path/filepath"
	"sf/utils"
)

var WorkspaceRoot, TempRoot string

func SetDirs(directory string) {
	var err error
	if WorkspaceRoot, err = filepath.Abs(directory); err != nil {
		panic(err)
	}
	TempRoot = tempLocation()
	utils.MakeDir(TempRoot)
}

func WorkspaceDir(relativePath string) (dirPath string) {
	dirPath = filepath.Join(WorkspaceRoot, relativePath)
	return
}

func TempDir(relativePath string) (dirPath string) {
	dirPath = filepath.Join(tempLocation(), ".shapeframe", relativePath)
	return
}

func tempLocation() (dirPath string) {
	dirPath = os.Getenv("SHAPEFRAME_TEMP_DIR")
	if dirPath == "" {
		dirPath = filepath.Dir(os.TempDir())
	}
	if dirPath == "/" {
		dirPath = "/tmp"
	}
	return
}
