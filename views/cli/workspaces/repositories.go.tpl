{{ define "workspaces/repositories" -}}
Repositories:{{ if eq (len .) 0 }} none{{ else }} 
{{ range . -}}
- {{ include "workspaces/repository" . 2 0 }}
{{ end -}}
{{ end -}}
{{ end -}}
