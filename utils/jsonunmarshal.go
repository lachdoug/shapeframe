package utils

import "encoding/json"

func JsonUnmarshal(j []byte, v any) {
	err := json.Unmarshal(j, v)
	checkErr(err)
}
