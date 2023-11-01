package utils

import (
	"encoding/json"
)

func JsonMarshalPretty(v any) (j []byte) {
	j, err := json.MarshalIndent(v, "", "  ")
	checkErr(err)
	return
}
