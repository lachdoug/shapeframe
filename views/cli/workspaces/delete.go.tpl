{{ define "workspaces/delete" -}}
{{ with .Payload -}}
Successfully removed workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
