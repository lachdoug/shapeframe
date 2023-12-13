{{ define "contexts/update" -}}
{{ with .Payload -}}
Successfully changed context
From: {{ include "contexts/context" .From 2 0 }}
To: {{ include "contexts/context" .To 2 0 }}
{{ end -}}
{{ end -}}