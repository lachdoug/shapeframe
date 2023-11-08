{{ define "shapes/destroy" -}}
{{ with .Result -}}
Successfully removed shape {{ .Shape }} from frame {{ .Frame }} workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
