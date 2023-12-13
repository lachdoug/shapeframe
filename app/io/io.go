package io

import (
	"os"
)

var Out *os.File
var In *os.File
var Err *os.File

func SetIO(o *os.File, i *os.File, e *os.File) {
	Out = o
	In = i
	Err = e
}
