package app

import (
	"fmt"
	"strings"
)

type Validation struct {
	Failures []*InvalidField
}

type InvalidField struct {
	Key     string
	Message string
}

func (validation *Validation) Add(name string, message string) {
	validation.Failures = append(validation.Failures, &InvalidField{name, message})
}

func (validation *Validation) IsInvalid() (is bool) {
	if validation.Failures != nil {
		is = true
	}
	return
}

func (validation *Validation) Error() (s string) {
	fs := []string{}
	for _, f := range validation.Failures {
		fs = append(fs, fmt.Sprintf("%s %s", f.Key, f.Message))
	}
	s = strings.Join(fs, "\n")
	return
}
