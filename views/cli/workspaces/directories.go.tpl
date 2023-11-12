{{ define "workspaces/directories" -}}
Directories:{{ if eq (len .) 0 }} <none>
{{ else }}
{{ range . -}}
- {{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
