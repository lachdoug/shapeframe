{{ define "contexts/to" -}}
To:{{ if not (.Workspace) }} none{{ else }}
{{ include "contexts/context" . 2 }}
{{ end -}}
{{ end -}}
