{{ define "repositories/delete" -}}
{{ with .Payload -}}
Successfully removed repository {{ .URI }} from workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
