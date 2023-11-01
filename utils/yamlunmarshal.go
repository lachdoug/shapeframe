package utils

import "gopkg.in/yaml.v2"

func YamlUnmarshal(y []byte, model any) (err error) {
	err = yaml.Unmarshal(y, model)
	return
}
