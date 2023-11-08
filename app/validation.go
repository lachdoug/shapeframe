package app

import (
	"fmt"
)

type Validation struct {
	Failures []*InvalidField
}

type InvalidField struct {
	Key     string
	Message string
}

func (v *Validation) Add(name string, message string) {
	v.Failures = append(v.Failures, &InvalidField{name, message})
}

func (v *Validation) IsInvalid() (is bool) {
	if v.Failures != nil {
		is = true
	}
	return
}

func (v *Validation) Error() (s string) {
	for _, f := range v.Failures {
		s = s + fmt.Sprintf("\n  - %s %s", f.Key, f.Message)
	}
	return
}
