{{ define "gitrepos/shapers" -}}
{{ range . -}}
- {{ include "gitrepos/shaper" . 2 0 }}
{{ end -}}
{{ end -}}
