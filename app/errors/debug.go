package errors

var Debug bool

func SetDebug(args []string) {
	if len(args) > 1 && args[1] == "--debug" {
		Debug = true
	}
}
