{{ define "workspaces/create" -}}
{{ with .Result -}}
Successfully added workspace {{ .Workspace }}
{{ end -}}
{{ end -}}
