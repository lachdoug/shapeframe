package app

import (
	"fmt"
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

func Print(a ...any) {
	if _, err := fmt.Fprint(Out, a...); err != nil {
		panic(err)
	}
}

func Println(a ...any) {
	if _, err := fmt.Fprintln(Out, a...); err != nil {
		panic(err)
	}
}

func Printf(format string, a ...any) {
	if _, err := fmt.Fprintf(Out, format, a...); err != nil {
		panic(err)
	}
}

func PrintErr(a ...any) {
	if _, err := fmt.Fprint(Err, a...); err != nil {
		panic(err)
	}
}

func PrintlnErr(a ...any) {
	if _, err := fmt.Fprintln(Err, a...); err != nil {
		panic(err)
	}
}

func PrintfErr(format string, a ...any) {
	if _, err := fmt.Fprintf(Err, format, a...); err != nil {
		panic(err)
	}
}
