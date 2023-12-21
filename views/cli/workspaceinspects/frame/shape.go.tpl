{{ define "workspaceinspects/frame/shape" -}}
Name: {{ .Name }}
About: {{ .About }}
Shaper: {{ .Shaper }}
{{ with .Configuration -}}
Configuration:
  Shape: {{ include "configurations/configuration" .Shape 4 0 }}
  Frame: {{ include "configurations/configuration" .Frame 4 0 }}
{{ end -}}
{{ end -}}
