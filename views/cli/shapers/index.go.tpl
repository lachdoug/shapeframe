{{ define "shapers/index" -}}
{{ $lines := .Lines -}}
{{ if le (len $lines) 1 -}}
No shapers
{{ else -}}
{{ range $lines -}}
{{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
