{{ define "workspaceinspects/frame/shape" -}}
Name: {{ .Name }}
About: {{ .About }}
Shaper: {{ .Shaper }}
{{ include "configurations/configuration" (.Configuration) }}
{{ end -}}