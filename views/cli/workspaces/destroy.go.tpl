{{ define "workspaces/destroy" -}}
{{ with .Result -}}
Successfully removed workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
