{{ with index . "Result" -}}
Hello {{ index . "Name" }}!{{ if index . "Extra"}}!!!!!!!!!!{{ end }}
{{ end -}}