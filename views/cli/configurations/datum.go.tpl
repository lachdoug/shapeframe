{{ define "configurations/datum" -}}
{{ if stringContains .Setting "\n" -}}
{{ .Key }}:
{{ indentLines .Setting 2 }}
{{ else -}}
{{ .Key }}: {{ .Setting }}
{{ end -}}
{{ end -}}
