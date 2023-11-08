{{ define "workspaces/directories" -}}
Directories:{{ if eq (len .) 0 }} none{{ else }}
{{ range . -}}
- {{ include "workspaces/directory" . 2 0 }}
{{ end -}}
{{ end -}}
{{ end -}}
