package utils

import "regexp"

func ScriptInterpreter(filePath string) (in string) {
	c := ReadFile(filePath)
	r := regexp.MustCompile(`^#![^\s]*\s(\w+)`)
	in = r.FindString(string(c))
	return
}
