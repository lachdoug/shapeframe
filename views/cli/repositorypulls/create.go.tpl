{{ define "repositorypulls/create" -}}
{{ with .Payload -}}
Successfully pulled repository {{ .URI }} in workspace {{ .Workspace }}
{{ end -}}
{{ end -}}

