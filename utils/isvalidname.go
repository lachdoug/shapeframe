package utils

import "regexp"

func IsValidName(name string) (is bool) {
	validNameRegexp := regexp.MustCompile(`^[\w\d\-_]+$`)
	is = validNameRegexp.MatchString(name)
	return
}
