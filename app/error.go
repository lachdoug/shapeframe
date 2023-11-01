package app

import "fmt"

func Error(errin error, f string, a ...any) (errout error) {
	msg := fmt.Sprintf(f, a...)
	if errin == nil {
		errout = fmt.Errorf(msg)
	} else {
		errout = fmt.Errorf("%s: %s", msg, errin.Error())
	}
	return
}

func NewConfigValidationError(kind string, name string, messages []string) (err error) {
	err = &ConfigValidationError{
		Kind:     kind,
		Name:     name,
		Messages: messages,
	}
	return
}

type ConfigValidationError struct {
	Kind     string
	Name     string
	Messages []string
}

func (err *ConfigValidationError) Error() (s string) {
	s = fmt.Sprintf("%s %s configuration validation", err.Kind, err.Name)
	for _, m := range err.Messages {
		s = fmt.Sprintf("%s\n  %s", s, m)
	}
	return
}
