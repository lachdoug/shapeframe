{{ define "workspaces/repositories" -}}
Repositories:{{ if eq (len .) 0 }} <none>
{{ else }} 
{{ range . -}}
- {{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
