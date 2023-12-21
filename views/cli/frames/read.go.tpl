{{ define "frames/read" -}}
{{ with .Payload -}}
Name: {{ .Name }}
About: {{ .About }}
Workspace: {{ .Workspace }}
Parent: {{ if eq .Parent "" }}<unset>{{ else }}{{ .Parent }}{{ end }}
Framer: {{ .Framer }}
Configuration: {{ include "configurations/configuration" .Configuration }}
Shapes: {{ include "frames/shapes" .Shapes }}
{{ end -}}
{{ end -}}
