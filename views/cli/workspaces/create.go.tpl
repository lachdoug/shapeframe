{{ define "workspaces/create" -}}
{{ with .Payload -}}
Successfully initialized workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
