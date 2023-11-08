{{ define "workspaces/shape" -}}
Name: {{ .Name }}
About: {{ .About }}
Shaper: {{ .Shaper }}
{{ include "configurations/configuration" (.Configuration) }}
{{ end -}}