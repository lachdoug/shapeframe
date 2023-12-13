{{ define "shapes/create" -}}
{{ with .Payload -}}
Successfully added shape {{ .Shape }} to frame {{ .Frame }} workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
