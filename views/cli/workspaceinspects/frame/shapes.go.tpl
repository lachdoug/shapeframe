{{ define "workspaceinspects/frame/shapes" -}}
{{ if eq (len .) 0 }}<none>{{ else }}
{{ range . -}}
- {{ include "workspaceinspects/frame/shape" . 2 0 }}
{{ end -}}
{{ end -}}
{{ end -}}
