package models

import (
	"fmt"
	"regexp"
	"sf/app/validations"
)

type FormComponent struct {
	Type     string
	Key      string
	As       string
	Label    string
	Default  string
	Required bool
	Pattern  string
	Invalid  string
	Width    int
	Depend   *FormDepend
	Options  []*FormOption
	Cols     []*FormComponent
}

func (fmc *FormComponent) validation(id string, vn *validations.Validation) (msgs []string) {
	var err error
	if fmc.Type != "row" {
		if fmc.Key == "" {
			vn.Add(id, "must have a key")
		}
		if fmc.Pattern != "" {
			if _, err = regexp.Compile(fmc.Pattern); err != nil {
				vn.Add(id, "must have a valid regexp")
			}
		}
		if fmc.Options != nil {
			for i, fmo := range fmc.Options {
				fmo.validation(fmt.Sprintf("%s options[%d]", id, i), vn)
			}
		}
	}
	return
}

func (fmc *FormComponent) Load() {
	fmc.setType()
	fmc.setWidth()
	switch fmc.Type {
	case "row":
		fmc.loadRow()
	default:
		fmc.loadField()
	}
}

func (fmc *FormComponent) setType() {
	if fmc.Type == "" {
		fmc.Type = "string"
	}
}

func (fmc *FormComponent) setWidth() {
	if fmc.Width == 0 {
		fmc.Width = 12
	}
}

func (fmc *FormComponent) loadField() {
	fmc.setAs()
	fmc.setLabel()
	fmc.setInvalid()
	fmc.setDepend()
}

func (fmc *FormComponent) loadRow() {
	for _, fmc := range fmc.Cols {
		fmc.Load()
	}
}

func (fmc *FormComponent) setAs() {
	if fmc.As == "" {
		fmc.As = fmc.defaultAs()
	}
}

func (fmc *FormComponent) setDepend() {
	if fmc.Depend != nil && fmc.Depend.Pattern == "" {
		fmc.Depend.Pattern = ".+"
	}
}

func (fmc *FormComponent) defaultAs() (a string) {
	switch fmc.Type {
	case "select", "selects", "text":
		a = fmc.Type
	default:
		a = "input"
	}
	return
}

func (fmc *FormComponent) setLabel() {
	if fmc.Label == "" {
		fmc.Label = fmc.Key
	}
}

func (fmc *FormComponent) setInvalid() {
	if fmc.Pattern != "" && fmc.Invalid == "" {
		fmc.Invalid = fmt.Sprintf("must match pattern /%s/", fmc.Pattern)
	}
}

func (fmc *FormComponent) ValueValidation(v string, vn *validations.Validation) *validations.Validation {
	if fmc.Required && v == "" {
		vn.Add(fmc.Key, "must not be blank")
	} else if !fmc.isMatch(v) {
		vn.Add(fmc.Key, fmc.Invalid)
	}
	return vn
}

func (fmc *FormComponent) isMatch(v string) (is bool) {
	if fmc.Pattern == "" {
		is = true
		return
	}
	var r *regexp.Regexp
	var err error
	if r, err = regexp.Compile(fmc.Pattern); err != nil {
		panic(err)
	}
	is = r.MatchString(v)
	return
}
