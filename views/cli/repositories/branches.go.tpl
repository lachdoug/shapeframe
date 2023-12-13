{{ define "repositories/branches" -}}
{{ if eq (len .) 0 }}<none>{{ else }}
{{ range . -}}
- {{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
