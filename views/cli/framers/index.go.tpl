{{ $lines := index . "Lines" -}}
{{ if le (len $lines) 1 -}}
No framers
{{ else -}}
{{ range $lines -}}
{{ . }}
{{ end -}}
{{ end -}}
