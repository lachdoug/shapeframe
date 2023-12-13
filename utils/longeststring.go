package utils

func LongestString(ss []string) (max int) {
	for _, value := range ss {
		if l := len(value); l > max {
			max = l
		}
	}
	return
}
