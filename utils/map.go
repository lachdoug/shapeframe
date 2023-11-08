package utils

func Map(a any) (m *map[string]any) {
	m = &map[string]any{}
	JsonUnmarshal(JsonMarshal(a), m)
	return
}
