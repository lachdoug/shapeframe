{{ define "shapers/index" -}}
{{ with .Payload -}}
{{ if le (len .Lines) 1 -}}
No shapers
{{ else -}}
{{ range .Lines -}}
{{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
