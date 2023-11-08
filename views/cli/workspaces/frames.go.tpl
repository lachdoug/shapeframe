{{ define "workspaces/frames" -}}
Frames:{{ if eq (len .) 0 }} none{{ else }}
{{ range . -}}
- {{ include "workspaces/frame" . 2 0 }}
{{ end -}}
{{ end -}}
{{ end -}}
