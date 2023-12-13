package scripts

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sf/app/streams"
)

func Exec(filePath string, st *streams.Stream) (isOutput bool, err error) {
	// interpreter := utils.ScriptInterpreter(filePath)

	os.Chmod(filePath, 0700)
	outw := NewOutWriter(st)
	errw := NewErrWriter()
	cmd := exec.Command(filePath)
	cmd.Dir = filepath.Dir(filePath)
	cmd.Stdout = outw
	cmd.Stderr = errw
	if err = cmd.Run(); err != nil {
		err = fmt.Errorf("execute %s%s: %s", filepath.Base(filePath), errw.Error(), err)
		return
	} else if outw.length() > 0 {
		isOutput = true
	}
	return
}
