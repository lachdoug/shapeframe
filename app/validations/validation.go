package validations

type Validation struct {
	Failures []*InvalidField
}

type InvalidField struct {
	Key     string
	Message string
}

func NewValidation() (vn *Validation) {
	vn = &Validation{}
	return
}

func (vn *Validation) Add(name string, message string) {
	vn.Failures = append(vn.Failures, &InvalidField{name, message})
}

func (vn *Validation) IsInvalid() (is bool) {
	is = !vn.IsValid()
	return
}

func (vn *Validation) IsValid() (is bool) {
	if vn == nil || vn.Failures == nil {
		is = true
	}
	return
}

func (vn *Validation) Maps() (ms []map[string]string) {
	if vn.IsValid() {
		return
	}
	ms = []map[string]string{}
	for _, f := range vn.Failures {
		ms = append(ms, map[string]string{
			"Key":     f.Key,
			"Message": f.Message,
		})
	}
	return
}
