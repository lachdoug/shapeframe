{{ define "configurations/settings" -}}
{{ if or (eq . nil) (eq (len .) 0) }}<unset>{{ else }}
{{ range . -}}{{ include "configurations/setting" . 2 }}
{{ end -}}
{{ end -}}
{{ end -}}
