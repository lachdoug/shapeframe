package utils

func YamlReadFile(dirPath string, name string, model any) (err error) {
	err = YamlUnmarshal(ReadFile(YamlFilePath(dirPath, name)), model)
	return
}
