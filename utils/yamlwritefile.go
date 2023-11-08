package utils

func YamlWriteFile(dirPath string, fileName string, data any) {
	WriteFile(YamlFilePath(dirPath, fileName), YamlMarshal(data))
	return
}
