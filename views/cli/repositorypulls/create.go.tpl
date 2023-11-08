{{ define "repositorypulls/create" -}}
{{ with .Result -}}
Successfully pulled repository {{ .URI }} in workspace {{ .Workspace }}
{{ end -}}
{{ end -}}

