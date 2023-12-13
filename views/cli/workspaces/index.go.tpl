{{ define "workspaces/index" -}}
{{ with .Payload -}}
{{ if le (len .Lines) 1 -}}
No workspaces
{{ else -}}
{{ range .Lines -}}
{{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
