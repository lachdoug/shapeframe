{{ define "contexts/update" -}}
{{ with .Result -}}
Successfully changed context
{{ include "contexts/from" (.From) }}
{{ include "contexts/to" (.To) }}
{{ end -}}
{{ end -}}