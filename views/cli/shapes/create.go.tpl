{{ define "shapes/create" -}}
{{ with .Result -}}
Successfully added shape {{ .Shape }} to frame {{ .Frame }} workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
