{{ define "shapes/read" -}}
{{ with .Result -}}
Name: {{ .Name }}
About: {{ .About }}
Workspace: {{ .Workspace }}
Frame: {{ .Frame }}
{{ include "configurations/configuration" (.Configuration) }}
{{ end -}}
{{ end -}}
