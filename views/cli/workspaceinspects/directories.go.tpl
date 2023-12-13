{{ define "workspaceinspects/directories" -}}
{{ if eq (len .) 0 }}<none>{{ else }}
{{ range . -}}
- {{ include "workspaceinspects/directory" . 2 0 }}
{{ end -}}
{{ end -}}
{{ end -}}
