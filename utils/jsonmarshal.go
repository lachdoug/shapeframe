package utils

import (
	"encoding/json"
)

func JsonMarshal(v any) (j []byte) {
	j, err := json.Marshal(v)
	checkErr(err)
	return
}
