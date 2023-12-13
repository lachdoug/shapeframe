package models

import "sf/app/validations"

type FormOption struct {
	Value string
	Label string
}

func (cspos *FormOption) validation(
	id string,
	vn *validations.Validation,
) (msgs []string) {
	if cspos.Value == "" {
		vn.Add(id, "must have a value")
	}
	return
}
