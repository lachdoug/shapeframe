{{ define "workspaces/frame" -}}
Name: {{ .Name }}
About: {{ .About }}
Framer: {{ .Framer }}
{{ include "configurations/configuration" (.Configuration) }}
{{ include "workspaces/shapes" (.Shapes) }}
{{ end -}}