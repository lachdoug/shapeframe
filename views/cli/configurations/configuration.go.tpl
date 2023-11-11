{{ define "configurations/configuration" -}}
Configuration:{{ if eq (len .) 0 }} <none>{{ else }}
{{ range . -}}{{ include "configurations/setting" . 2 }}
{{ end -}}
{{ end -}}
{{ end -}}
