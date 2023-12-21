{{ define "shapes/read" -}}
{{ with .Payload -}}
Name: {{ .Name }}
About: {{ .About }}
Workspace: {{ .Workspace }}
Frame: {{ .Frame }}
{{ with .Configuration -}}
Configuration:
  Shape: {{ include "configurations/configuration" .Shape 4 0 }}
  Frame: {{ include "configurations/configuration" .Frame 4 0 }}
{{ end -}}
{{ end -}}
{{ end -}}

