{{ define "gitrepos/shapers" -}}
{{ if eq (len .) 0 }}<none>{{ else }}
{{ range . -}}
- {{ include "gitrepos/shaper" . 2 0 }}
{{ end -}}
{{ end -}}
{{ end -}}
