package logs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sf/app/dirs"
)

var logger *log.Logger

func SetLogger() {
	if file, err := os.OpenFile(dirs.WorkspaceDir("sf.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		panic(err)
	} else {
		logger = log.New(file, "", log.LstdFlags)
	}
}

func Print(a ...any) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		prefix := fmt.Sprintf("%s:%d", filepath.Base(file), no)
		a = append([]any{prefix}, a...)
		logger.Println(a...)
	}

}

func Printf(format string, a ...any) {
	Print(fmt.Sprintf(format, a...))
}
