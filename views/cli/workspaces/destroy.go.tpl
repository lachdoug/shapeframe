{{ with index . "Result" -}}
Destroyed workspace {{ index . "Name" }}
{{ end -}}