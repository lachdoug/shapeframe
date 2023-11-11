{{ define "contexts/read" -}}
{{ with .Result -}}
{{ include "contexts/context" . }}
{{ end -}}
{{ end -}}
