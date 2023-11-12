{{ define "workspaceinspects/frame/ancestors" -}}
Ancestors:{{ if eq (len .) 0 }} <none>
{{ else }}
{{ range . -}}
- {{ . }}
{{ end -}}
{{ end -}}
{{ end -}}
