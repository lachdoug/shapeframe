{{ define "framers/index" -}}
{{ with .Payload -}}
{{ $lines := .Lines -}}
{{ if le (len $lines) 1 -}}
No framers
{{ else -}}
{{ range $lines -}}
{{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
