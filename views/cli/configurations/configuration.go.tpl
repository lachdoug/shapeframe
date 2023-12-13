{{ define "configurations/configuration" -}}
{{ if or (eq . nil) (eq (len .) 0) }}<none>{{ else }}
{{ range . -}}{{ include "configurations/datum" . 2 }}
{{ end -}}
{{ end -}}
{{ end -}}
