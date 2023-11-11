{{ define "contexts/update" -}}
{{ with .Result -}}
Successfully changed context
From:
{{ include "contexts/context" (.From) 2 }}
To:
{{ include "contexts/context" (.To) 2 }}
{{ end -}}
{{ end -}}