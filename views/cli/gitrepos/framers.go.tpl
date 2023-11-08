{{ define "gitrepos/framers" -}}
{{ range . -}}
- {{ include "gitrepos/framer" . 2 0 }}
{{ end -}}
{{ end -}}
