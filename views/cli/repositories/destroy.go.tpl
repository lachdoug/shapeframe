{{ define "repositories/destroy" -}}
{{ with .Result -}}
Successfully removed repository {{ .URI }} from workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
