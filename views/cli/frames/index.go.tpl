{{ $lines := index . "Lines" -}}
{{ if le (len $lines) 1 -}}
No frames
{{ else -}}
{{ range $lines -}}
{{ . }}
{{ end -}}
{{ end -}}
