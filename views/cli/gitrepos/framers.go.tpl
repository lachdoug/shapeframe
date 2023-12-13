{{ define "gitrepos/framers" -}}
{{ if eq (len .) 0 }}<none>{{ else }}
{{ range . -}}
- {{ include "gitrepos/framer" . 2 0 }}
{{ end -}}
{{ end -}}
{{ end -}}
