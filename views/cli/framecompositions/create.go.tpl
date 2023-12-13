{{ define "framecompositions/create" -}}
{{ with .Payload -}}
Successfully composed frame {{ .Frame }} in workspace {{ .Workspace }}
{{ end -}}
{{ end -}}

