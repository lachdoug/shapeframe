{{ define "workspaces/shapes" -}}
Shapes:{{ if eq (len .) 0 }} none{{ else }}
{{ range . -}}
- {{ include "workspaces/shape" . 2 0 }}
{{ end -}}
{{ end -}}
{{ end -}}
