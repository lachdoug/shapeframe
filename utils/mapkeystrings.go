package utils

func MapKeyStrings(i any) any {
	switch x := i.(type) {
	case map[any]any:
		m2 := map[string]any{}
		for k, v := range x {
			m2[k.(string)] = MapKeyStrings(v)
		}
		return m2
	case []any:
		for i, v := range x {
			x[i] = MapKeyStrings(v)
		}
	}
	return i
}
