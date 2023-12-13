{{ define "workspaceinspects/frame" -}}
Name: {{ .Name }}
About: {{ .About }}
Framer: {{ .Framer }}
Configuration: {{ include "configurations/configuration" .Configuration 2 0 }}
Relations:
{{ include "workspaceinspects/frame/relations" .Relations 2 }}
Shapes: {{ include "workspaceinspects/frame/shapes" .Shapes 2 0 }}
{{ end -}}