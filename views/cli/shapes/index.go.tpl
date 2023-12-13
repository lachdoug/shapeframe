{{ define "shapes/index" -}}
{{ with .Payload -}}
{{ if le (len .Lines) 1 -}}
No shapes
{{ else -}}
{{ range .Lines -}}
{{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
{{ end -}}
