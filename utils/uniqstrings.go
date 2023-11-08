package utils

import "github.com/mpvl/unique"

func UniqStrings(s *[]string) {
	unique.Sort(unique.StringSlice{P: s})
}
