{{ define "workspaceinspects/frame/descendents" -}}
Descendents:{{ if eq (len .) 0 }} <none>{{ else }}
{{ range . -}}
- {{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
