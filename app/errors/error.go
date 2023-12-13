package errors

import (
	"fmt"
	"os"
	"sf/app/io"
	"strings"

	"github.com/go-errors/errors"
)

var AppError *errors.Error

func ValidationError(vnmaps []map[string]string) error {
	s := ""
	for _, vnmap := range vnmaps {
		key := vnmap["Key"]
		msg := vnmap["Message"]
		s = s + fmt.Sprintf("\n  - %s %s", key, msg)
	}
	return Error("invalid:" + s)
}

func ErrorWrap(err error, msgs ...string) error {
	msgs = append(msgs, err.Error())
	return appError(strings.Join(msgs, ": "))
}

func ErrorWrapf(err error, f string, a ...any) error {
	msg := fmt.Sprintf("%s: %s", fmt.Sprintf(f, a...), err.Error())
	return appError(msg)
}

func Errorf(f string, a ...any) error {
	msg := fmt.Sprintf(f, a...)
	return appError(msg)
}

func Error(msg string) error {
	return appError(msg)
}

func appError(msg string) error {
	AppError = errors.New(msg)
	return AppError
}

func ErrorHandler(err error) {
	if err == nil {
		panic("err should not be nil")
	}
	if AppError == nil {
		panic(fmt.Sprintf("AppErr should not be nil for err %s", err))
	}
	if errors.Is(err, AppError) {
		io.PrintlnErr("Error:", err)
		if Debug {
			stack := string(err.(*errors.Error).Stack())
			// Skip first four lines since they merely show AppError being created.
			lines := strings.Split(stack, "\n")[4:]
			io.PrintErr(strings.Join(lines, "\n"))
		}
	} else {
		panic(err)
	}
	os.Exit(1)
}
