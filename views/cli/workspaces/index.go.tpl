{{ $lines := index . "Lines" -}}
{{ if le (len $lines) 1 -}}
No workspaces
{{ else -}}
{{ range $lines -}}
{{ . }}
{{ end -}}
{{ end -}}
