{{ define "workspaces/create" -}}
{{ with .Payload -}}
Successfully added workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
