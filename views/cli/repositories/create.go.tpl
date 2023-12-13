{{ define "repositories/create" -}}
{{ with .Payload -}}
Successfully added repository {{ .URI }} to workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
