{{ define "contexts/context" -}}
{{ if not (.Workspace) }}<none>{{ else }}
Workspace: {{ .Workspace -}}
{{ if .Frame }}
Frame: {{ .Frame -}}
{{ end -}}
{{ if .Shape }}
Shape: {{ .Shape -}}
{{ end -}}
{{ end -}}
{{ end -}}

