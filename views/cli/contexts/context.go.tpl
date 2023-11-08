{{ define "contexts/context" -}}
{{ if (.Workspace) }}
Workspace: {{ .Workspace -}}
{{ end -}}
{{ if .Frame }}
Frame: {{ .Frame -}}
{{ end -}}
{{ if .Shape }}
Shape: {{ .Shape -}}
{{ end -}}
{{ end -}}

