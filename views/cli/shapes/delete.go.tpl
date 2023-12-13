{{ define "shapes/delete" -}}
{{ with .Payload -}}
Successfully removed shape {{ .Shape }} from frame {{ .Frame }} workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
