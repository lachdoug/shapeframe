{{ define "contexts/from" -}}
From:{{ if not (.Workspace) }} none{{ else }}
{{ include "contexts/context" . 2 }}
{{ end -}}
{{ end -}}
