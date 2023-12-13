{{ define "contexts/read" -}}
{{ with .Payload -}}
Context: {{ include "contexts/context" . 2 0 }}
{{ end -}}
{{ end -}}
