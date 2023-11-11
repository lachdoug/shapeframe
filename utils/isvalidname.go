package utils

import "regexp"

func IsValidName(name string) (is bool) {
	validName := regexp.MustCompile(`^[\w\d\-_]+$`)
	is = validName.MatchString(name)
	return
}
