package logs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sf/utils"
)

var logger *log.Logger

func SetLogger() {
	if file, err := os.OpenFile(utils.LogDir("sf.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		panic(err)
	} else {
		logger = log.New(file, "", log.LstdFlags)
	}
}

func Log(args ...any) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		prefix := fmt.Sprintf("%s:%d", filepath.Base(file), no)
		args = append([]any{prefix}, args...)
		logger.Println(args...)
	}

}

func Logf(format string, args ...any) {
	Log(fmt.Sprintf(format, args...))
}
