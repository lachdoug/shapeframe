package utils

func MapKeys(m map[string]any) (keys []string) {
	keys = make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return
}
