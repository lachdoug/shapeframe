package utils

func StringChunks(s string, max int) (ss []string) {
	if len(s) == 0 {
		return nil
	}
	if max >= len(s) {
		return []string{s}
	}
	ss = make([]string, 0, (len(s)-1)/max+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == max {
			ss = append(ss, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	ss = append(ss, s[currentStart:])
	return ss
}
