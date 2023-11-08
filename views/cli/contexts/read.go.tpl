{{ define "contexts/read" -}}
{{ with .Result -}}
Context:{{ if not (.Workspace) }} none
{{ else }}
{{ include "contexts/context" . 2 }}
{{ end -}}
{{ end -}}
{{ end -}}
