package errors

var Debug bool

func SetDebug(is bool) {
	if is {
		Debug = true
	}
}
