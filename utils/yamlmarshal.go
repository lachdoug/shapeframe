package utils

import "gopkg.in/yaml.v2"

func YamlMarshal(data any) (y []byte) {
	y, err := yaml.Marshal(data)
	checkErr(err)
	return
}
