{{ define "workspaces/frames" -}}
Frames:{{ if eq (len .) 0 }} <none>{{ else }}
{{ range . -}}
- {{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
