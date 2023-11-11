package app

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

func (v *Validation) IsValid() (is bool) {
	is = !v.IsInvalid()
	return
}

func (v *Validation) Map() (m map[string]string) {
	if v.IsValid() {
		return
	}
	m = map[string]string{}
	for _, f := range v.Failures {
		m[f.Key] = f.Message
	}
	return
}
