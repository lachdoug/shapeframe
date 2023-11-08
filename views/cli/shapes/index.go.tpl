{{ define "shapes/index" -}}
{{ $lines := .Lines -}}
{{ if le (len $lines) 1 -}}
No shapes
{{ else -}}
{{ range $lines -}}
{{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
