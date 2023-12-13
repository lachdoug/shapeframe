package utils

import (
	"fmt"
)

func FixedLengthString(in string, length int) (out string) {
	if length == 0 {
		out = ""
	} else if len(in) > length {
		format := fmt.Sprintf("%%.%ds…", length-1)
		out = fmt.Sprintf(format, in)
	} else {
		format := fmt.Sprintf("%%-%ds", length)
		out = fmt.Sprintf(format, in)
	}
	return
}
