{{ define "workspaceinspects/frames" -}}
Frames:{{ if eq (len .) 0 }} <none>{{ else }}
{{ range . -}}
- {{ include "workspaceinspects/frame" . 2 0 }}
{{ end -}}
{{ end -}}
{{ end -}}
