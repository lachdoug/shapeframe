{{ define "contexts/context" -}}
{{ if not (.Frame) }}<unset>{{ else }}
Frame: {{ .Frame -}}
{{ end -}}
{{ if .Shape }}
Shape: {{ .Shape -}}
{{ end -}}
{{ end -}}

