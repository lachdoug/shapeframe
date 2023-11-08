package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-errors/errors"
)

var Debug bool
var AppError *errors.Error

func SetDebug(args []string) {
	if len(args) > 1 && args[1] == "-debug" {
		Debug = true
	}
}

func ErrorWith(err error, f string, a ...any) error {
	msg := fmt.Sprintf("%s: %s", fmt.Sprintf(f, a...), err.Error())
	return appError(msg)
}

func Error(f string, a ...any) error {
	msg := fmt.Sprintf(f, a...)
	return appError(msg)
}

func appError(msg string) error {
	AppError = errors.New(msg)
	return AppError
}

func ErrorHandler(err error) {
	if errors.Is(err, AppError) {
		PrintfErr("Error: %s\n", err)
		if Debug {
			stack := string(err.(*errors.Error).Stack())
			// Skip first four lines since they merely show AppError being created.
			lines := strings.Split(stack, "\n")[4:]
			PrintErr(strings.Join(lines, "\n"))
		}
	} else {
		panic(err)
	}
	os.Exit(1)
}
