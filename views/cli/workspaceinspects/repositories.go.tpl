{{ define "workspaceinspects/repositories" -}}
Repositories:{{ if eq (len .) 0 }} <none>
{{ else }} 
{{ range . -}}
- {{ include "workspaceinspects/repository" . 2 0 }}
{{ end -}}
{{ end -}}
{{ end -}}
