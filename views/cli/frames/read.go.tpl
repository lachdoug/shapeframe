{{ define "frames/read" -}}
{{ with .Result -}}
Name: {{ .Name }}
About: {{ .About }}
Workspace: {{ .Workspace }}
Parent: {{ if eq .Parent "" }}<none>{{ else }}{{ .Parent }}{{ end }}
Framer: {{ .Framer }}
{{ include "configurations/configuration" (.Configuration) }}
{{ include "frames/shapes" (.Shapes) }}
{{ end -}}
{{ end -}}
