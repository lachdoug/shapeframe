package utils

import "regexp"

func MaskSecret(in string) (out string) {
	anyCharRegexp := regexp.MustCompile(`.`)
	out = anyCharRegexp.ReplaceAllString(in, "*")
	return
}
