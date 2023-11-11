{{ define "workspaceinspects/frame" -}}
Name: {{ .Name }}
About: {{ .About }}
Framer: {{ .Framer }}
{{ include "configurations/configuration" (.Configuration) }}
{{ include "workspaceinspects/frame/relations" (.Relations) }}
{{ include "workspaceinspects/frame/shapes" (.Shapes) }}
{{ end -}}